package utils

import (
	"context"
	"math/rand"
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

func GenerateRandomIntList(size int, edge1, edge2 int) []int {
	list := make([]int, size, size)
	if edge1 > edge2 {
		edge1, edge2 = edge2, edge1
	}
	for i := 0; i < size; i++ {
		list[i] = rand.Intn(edge2-edge1) + edge1
	}
	return list
}
