package rest

import (
	"strconv"

	"github.com/alpardfm/e-commerce/src/business/usecase/orders"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/gin-gonic/gin"
)

func (r *rest) CreateOrder(ctx *gin.Context) {
	var param orders.CreateOrderInput
	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	result, err := r.uc.Orders.Create(ctx.Request.Context(), param, "user")
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) GetListOrders(ctx *gin.Context) {
	userID, _ := strconv.ParseInt(ctx.DefaultQuery("user_id", "0"), 10, 64)

	result, err := r.uc.Orders.GetListByUser(ctx.Request.Context(), userID)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) UpdateOrderStatus(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	var body struct {
		Status string `json:"status"`
	}
	if err := r.Bind(ctx, &body); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	result, err := r.uc.Orders.UpdateStatus(ctx.Request.Context(), id, body.Status, "admin")
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}
