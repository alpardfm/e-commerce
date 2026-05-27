package rest

import (
	"strconv"

	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/gin-gonic/gin"
)

func (r *rest) GetListProductsDashboard(ctx *gin.Context) {
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 64)

	paginate := entity.PaginationProducts{
		Page:  page,
		Limit: limit,
	}

	result, err := r.uc.Products.GetList(ctx.Request.Context(), entity.Products{}, paginate)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) GetDetailProducts(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	result, err := r.uc.Products.GetDetail(ctx.Request.Context(), entity.Products{ID: id})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) CreateProducts(ctx *gin.Context) {
	var param entity.Products
	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	result, err := r.uc.Products.Create(ctx.Request.Context(), param, "admin")
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) UpdateProducts(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	var param entity.Products
	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}
	param.ID = id

	result, err := r.uc.Products.Update(ctx.Request.Context(), param, "admin")
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) DeleteProducts(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	result, err := r.uc.Products.Delete(ctx.Request.Context(), entity.Products{ID: id}, "admin")
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}
