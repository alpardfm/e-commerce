package rest

import (
	"strconv"

	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/gin-gonic/gin"
)

func (r *rest) GetListCart(ctx *gin.Context) {
	// TODO: get user_id from JWT token
	userID, _ := strconv.ParseInt(ctx.DefaultQuery("user_id", "0"), 10, 64)

	result, err := r.uc.Cart.GetList(ctx.Request.Context(), userID)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) CreateCart(ctx *gin.Context) {
	var param entity.Cart
	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	result, err := r.uc.Cart.Create(ctx.Request.Context(), param, "user")
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) UpdateCart(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	var param entity.Cart
	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}
	param.ID = id

	result, err := r.uc.Cart.Update(ctx.Request.Context(), param, "user")
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) DeleteCart(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	result, err := r.uc.Cart.Delete(ctx.Request.Context(), entity.Cart{ID: id}, "user")
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}
