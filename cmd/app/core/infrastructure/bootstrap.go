package infrastructure

import (
	"context"
	"git.repo.services.lenvendo.ru/grade-factor/echo/configs"
)

//initInfrastructure init - инициализация юнитов
//closer - закрытие, если того требует юнит
type initInfrastructure interface {
	init(ctx context.Context, cfg *configs.Config) (context.Context, error)
	closer()
}

type units []initInfrastructure
type bootstrap struct {
	units
}

//infra добавление структуры с реализацией initInfrastructure
func (b *bootstrap) initBootstrap(cfg *configs.Config) (boot []initInfrastructure) {
	boot = append(boot, &logger{})
	if !cfg.Metrics.Enabled {
		boot = append(boot, &metric{})
	}

	if cfg.Sentry.Enabled {
		boot = append(boot, &sentry{})
	}

	if cfg.Tracer.Enabled {
		boot = append(boot, &tracer{})
	}
	return boot
}
