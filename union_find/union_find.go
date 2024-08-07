package union_find

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func InitUnionFind(n int) *UnionFind {
	// 初始时，总共有n个分组
	uf := &UnionFind{count: n}
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := 0; i < n; i++ {
		// 初始时，每个元素单独一个分组
		uf.parent[i] = i
		// 每个分组里只有一个元素
		uf.rank[i] = 1
	}
	return uf
}

// Find 递归查找元素x的根结点，查找的同时将该组所有元素(遍历到的)都直接指向根节点（路径压缩）
func (u *UnionFind) Find(x int) int {
	if u.parent[x] == x {
		return x
	}
	u.parent[x] = u.Find(u.parent[x])
	return u.parent[x]
}

//func (u *UnionFind) Find1(x int) int {
//	for x != u.parent[x] {
//		u.parent[x] = u.parent[u.parent[x]]
//		x = u.parent[x]
//	}
//	return x
//}

// Union 合并两个分组
func (u *UnionFind) Union(x, y int) {
	xp := u.Find(x)
	yp := u.Find(y)
	if xp == yp {
		// 已经是同一个分组了，直接返回
		return
	}
	// 我们将小分组合并到大分组（这一步不是必须的）
	if u.rank[yp] > u.rank[xp] {
		xp, yp = yp, xp
	}
	// 大分组的元素数量增加
	u.rank[xp] += u.rank[yp]
	// 小分组消失，让元素数量变0
	u.rank[yp] = 0
	// 合并只需要小分组的根指向大分组任意一个元素即可
	// 这里我们让小分组的根指向大分组的根
	u.parent[yp] = xp
	// 总的分组数减少
	u.count--
}
