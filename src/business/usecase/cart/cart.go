package cart

import (
	"context"
	"fmt"
	"time"

	cartDom "github.com/alpardfm/e-commerce/src/business/domain/cart"
	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/e-commerce/src/utils/config"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/alpardfm/go-toolkit/errors"
	"github.com/alpardfm/go-toolkit/log"
)

type Interface interface {
	GetList(ctx context.Context, userID int64) ([]entity.Cart, error)
	Create(ctx context.Context, param entity.Cart, createdBy string) (entity.Cart, error)
	Update(ctx context.Context, param entity.Cart, updatedBy string) (entity.Cart, error)
	Delete(ctx context.Context, param entity.Cart, deletedBy string) (entity.Cart, error)
}

type cartUC struct {
	log log.Interface
	cfg config.Application
	dom cartDom.Interface
}

func Init(log log.Interface, cfg config.Application, dom cartDom.Interface) Interface {
	return &cartUC{
		log: log,
		cfg: cfg,
		dom: dom,
	}
}

func (c *cartUC) GetList(ctx context.Context, userID int64) ([]entity.Cart, error) {
	results, err := c.dom.GetList(ctx, entity.Cart{UserID: userID}, func(_, suffix *string) error {
		*suffix = fmt.Sprintf("AND is_deleted = %d AND user_id = %d", 0, userID)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (c *cartUC) Create(ctx context.Context, param entity.Cart, createdBy string) (entity.Cart, error) {
	if param.ProductID <= 0 {
		return entity.Cart{}, errors.NewWithCode(codes.CodeInvalidValue, "product_id is required")
	}
	if param.Quantity <= 0 {
		return entity.Cart{}, errors.NewWithCode(codes.CodeInvalidValue, "quantity must be greater than 0")
	}

	param.CreatedAt = time.Now().UTC()
	param.CreatedBy = createdBy
	param.IsDeleted = 0

	result, err := c.dom.Create(ctx, param)
	if err != nil {
		return entity.Cart{}, err
	}

	return result, nil
}

func (c *cartUC) Update(ctx context.Context, param entity.Cart, updatedBy string) (entity.Cart, error) {
	if param.ID <= 0 {
		return entity.Cart{}, errors.NewWithCode(codes.CodeInvalidValue, "cart id is required")
	}
	if param.Quantity <= 0 {
		return entity.Cart{}, errors.NewWithCode(codes.CodeInvalidValue, "quantity must be greater than 0")
	}

	param.UpdatedAt = time.Now().UTC()
	param.UpdatedBy = updatedBy

	result, err := c.dom.Update(ctx, param)
	if err != nil {
		return entity.Cart{}, err
	}

	return result, nil
}

func (c *cartUC) Delete(ctx context.Context, param entity.Cart, deletedBy string) (entity.Cart, error) {
	if param.ID <= 0 {
		return entity.Cart{}, errors.NewWithCode(codes.CodeInvalidValue, "cart id is required")
	}

	param.DeletedAt = time.Now().UTC()
	param.DeletedBy = deletedBy
	param.IsDeleted = 1

	result, err := c.dom.Delete(ctx, param)
	if err != nil {
		return entity.Cart{}, err
	}

	return result, nil
}
