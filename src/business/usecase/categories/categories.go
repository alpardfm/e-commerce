package categories

import (
	"context"
	"fmt"
	"strconv"
	"time"

	categoriesDom "github.com/alpardfm/e-commerce/src/business/domain/categories"
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
	GetListDashboard(ctx context.Context, param entity.Categories, paginate entity.PaginationCategories, token string) (entity.ResponseCategories, error)
	GetDetail(ctx context.Context, param entity.Categories, token string) (entity.Categories, error)
	Create(ctx context.Context, param entity.Categories, token string) (entity.Categories, error)
	Update(ctx context.Context, param entity.Categories, token string) (entity.Categories, error)
	Delete(ctx context.Context, param entity.Categories, token string) (entity.Categories, error)
}

type categories struct {
	log log.Interface
	cfg config.Application
	dom domain
}

type domain struct {
	categories categoriesDom.Interface
	role       roleDom.Interface
}

func Init(log log.Interface, cfg config.Application, categororiesDom categoriesDom.Interface, roleDom roleDom.Interface) Interface {
	return &categories{
		log: log,
		cfg: cfg,
		dom: domain{
			categories: categororiesDom,
			role:       roleDom,
		},
	}
}

func (c *categories) GetListDashboard(ctx context.Context, param entity.Categories, paginate entity.PaginationCategories, token string) (entity.ResponseCategories, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(c.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.ResponseCategories{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.ResponseCategories{}, err
	}

	roleIdInt, err := strconv.Atoi(claims.RoleID)
	if err != nil {
		return entity.ResponseCategories{}, err
	}

	role, err := c.dom.role.GetDetail(ctx, entity.Role{
		ID: int64(roleIdInt),
	})
	if err != nil {
		return entity.ResponseCategories{}, err
	}

	if role.Name != "admin" {
		return entity.ResponseCategories{}, errors.NewWithCode(codes.CodeUnauthorized, err.Error())
	}

	c.log.Debug(ctx, fmt.Sprintf("Get List Categories Dashboard By %v", claims.UID))

	results, err := c.dom.categories.GetList(ctx, param, func(_, suffix *string) error {
		*suffix = fmt.Sprintf("AND is_deleted = %d", 0)
		return nil
	})
	if err != nil {
		return entity.ResponseCategories{}, err
	}

	resultListWithPagination := helper.Paginate[entity.Categories](results, int(paginate.Page), int(paginate.Limit))
	var totalRows, totalPages int64
	totalRows = int64(len(results))

	if totalRows != 0 && paginate.Limit != 0 {
		totalPages = (totalRows + paginate.Limit - 1) / paginate.Limit
	}

	return entity.ResponseCategories{
		Limit:      paginate.Limit,
		Page:       paginate.Page,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       resultListWithPagination,
	}, nil
}

func (c *categories) GetDetail(ctx context.Context, param entity.Categories, token string) (entity.Categories, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(c.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Categories{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Categories{}, err
	}

	c.log.Debug(ctx, fmt.Sprintf("Get Detail Categories By %v", claims.UID))

	result, err := c.dom.categories.GetDetail(ctx, param)
	if err != nil {
		return entity.Categories{}, err
	}

	return result, nil
}

func (c *categories) Create(ctx context.Context, param entity.Categories, token string) (entity.Categories, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(c.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Categories{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Categories{}, err
	}

	c.log.Debug(ctx, fmt.Sprintf("Create New Categories By %v", claims.UID))

	param.CreatedAt = time.Now().UTC()
	param.CreatedBy = fmt.Sprintf("%v", claims.UID)
	param.IsDeleted = 0

	result, err := c.dom.categories.Create(ctx, param)
	if err != nil {
		return entity.Categories{}, err
	}

	return result, nil
}

func (c *categories) Update(ctx context.Context, param entity.Categories, token string) (entity.Categories, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(c.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Categories{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Categories{}, err
	}

	c.log.Debug(ctx, fmt.Sprintf("Update Categories By %v", claims.UID))

	param.UpdatedAt = time.Now().UTC()
	param.UpdatedBy = fmt.Sprintf("%v", claims.UID)

	result, err := c.dom.categories.Update(ctx, param)
	if err != nil {
		return entity.Categories{}, err
	}

	return result, nil
}

func (c *categories) Delete(ctx context.Context, param entity.Categories, token string) (entity.Categories, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(c.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Categories{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Categories{}, err
	}

	c.log.Debug(ctx, fmt.Sprintf("Delete Categories By %v", claims.UID))

	param.DeletedAt = time.Now().UTC()
	param.DeletedBy = fmt.Sprintf("%v", claims.UID)
	param.IsDeleted = 1

	result, err := c.dom.categories.Delete(ctx, param)
	if err != nil {
		return entity.Categories{}, err
	}

	return result, nil
}
