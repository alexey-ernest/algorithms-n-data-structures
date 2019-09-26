package ch1

type unionFindQuickUnion struct {
	data []int
	n int
}

func NewUnionFindQuickUnion(n int) unionFindQuickUnion {
	d := make([]int, n)
	for i := range d {
		d[i] = i
	}

	return unionFindQuickUnion {
		data: d,
		n: n,
	}
}

func (uf *unionFindQuickUnion) count() int {
	return uf.n
}

func (uf *unionFindQuickUnion) find(p int) int {
	for p != uf.data[p] {
		p = uf.data[p]
	}
	return p
}

func (uf *unionFindQuickUnion) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *unionFindQuickUnion) union(p, q int) {
	pid := uf.find(p)
	qid := uf.find(q)
	if pid == qid {
		return
	}

	// assigning first tree to the second one's root
	uf.data[pid] = qid
	uf.n -= 1
}