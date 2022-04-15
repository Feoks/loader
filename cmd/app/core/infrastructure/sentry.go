package infrastructure

import (
	"context"
	"git.repo.services.lenvendo.ru/grade-factor/echo/configs"
	s "git.repo.services.lenvendo.ru/grade-factor/echo/tools/sentry"
	"github.com/pkg/errors"
)

type sentry struct{}

func (*sentry) init(ctx context.Context, cfg *configs.Config) (context.Context, error) {
	if err := s.NewSentry(cfg); err != nil {
		return nil, errors.Wrap(err, "failed to init sentry")
	}

	return ctx, nil
}

func (*sentry) closer() {}
