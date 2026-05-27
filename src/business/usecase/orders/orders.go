package orders

import (
	"context"
	"fmt"
	"time"

	ordersDom "github.com/alpardfm/e-commerce/src/business/domain/orders"
	paymentsDom "github.com/alpardfm/e-commerce/src/business/domain/payments"
	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/e-commerce/src/utils/config"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/alpardfm/go-toolkit/errors"
	"github.com/alpardfm/go-toolkit/log"
	"github.com/google/uuid"
)

type Interface interface {
	Create(ctx context.Context, param CreateOrderInput, createdBy string) (entity.Orders, error)
	GetListByUser(ctx context.Context, userID int64) ([]entity.Orders, error)
	UpdateStatus(ctx context.Context, orderID int64, status string, updatedBy string) (entity.Orders, error)
}

type CreateOrderInput struct {
	UserID        int64   `json:"user_id"`
	TotalPrice    float64 `json:"total_price"`
	PaymentMethod string  `json:"payment_method"`
}

type ordersUC struct {
	log      log.Interface
	cfg      config.Application
	orderDom ordersDom.Interface
	payDom   paymentsDom.Interface
}

func Init(log log.Interface, cfg config.Application, orderDom ordersDom.Interface, payDom paymentsDom.Interface) Interface {
	return &ordersUC{
		log:      log,
		cfg:      cfg,
		orderDom: orderDom,
		payDom:   payDom,
	}
}

func (o *ordersUC) Create(ctx context.Context, param CreateOrderInput, createdBy string) (entity.Orders, error) {
	if param.UserID <= 0 {
		return entity.Orders{}, errors.NewWithCode(codes.CodeInvalidValue, "user_id is required")
	}
	if param.TotalPrice <= 0 {
		return entity.Orders{}, errors.NewWithCode(codes.CodeInvalidValue, "total_price must be greater than 0")
	}
	if param.PaymentMethod == "" {
		return entity.Orders{}, errors.NewWithCode(codes.CodeInvalidValue, "payment_method is required")
	}

	now := time.Now().UTC()

	// Create order
	order := entity.Orders{
		UserID:     param.UserID,
		TotalPrice: param.TotalPrice,
		Status:     "pending",
		CreatedAt:  now,
		CreatedBy:  createdBy,
		IsDeleted:  0,
	}

	createdOrder, err := o.orderDom.Create(ctx, order)
	if err != nil {
		return entity.Orders{}, err
	}

	// Create payment record
	payment := entity.Payments{
		OrderID:       createdOrder.ID,
		PaymentMethod: param.PaymentMethod,
		PaymentStatus: "pending",
		TransactionID: uuid.New().String(),
		CreatedAt:     now,
		CreatedBy:     createdBy,
		IsDeleted:     0,
	}

	if _, err := o.payDom.Create(ctx, payment); err != nil {
		return entity.Orders{}, err
	}

	return createdOrder, nil
}

func (o *ordersUC) GetListByUser(ctx context.Context, userID int64) ([]entity.Orders, error) {
	results, err := o.orderDom.GetList(ctx, entity.Orders{UserID: userID}, func(_, suffix *string) error {
		*suffix = fmt.Sprintf("AND is_deleted = %d AND user_id = %d", 0, userID)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (o *ordersUC) UpdateStatus(ctx context.Context, orderID int64, status string, updatedBy string) (entity.Orders, error) {
	if orderID <= 0 {
		return entity.Orders{}, errors.NewWithCode(codes.CodeInvalidValue, "order_id is required")
	}

	validStatuses := map[string]bool{
		"pending": true, "paid": true, "shipped": true, "completed": true, "canceled": true,
	}
	if !validStatuses[status] {
		return entity.Orders{}, errors.NewWithCode(codes.CodeInvalidValue, "invalid order status")
	}

	result, err := o.orderDom.Update(ctx, entity.Orders{
		ID:        orderID,
		Status:    status,
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: updatedBy,
	})
	if err != nil {
		return entity.Orders{}, err
	}

	return result, nil
}
