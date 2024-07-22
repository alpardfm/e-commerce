package rest

import (
	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/gin-gonic/gin"
)

func (r *rest) LoginDashboard(ctx *gin.Context) {
	paramHeader := entity.AuthLoginDashboardHeader{}
	paramBody := entity.AuthLoginDashboardBody{}

	paramHeader.Lat = ctx.GetHeader("lat")
	paramHeader.Long = ctx.GetHeader("long")
	ctx.Bind(&paramBody)

	result, err := r.uc.Auth.LoginDashboard(ctx, paramBody, paramHeader)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}
