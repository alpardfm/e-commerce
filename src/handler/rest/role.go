package rest

import (
	"strconv"

	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/gin-gonic/gin"
)

func (r *rest) GetListRoleDashboard(ctx *gin.Context) {
	tokens := ctx.GetHeader("Authorization")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	name := ctx.Query("name")

	paginate := entity.PaginationRole{}
	param := entity.Role{}

	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			r.httpRespError(ctx, err)
			return
		}

		paginate.Page = int64(pageInt)
	} else {
		paginate.Page = 1
	}

	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			r.httpRespError(ctx, err)
			return
		}

		paginate.Limit = int64(limitInt)
	} else {
		paginate.Limit = 10
	}

	if name != "" {
		param.Name = name
	}

	result, err := r.uc.Role.GetListDashboard(ctx, param, paginate, tokens)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) GetDetailRole(ctx *gin.Context) {
	tokens := ctx.GetHeader("Authorization")
	id := ctx.Param("id")

	param := entity.Role{}

	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			r.httpRespError(ctx, err)
			return
		}

		param.ID = int64(idInt)
	}

	result, err := r.uc.Role.GetDetail(ctx, param, tokens)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) CreateRole(ctx *gin.Context) {
	tokens := ctx.GetHeader("Authorization")
	var body entity.BodyRole
	ctx.Bind(&body)

	result, err := r.uc.Role.Create(ctx, entity.Role{
		Name: body.Name,
	}, tokens)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) UpdateRole(ctx *gin.Context) {
	tokens := ctx.GetHeader("Authorization")
	id := ctx.Param("id")

	param := entity.Role{}

	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			r.httpRespError(ctx, err)
			return
		}

		param.ID = int64(idInt)
	}

	var body entity.BodyRole
	ctx.Bind(&body)
	param.Name = body.Name

	result, err := r.uc.Role.Update(ctx, param, tokens)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) DeleteRole(ctx *gin.Context) {
	tokens := ctx.GetHeader("Authorization")
	id := ctx.Param("id")

	param := entity.Role{}

	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			r.httpRespError(ctx, err)
			return
		}

		param.ID = int64(idInt)
	}

	result, err := r.uc.Role.Delete(ctx, param, tokens)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}
