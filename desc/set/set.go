package set

type Set[K comparable] interface {
	Set(keys ...K)
	HasKey(keys ...K) bool
	HasAny(keys ...K) bool
	Drop(key ...K)
	Len() int
	DropAll() Set[K]
	ToSlice() []K
}

type set[K comparable] map[K]struct{}

func (s set[K]) Set(keys ...K) {
	for _, key := range keys {
		s[key] = struct{}{}
	}
}

// 是否包含全部key
func (s set[K]) HasKey(keys ...K) bool {
	if len(keys) == 0 {
		return false
	}

	for _, k := range keys {
		if _, ok := s[k]; !ok {
			return false
		}
	}
	return true
}

func (s set[K]) Drop(keys ...K) {
	for _, k := range keys {
		delete(s, k)
	}
}

func (s set[K]) Len() int {
	return len(s)
}

func (s set[K]) DropAll() Set[K] {
	return make(set[K])
}

func InitSet[K comparable](length int) Set[K] {
	return make(set[K], length)
}

func Setify[K comparable](keys ...K) Set[K] {
	s := InitSet[K](len(keys))
	s.Set(keys...)
	return s
}

func (s set[K]) ToSlice() []K {
	sl := make([]K, 0, s.Len())
	for k := range s {
		sl = append(sl, k)
	}
	return sl
}

func (s set[K]) HasAny(keys ...K) bool {
	for _, key := range keys {
		if s.HasKey(key) {
			return true
		}
	}
	return false
}
