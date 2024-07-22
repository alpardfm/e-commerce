package rest

import (
	"strconv"

	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/gin-gonic/gin"
)

func (r *rest) GetListLocationDashboard(ctx *gin.Context) {
	tokens := ctx.GetHeader("Authorization")
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	paginate := entity.PaginationLocation{}
	param := entity.Location{}

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

	result, err := r.uc.Location.GetListDashboard(ctx, param, paginate, tokens)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) GetDetailLocation(ctx *gin.Context) {
	tokens := ctx.GetHeader("Authorization")
	id := ctx.Param("id")

	param := entity.Location{}

	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			r.httpRespError(ctx, err)
			return
		}

		param.ID = int64(idInt)
	}

	result, err := r.uc.Location.GetDetail(ctx, param, tokens)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) CreateLocation(ctx *gin.Context) {
	tokens := ctx.GetHeader("Authorization")
	var body entity.BodyLocation
	ctx.Bind(&body)

	result, err := r.uc.Location.Create(ctx, entity.Location{
		Lat:      body.Lat,
		Long:     body.Long,
		Distance: body.Distance,
	}, tokens)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) UpdateLocation(ctx *gin.Context) {
	tokens := ctx.GetHeader("Authorization")
	id := ctx.Param("id")

	param := entity.Location{}
	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			r.httpRespError(ctx, err)
			return
		}

		param.ID = int64(idInt)
	}

	var body entity.BodyLocation
	ctx.Bind(&body)
	param.Lat = body.Lat
	param.Long = body.Long
	param.Distance = body.Distance

	result, err := r.uc.Location.Update(ctx, param, tokens)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}

func (r *rest) DeleteLocation(ctx *gin.Context) {
	tokens := ctx.GetHeader("Authorization")
	id := ctx.Param("id")

	param := entity.Location{}

	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			r.httpRespError(ctx, err)
			return
		}

		param.ID = int64(idInt)
	}

	result, err := r.uc.Location.Delete(ctx, param, tokens)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, result, nil)
}
