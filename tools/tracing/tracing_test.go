// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
package tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromContext(t *testing.T) {
	ctx := context.Background()
	tracer := FromContext(ctx)
	assert.Equal(t, opentracing.NoopTracer{}, tracer)
}

func TestWithContext(t *testing.T) {
	ctx := context.Background()
	expected := opentracing.NoopTracer{}
	ctx = WithContext(ctx, expected)
	actual := FromContext(ctx)
	assert.Equal(t, expected, actual)
}