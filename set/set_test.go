package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := InitSet[int](12)
	s.HasKey()
	for i := 0; i < 12; i++ {
		s.Set(i + 1)
	}
	for i := 0; i < 6; i++ {
		s.Drop(i*2 + 1)
	}
	t.Logf("s has %d keys", s.Len())
	for i := 0; i < 12; i++ {
		t.Logf("s has %d : %v", i+1, s.HasKey(i+1))
	}
	t.Log(s)
	s = s.DropAll()
	t.Logf("s has %d keys, s is\n %v", s.Len(), s)
	s.Set(1, 2, 3)
	Setify[int](s.ToSlice()...)
	s.HasAny(1)
	s.HasAny(4)
}
