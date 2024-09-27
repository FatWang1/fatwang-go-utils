package union_find

import (
	"testing"
)

// todo: 有缘就补上
func TestUnionFind(t *testing.T) {
	n := 10
	uf := InitUnionFind(n)

	// Test initial state
	if uf.count != n {
		t.Errorf("Expected initial count to be %d, got %d", n, uf.count)
	}
	for i := 0; i < n; i++ {
		if uf.Find(i) != i {
			t.Errorf("Expected element %d to be its own root", i)
		}
	}

	// Test union and find
	uf.Union(0, 1)
	if uf.Find(0) != uf.Find(1) {
		t.Errorf("Elements 0 and 1 should be in the same group")
	}
	if uf.count != n-1 {
		t.Errorf("Expected count to be %d, got %d", n-1, uf.count)
	}
	uf.Union(1, 1)

	// Test rank
	uf.Union(2, 0)
	if uf.rank[uf.Find(0)] != 3 {
		t.Errorf("Expected rank of group containing element 0 to be 3, got %d", uf.rank[uf.Find(0)])
	}
	//if uf.rank[uf.Find1(0)] != 3 {
	//	t.Errorf("Expected rank of group containing element 0 to be 3, got %d", uf.rank[uf.Find1(0)])
	//}
}
