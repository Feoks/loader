package infrastructure

import (
	"context"
	"git.repo.services.lenvendo.ru/grade-factor/echo/configs"
	"git.repo.services.lenvendo.ru/grade-factor/echo/tools/logging"
	"github.com/pkg/errors"
)

type logger struct{}

func (*logger) init(ctx context.Context, cfg *configs.Config) (context.Context, error) {
	l, err := logging.NewLogger(cfg.Logger.Level, cfg.Logger.TimeFormat)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init logger: %s")
	}

	return logging.WithContext(ctx, l), nil
}

func (*logger) closer() {}
