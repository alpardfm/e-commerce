package usecase

import (
	"github.com/alpardfm/e-commerce/src/business/domain"
	"github.com/alpardfm/e-commerce/src/business/usecase/auth"
	"github.com/alpardfm/e-commerce/src/business/usecase/categories"
	"github.com/alpardfm/e-commerce/src/business/usecase/location"
	"github.com/alpardfm/e-commerce/src/business/usecase/role"
	"github.com/alpardfm/e-commerce/src/utils/config"
	"github.com/alpardfm/go-toolkit/log"
	"github.com/alpardfm/go-toolkit/parser"
)

type Usecases struct {
	Categories categories.Interface
	Location   location.Interface
	Role       role.Interface
	Auth       auth.Interface
}

func Init(log log.Interface, d *domain.Domains, jsonParser parser.JSONInterface, cfg config.Application) *Usecases {
	return &Usecases{
		Categories: categories.Init(log, cfg, d.Categories, d.Role),
		Location:   location.Init(log, cfg, d.Location, d.Role),
		Role:       role.Init(log, cfg, d.Role),
		Auth:       auth.Init(log, cfg, d.Users, d.Location, d.Role),
	}
}
