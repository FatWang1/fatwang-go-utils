package utils

import "gopkg.in/errgo.v2/errors"

var (
	ErrOutOfRange = errors.New("index out of range")
)

func Contain[T comparable](list []T, item T) bool {
	for _, t := range list {
		if t == item {
			return true
		}
	}
	return false
}

func GetItem[T any](idx int, list []T) (T, error) {
	var t T
	if len(list) <= idx || idx < 0 {
		return t, ErrOutOfRange
	}
	return list[idx], nil
}

// 移除list中的元素 原地交换省内存性能好
func RemoveItemByValue[T comparable](list []T, val T) []T {
	left, right := 0, 0
	for right < len(list) {
		if list[right] != val {
			list[left] = list[right]
			left++
		}
		right++
	}
	return list[:left]
}

func InsertItems[T comparable](list []T, idx int, vals ...T) ([]T, error) {
	if idx < 0 || idx > len(list) {
		return nil, ErrOutOfRange
	}
	list = append(list, vals...)
	copy(list[idx+len(vals):], list[idx:])
	copy(list[idx:], vals)
	return list, nil
}
