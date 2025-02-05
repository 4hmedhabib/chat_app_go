package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey int

const (
	writeKey ctxKey = iota + 1
	traceIDKey
)

func SetTraceID(ctx context.Context, traceID uuid.UUID) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func GetTraceID(ctx context.Context) uuid.UUID {
	v, ok := ctx.Value(traceIDKey).(uuid.UUID)
	if !ok {
		return uuid.UUID{}
	}

	return v
}

func setWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, writeKey, w)
}

func GetWriter(ctx context.Context) http.ResponseWriter {
	v, ok := ctx.Value(writeKey).(http.ResponseWriter)
	if !ok {
		return nil
	}

	return v
}
