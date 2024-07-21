package domain

import (
	"github.com/alpardfm/e-commerce/src/business/domain/cart"
	"github.com/alpardfm/e-commerce/src/business/domain/categories"
	"github.com/alpardfm/e-commerce/src/business/domain/location"
	"github.com/alpardfm/e-commerce/src/business/domain/order_items"
	"github.com/alpardfm/e-commerce/src/business/domain/orders"
	"github.com/alpardfm/e-commerce/src/business/domain/otp"
	"github.com/alpardfm/e-commerce/src/business/domain/payments"
	"github.com/alpardfm/e-commerce/src/business/domain/products"
	"github.com/alpardfm/e-commerce/src/business/domain/refund"
	"github.com/alpardfm/e-commerce/src/business/domain/reviews"
	"github.com/alpardfm/e-commerce/src/business/domain/role"
	"github.com/alpardfm/e-commerce/src/business/domain/users"
	"github.com/alpardfm/e-commerce/src/utils/config"
	"github.com/alpardfm/go-toolkit/log"
	"github.com/alpardfm/go-toolkit/parser"
	"github.com/alpardfm/go-toolkit/sql"
)

type Domains struct {
	Users      users.Interface
	Cart       cart.Interface
	Categories categories.Interface
	Location   location.Interface
	OrderItems order_items.Interface
	Orders     orders.Interface
	Otp        otp.Interface
	Payments   payments.Interface
	Products   products.Interface
	Refund     refund.Interface
	Reviews    reviews.Interface
	Role       role.Interface
}

func Init(log log.Interface, db sql.Interface, parser parser.JSONInterface, cfg config.Application) *Domains {
	return &Domains{
		Users:      users.Init(log, db),
		Cart:       cart.Init(log, db),
		Categories: categories.Init(log, db),
		Location:   location.Init(log, db),
		OrderItems: order_items.Init(log, db),
		Orders:     orders.Init(log, db),
		Otp:        otp.Init(log, db),
		Payments:   payments.Init(log, db),
		Products:   products.Init(log, db),
		Refund:     refund.Init(log, db),
		Reviews:    reviews.Init(log, db),
		Role:       role.Init(log, db),
	}
}
