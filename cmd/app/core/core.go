package core

import (
	"context"
	"git.repo.services.lenvendo.ru/grade-factor/echo/cmd/app/core/domain"
	"git.repo.services.lenvendo.ru/grade-factor/echo/cmd/app/core/infrastructure"
	"git.repo.services.lenvendo.ru/grade-factor/echo/configs"
)

type ServiceCore interface {
	Service() domain.EchoService
}

type serviceCore struct {
	service domain.EchoService
}

type InfrastructureCore interface {
	Cfg() *configs.Config
	Ctx() context.Context
	Closer()
}

type infraCore struct {
	infrastructure.Infrastructure
}

func InitCore(ctx context.Context) (ServiceCore, InfrastructureCore) {
	infra := infraCore{infrastructure.NewInfrastructure(ctx)}
	service := serviceCore{domain.NewDomain(ctx, infra)}

	return &service, &infra
}

func (c *serviceCore) Service() domain.EchoService {
	return c.service
}

func (c *infraCore) Ctx() context.Context {
	return c.Context()
}

func (c *infraCore) Cfg() *configs.Config {
	return c.Configs()
}

func (c *infraCore) Closer() {
	c.AllCloser()
}
