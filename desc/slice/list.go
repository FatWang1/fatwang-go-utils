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

// 移除list中的元素 不改变元素顺序
func RemoveListElement[T comparable](list []T, val T) []T {
	var newList = make([]T, len(list), len(list))
	left := 0
	for _, v := range list {
		if v != val {
			newList[left] = v
			left++
		}
	}
	return newList[0:left]
}

// 移除list中的元素 原地交换省内存性能好 改变元素顺序
func RemoveListElementInPlace[T comparable](list []T, val T) []T {
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
