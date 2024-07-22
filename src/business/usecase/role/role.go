package role

import (
	"context"
	"fmt"
	"strconv"
	"time"

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
	GetListDashboard(ctx context.Context, param entity.Role, paginate entity.PaginationRole, token string) (entity.ResponseRole, error)
	GetDetail(ctx context.Context, param entity.Role, token string) (entity.Role, error)
	Create(ctx context.Context, param entity.Role, token string) (entity.Role, error)
	Update(ctx context.Context, param entity.Role, token string) (entity.Role, error)
	Delete(ctx context.Context, param entity.Role, token string) (entity.Role, error)
}

type role struct {
	log log.Interface
	cfg config.Application
	dom domain
}

type domain struct {
	role roleDom.Interface
}

func Init(log log.Interface, cfg config.Application, roleDom roleDom.Interface) Interface {
	return &role{
		log: log,
		cfg: cfg,
		dom: domain{
			role: roleDom,
		},
	}
}

func (r *role) GetListDashboard(ctx context.Context, param entity.Role, paginate entity.PaginationRole, token string) (entity.ResponseRole, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(r.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.ResponseRole{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.ResponseRole{}, err
	}

	roleIdInt, err := strconv.Atoi(claims.RoleID)
	if err != nil {
		return entity.ResponseRole{}, err
	}

	role, err := r.dom.role.GetDetail(ctx, entity.Role{
		ID: int64(roleIdInt),
	})
	if err != nil {
		return entity.ResponseRole{}, err
	}

	if role.Name != "admin" {
		return entity.ResponseRole{}, errors.NewWithCode(codes.CodeUnauthorized, err.Error())
	}

	r.log.Debug(ctx, fmt.Sprintf("Get List Role Dashboard By %v", claims.UID))

	results, err := r.dom.role.GetList(ctx, param, func(_, suffix *string) error {
		*suffix = fmt.Sprintf("AND is_deleted = %d", 0)
		return nil
	})
	if err != nil {
		return entity.ResponseRole{}, err
	}

	resultListWithPagination := helper.Paginate[entity.Role](results, int(paginate.Page), int(paginate.Limit))
	var totalRows, totalPages int64
	totalRows = int64(len(results))

	if totalRows != 0 && paginate.Limit != 0 {
		totalPages = (totalRows + paginate.Limit - 1) / paginate.Limit
	}

	return entity.ResponseRole{
		Limit:      paginate.Limit,
		Page:       paginate.Page,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       resultListWithPagination,
	}, nil
}

func (r *role) GetDetail(ctx context.Context, param entity.Role, token string) (entity.Role, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(r.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Role{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Role{}, err
	}

	r.log.Debug(ctx, fmt.Sprintf("Get Detail Role By %v", claims.UID))

	result, err := r.dom.role.GetDetail(ctx, param)
	if err != nil {
		return entity.Role{}, err
	}

	return result, nil
}

func (r *role) Create(ctx context.Context, param entity.Role, token string) (entity.Role, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(r.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Role{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Role{}, err
	}

	r.log.Debug(ctx, fmt.Sprintf("Create New Role By %v", claims.UID))

	param.CreatedAt = time.Now().UTC()
	param.CreatedBy = fmt.Sprintf("%v", claims.UID)
	param.IsDeleted = 0

	results, err := r.dom.role.Create(ctx, param)
	if err != nil {
		return entity.Role{}, err
	}

	return results, nil
}

func (r *role) Update(ctx context.Context, param entity.Role, token string) (entity.Role, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(r.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Role{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Role{}, err
	}

	r.log.Debug(ctx, fmt.Sprintf("Update Role By %v", claims.UID))

	param.UpdatedAt = time.Now().UTC()
	param.UpdatedBy = fmt.Sprintf("%v", claims.UID)

	results, err := r.dom.role.Update(ctx, param)
	if err != nil {
		return entity.Role{}, err
	}

	return results, nil
}
func (r *role) Delete(ctx context.Context, param entity.Role, token string) (entity.Role, error) {
	jwtTokens, err := tokens.ValidateJWTToken[entity.TokenLoginDashboardClaims](token, []byte(r.cfg.JWT.JWTTokenKey), entity.TokenLoginDashboardClaims{})
	if err != nil {
		return entity.Role{}, err
	}

	claims, err := tokens.GetClaimsOfJWTToken[entity.TokenLoginDashboardClaims](*jwtTokens)
	if err != nil {
		return entity.Role{}, err
	}

	r.log.Debug(ctx, fmt.Sprintf("Delete Role By %v", claims.UID))

	param.DeletedAt = time.Now().UTC()
	param.DeletedBy = fmt.Sprintf("%v", claims.UID)
	param.IsDeleted = 1

	result, err := r.dom.role.Delete(ctx, param)
	if err != nil {
		return entity.Role{}, err
	}

	return result, nil
}
