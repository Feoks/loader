package infrastructure

import (
	"context"
	"github.com/Feoks/loader/configs"
	s "github.com/Feoks/loader/tools/sentry"
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
