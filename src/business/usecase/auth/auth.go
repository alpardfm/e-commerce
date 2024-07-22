package auth

import (
	"context"
	"fmt"
	"strconv"
	"time"

	locDom "github.com/alpardfm/e-commerce/src/business/domain/location"
	roleDom "github.com/alpardfm/e-commerce/src/business/domain/role"
	userDom "github.com/alpardfm/e-commerce/src/business/domain/users"
	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/e-commerce/src/utils/config"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/alpardfm/go-toolkit/distance"
	"github.com/alpardfm/go-toolkit/errors"
	"github.com/alpardfm/go-toolkit/log"
	"github.com/alpardfm/go-toolkit/tokens"
	"github.com/dgrijalva/jwt-go/v4"
)

type Interface interface {
	LoginDashboard(ctx context.Context, paramB entity.AuthLoginDashboardBody, paramH entity.AuthLoginDashboardHeader) (entity.AuthLoginDashboardResponse, error)
}

type auth struct {
	log log.Interface
	dom domain
	cfg config.Application
}

type domain struct {
	user     userDom.Interface
	location locDom.Interface
	role     roleDom.Interface
}

func Init(log log.Interface, cfg config.Application, userDom userDom.Interface, locDom locDom.Interface, roleDom roleDom.Interface) Interface {
	return &auth{
		log: log,
		cfg: cfg,
		dom: domain{
			user:     userDom,
			location: locDom,
			role:     roleDom,
		},
	}
}

func (a *auth) LoginDashboard(ctx context.Context, paramB entity.AuthLoginDashboardBody, paramH entity.AuthLoginDashboardHeader) (entity.AuthLoginDashboardResponse, error) {
	user, err := a.dom.user.GetDetail(ctx, entity.Users{
		Email:    paramB.Email,
		Password: paramB.Password,
	})
	if err != nil {
		if errors.GetCode(err) == codes.CodeSQLRead {
			return entity.AuthLoginDashboardResponse{}, errors.NewWithCode(codes.CodeUnauthorized, "Email Or Password Is Wrong")
		}
		return entity.AuthLoginDashboardResponse{}, err
	}

	loc, err := a.dom.location.GetDetail(ctx, entity.Location{
		Secret: paramB.Secret,
	})
	if err != nil {
		if errors.GetCode(err) == codes.CodeSQLRead {
			return entity.AuthLoginDashboardResponse{}, errors.NewWithCode(codes.CodeUnauthorized, "Secret Is Wrong")
		}
		return entity.AuthLoginDashboardResponse{}, err
	}

	latString := [4]string{loc.Lat, loc.Long, paramH.Lat, paramH.Long}
	latFloat := []float64{}
	for _, v := range latString {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return entity.AuthLoginDashboardResponse{}, errors.NewWithCode(codes.CodeInvalidValue, err.Error())
		}
		latFloat = append(latFloat, value)
	}

	distance := int64(distance.CalculateDistance(ctx, latFloat[0], latFloat[1], latFloat[2], latFloat[3]))
	if distance > loc.Distance {
		return entity.AuthLoginDashboardResponse{}, errors.NewWithCode(codes.CodeUnauthorized, "Your location is too far from the specified point")
	}

	claims := entity.TokenLoginDashboardClaims{
		UID:    fmt.Sprintf("%v", user.ID),
		Email:  user.Email,
		RoleID: fmt.Sprintf("%v", user.RoleID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * time.Duration(a.cfg.JWT.DashboardJWTTokenExpirationMinute))),
			IssuedAt:  jwt.Now(),
		},
	}

	jwtToken, err := tokens.NewJWTToken[entity.TokenLoginDashboardClaims](claims, []byte(a.cfg.JWT.JWTTokenKey))
	if err != nil {
		return entity.AuthLoginDashboardResponse{}, err
	}

	role, err := a.dom.role.GetDetail(ctx, entity.Role{
		ID: user.RoleID,
	})
	if err != nil {
		return entity.AuthLoginDashboardResponse{}, err
	}

	return entity.AuthLoginDashboardResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     role.Name,
		Token:    jwtToken,
	}, nil
}
