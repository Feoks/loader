package infrastructure

import (
	"context"
	"fmt"
	"git.repo.services.lenvendo.ru/grade-factor/echo/configs"
	"git.repo.services.lenvendo.ru/grade-factor/echo/tools/tracing"
	"github.com/pkg/errors"
	"io"
)

type tracer struct {
	enable       bool
	tracerCloser io.Closer
}

//execute если трассировщик включен, то кладем его в контекст
//возвращаем его в контексте и возвращаем io.Closer(), который необходимо закрыть в main
func (t *tracer) init(ctx context.Context, cfg *configs.Config) (context.Context, error) {
	t.enable = cfg.Tracer.Enabled
	if !t.enable {
		return ctx, nil
	}

	tracer, closer, err := tracing.NewJaegerTracer(
		ctx,
		fmt.Sprintf("%s:%d", cfg.Tracer.Host, cfg.Tracer.Port),
		cfg.Tracer.Name,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init tracer")
	}

	t.tracerCloser = closer
	return tracing.WithContext(ctx, tracer), nil
}

func (t *tracer) closer() {
	if t.enable {
		t.tracerCloser.Close()
	}
}
