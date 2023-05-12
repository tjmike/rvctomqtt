package utils

import (
	"context"
	"go.uber.org/zap"
)

func ApplyContext(ctx *context.Context, log *zap.Logger) *zap.Logger {
	var fields = (*ctx).Value("logFields")
	if fields != nil {
		var m, ok = fields.([]zap.Field)
		if ok {
			log = log.With(m...)
		}
	}
	return log
}
