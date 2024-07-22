package location

import (
	"context"
	"fmt"
	"strconv"
	"time"

	locDom "github.com/alpardfm/e-commerce/src/business/domain/location"
	roleDom "github.com/alpardfm/e-commerce/src/business/domain/role"
	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/e-commerce/src/utils/config"
	"github.com/alpardfm/e-commerce/src/utils/helper"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/alpardfm/go-toolkit/errors"
	"github.com/alpardfm/go-toolkit/log"
	"github.com/alpardfm/go-toolkit/tokens"
)

type Interface interface {
	GetListDashboard(ctx context.Context, param entity.Location, paginate entity.PaginationLocation, token string) (entity.ResponseLocation, error)
	GetDetail(ctx context.Context, param entity.Location, token string) (entity.Location, error)
	Create(ctx context.Context, param entity.Location, token string) (entity.Location, error)
	Update(ctx context.Context, param entity.Location, token string) (entity.Location, error)
	Delete(ctx context.Context, param entity.Location, token string) (entity.Location, error)
}

type location struct {
	log log.Interface
	cfg config.Application
	dom domain
}

type domain struct {
	location locDom.Interface
	role     roleDom.Interface
}

func Init(log log.Interface, cfg config.Application, locDom locDom.Interface, roleDom roleDom.Interface) Interface {
	return &location{
		log: log,
		cfg: cfg,
		dom: domain{
			location: locDom,
			role:     roleDom,
		},
	}
}

func (l *location) GetListDashboard(ctx context.Context, param entity.Location, paginate entity.PaginationLocation, token string) (entity.ResponseLocation, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(l.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.ResponseLocation{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.ResponseLocation{}, err
	}

	roleIdInt, err := strconv.Atoi(claims.RoleID)
	if err != nil {
		return entity.ResponseLocation{}, err
	}

	role, err := l.dom.role.GetDetail(ctx, entity.Role{
		ID: int64(roleIdInt),
	})
	if err != nil {
		return entity.ResponseLocation{}, err
	}

	if role.Name != "admin" {
		return entity.ResponseLocation{}, errors.NewWithCode(codes.CodeUnauthorized, err.Error())
	}

	l.log.Debug(ctx, fmt.Sprintf("Get List Location Dashboard By %v", claims.UID))

	results, err := l.dom.location.GetList(ctx, param, func(_, suffix *string) error {
		*suffix = fmt.Sprintf("AND is_deleted = %d", 0)
		return nil
	})
	if err != nil {
		return entity.ResponseLocation{}, err
	}

	resultListWithPagination := helper.Paginate[entity.Location](results, int(paginate.Page), int(paginate.Limit))
	var totalRows, totalPages int64
	totalRows = int64(len(results))

	if totalRows != 0 && paginate.Limit != 0 {
		totalPages = (totalRows + paginate.Limit - 1) / paginate.Limit
	}

	return entity.ResponseLocation{
		Limit:      paginate.Limit,
		Page:       paginate.Page,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       resultListWithPagination,
	}, nil
}
func (l *location) GetDetail(ctx context.Context, param entity.Location, token string) (entity.Location, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(l.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Location{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Location{}, err
	}

	l.log.Debug(ctx, fmt.Sprintf("Get Detail Location By %v", claims.UID))

	result, err := l.dom.location.GetDetail(ctx, param)
	if err != nil {
		return entity.Location{}, err
	}

	return result, nil
}

func (l *location) Create(ctx context.Context, param entity.Location, token string) (entity.Location, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(l.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Location{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Location{}, err
	}

	l.log.Debug(ctx, fmt.Sprintf("Create New Location By %v", claims.UID))

	param.CreatedAt = time.Now().UTC()
	param.CreatedBy = claims.UID
	param.IsDeleted = 0

	result, err := l.dom.location.Create(ctx, param)
	if err != nil {
		return entity.Location{}, err
	}

	return result, nil
}

func (l *location) Update(ctx context.Context, param entity.Location, token string) (entity.Location, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(l.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Location{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Location{}, err
	}

	l.log.Debug(ctx, fmt.Sprintf("Update Location By %v", claims.UID))

	param.UpdatedAt = time.Now().UTC()

	result, err := l.dom.location.Update(ctx, param)
	if err != nil {
		return entity.Location{}, err
	}

	return result, nil
}

func (l *location) Delete(ctx context.Context, param entity.Location, token string) (entity.Location, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(l.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Location{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Location{}, err
	}

	l.log.Debug(ctx, fmt.Sprintf("Delete Location By %v", claims.UID))

	param.DeletedAt = time.Now().UTC()
	param.IsDeleted = 1

	result, err := l.dom.location.Delete(ctx, param)
	if err != nil {
		return entity.Location{}, err
	}

	return result, nil
}
