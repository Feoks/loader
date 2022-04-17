package infrastructure

import (
	"context"
	"github.com/Feoks/loader/configs"
	"github.com/Feoks/loader/tools/metrics"
)

type metric struct{}

func (*metric) init(ctx context.Context, cfg *configs.Config) (context.Context, error) {
	return metrics.WithContext(ctx), nil
}

func (*metric) closer() {}
