package slice

func Contain[T comparable](list []T, item T) bool {
	for _, t := range list {
		if t == item {
			return true
		}
	}
	return false
}

func GetItem[T any](idx int, list []T) T {
	var t T
	if len(list) <= idx || idx < 0 {
		return t
	}
	return list[idx]
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

func InsertItems[T comparable](list []T, idx int, vals ...T) []T {
	if idx < 0 || idx > len(list) {
		return list
	}
	if idx == len(list) {
		return append(list, vals...)
	}
	length := len(list) + len(vals)
	if cap(list) > length {
		list = list[:length]
		copy(list[idx+len(vals):], list[idx:])
		copy(list[idx:], vals)
	} else {
		newList := make([]T, length)
		copy(newList, list[:idx])
		copy(newList[idx:], vals)
		copy(newList[idx+len(vals):], list[idx:])
		list = newList
	}
	return list
}
