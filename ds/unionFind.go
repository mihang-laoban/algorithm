package ds

type UnionFind struct {
	Count  int
	Parent []int
}

type UF interface {
	UnionFind(int)
	Find(int) int
	Union(int, int)
}

func (root *UnionFind) UnionFind(a int) {
	root.Count = a
	root.Parent = make([]int, a)
	for i := 0; i < a; i++ {
		root.Parent[i] = i
	}
}

func (root *UnionFind) Find(a int) int {
	for a != root.Parent[a] {
		root.Parent[a] = root.Parent[root.Parent[a]]
		a = root.Parent[a]
	}
	return a
}

func (root *UnionFind) Union(a, b int) {
	rootA := root.Find(a)
	rootB := root.Find(b)
	if rootA == rootB {
		return
	}
	root.Parent[rootA] = rootB
	root.Count--
}

func union(p []int, i, j int) {
	p1 := parent(p, i)
	p2 := parent(p, j)
	p[p2] = p1
}

func parent(p []int, i int) int {
	root := i
	for p[root] != root {
		root = p[root]
	}
	for p[i] != i {
		x := i
		i = p[i]
		p[x] = root
	}
	return root
}
