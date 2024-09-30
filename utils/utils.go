package utils

import (
	"context"
	"strings"

	"github.com/openzipkin/zipkin-go"
	uuid "github.com/satori/go.uuid"
)

func GetUuid() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

func TernaryOperator[T any](expr bool, trueVal, falseVal T) T {
	if expr {
		return trueVal
	} else {
		return falseVal
	}
}

func ContextCopy(ctx context.Context) context.Context {
	return zipkin.NewContext(context.Background(), zipkin.SpanOrNoopFromContext(ctx))
}
