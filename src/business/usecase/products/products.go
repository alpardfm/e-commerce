package products

import (
	"context"
	"fmt"
	"time"

	productsDom "github.com/alpardfm/e-commerce/src/business/domain/products"
	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/e-commerce/src/utils/config"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/alpardfm/go-toolkit/errors"
	"github.com/alpardfm/go-toolkit/log"
)

type Interface interface {
	GetList(ctx context.Context, param entity.Products, paginate entity.PaginationProducts) (entity.ResponseProducts, error)
	GetDetail(ctx context.Context, param entity.Products) (entity.Products, error)
	Create(ctx context.Context, param entity.Products, createdBy string) (entity.Products, error)
	Update(ctx context.Context, param entity.Products, updatedBy string) (entity.Products, error)
	Delete(ctx context.Context, param entity.Products, deletedBy string) (entity.Products, error)
}

type products struct {
	log log.Interface
	cfg config.Application
	dom productsDom.Interface
}

func Init(log log.Interface, cfg config.Application, dom productsDom.Interface) Interface {
	return &products{
		log: log,
		cfg: cfg,
		dom: dom,
	}
}

func (p *products) GetList(ctx context.Context, param entity.Products, paginate entity.PaginationProducts) (entity.ResponseProducts, error) {
	results, err := p.dom.GetList(ctx, param, func(_, suffix *string) error {
		*suffix = fmt.Sprintf("AND is_deleted = %d", 0)
		return nil
	})
	if err != nil {
		return entity.ResponseProducts{}, err
	}

	// Calculate pagination
	totalRows := int64(len(results))
	var totalPages int64
	if totalRows != 0 && paginate.Limit != 0 {
		totalPages = (totalRows + paginate.Limit - 1) / paginate.Limit
	}

	// Apply pagination
	start := (paginate.Page - 1) * paginate.Limit
	end := start + paginate.Limit
	if start > totalRows {
		start = totalRows
	}
	if end > totalRows {
		end = totalRows
	}

	return entity.ResponseProducts{
		Limit:      paginate.Limit,
		Page:       paginate.Page,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       results[start:end],
	}, nil
}

func (p *products) GetDetail(ctx context.Context, param entity.Products) (entity.Products, error) {
	result, err := p.dom.GetDetail(ctx, param, func(_, suffix *string) error {
		*suffix = fmt.Sprintf("AND is_deleted = %d", 0)
		return nil
	})
	if err != nil {
		return entity.Products{}, err
	}
	return result, nil
}

func (p *products) Create(ctx context.Context, param entity.Products, createdBy string) (entity.Products, error) {
	if param.Name == "" {
		return entity.Products{}, errors.NewWithCode(codes.CodeInvalidValue, "product name is required")
	}
	if param.Price <= 0 {
		return entity.Products{}, errors.NewWithCode(codes.CodeInvalidValue, "product price must be greater than 0")
	}
	if param.Stock < 0 {
		return entity.Products{}, errors.NewWithCode(codes.CodeInvalidValue, "product stock cannot be negative")
	}

	param.CreatedAt = time.Now().UTC()
	param.CreatedBy = createdBy
	param.IsDeleted = 0

	result, err := p.dom.Create(ctx, param)
	if err != nil {
		return entity.Products{}, err
	}

	return result, nil
}

func (p *products) Update(ctx context.Context, param entity.Products, updatedBy string) (entity.Products, error) {
	if param.ID <= 0 {
		return entity.Products{}, errors.NewWithCode(codes.CodeInvalidValue, "product id is required")
	}
	if param.Name == "" {
		return entity.Products{}, errors.NewWithCode(codes.CodeInvalidValue, "product name is required")
	}

	param.UpdatedAt = time.Now().UTC()
	param.UpdatedBy = updatedBy

	result, err := p.dom.Update(ctx, param)
	if err != nil {
		return entity.Products{}, err
	}

	return result, nil
}

func (p *products) Delete(ctx context.Context, param entity.Products, deletedBy string) (entity.Products, error) {
	if param.ID <= 0 {
		return entity.Products{}, errors.NewWithCode(codes.CodeInvalidValue, "product id is required")
	}

	param.DeletedAt = time.Now().UTC()
	param.DeletedBy = deletedBy
	param.IsDeleted = 1

	result, err := p.dom.Delete(ctx, param)
	if err != nil {
		return entity.Products{}, err
	}

	return result, nil
}
