package domain

import (
	"context"

	"git.repo.services.lenvendo.ru/grade-factor/echo/pkg/health"
)

const HEALTH_PATH = ""

type healthService interface {
	initHealthService(context.Context)
	getHealth() health.Service
}

func (s *service) initHealthService(ctx context.Context) {
	healthService := health.NewHealthService()
	if s.infra.Configs().Metrics.Enabled {
		healthService = health.NewMetricsService(ctx, healthService)
	}
	healthService = health.NewLoggingService(ctx, healthService)

	if s.infra.Configs().Sentry.Enabled {
		healthService = health.NewSentryService(healthService)
	}
	s.healthService = healthService
}

func (s *service) getHealth() health.Service {
	return s.healthService
}
