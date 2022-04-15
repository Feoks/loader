package infrastructure

import (
	"context"
	"git.repo.services.lenvendo.ru/grade-factor/echo/configs"
	"git.repo.services.lenvendo.ru/grade-factor/echo/tools/metrics"
)

type metric struct{}

func (*metric) init(ctx context.Context, cfg *configs.Config) (context.Context, error) {
	return metrics.WithContext(ctx), nil
}

func (*metric) closer() {}
