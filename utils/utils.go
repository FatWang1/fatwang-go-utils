package utils

import (
	"strings"

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
