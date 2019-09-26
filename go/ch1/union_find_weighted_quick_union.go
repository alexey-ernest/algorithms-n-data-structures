package ch1

type unionFindWeightedQuickUnion struct {
	data []int
	n int
	sizes []int
}

func NewUnionFindWeightedQuickUnion(n int) unionFindWeightedQuickUnion {
	d := make([]int, n)
	s := make([]int, n)
	for i := range d {
		d[i] = i
		s[i] = 1
	}

	return unionFindWeightedQuickUnion {
		data: d,
		n: n,
		sizes: s,
	}
}

func (uf *unionFindWeightedQuickUnion) count() int {
	return uf.n
}

func (uf *unionFindWeightedQuickUnion) find(p int) int {
	for p != uf.data[p] {
		p = uf.data[p]
	}
	return p
}

func (uf *unionFindWeightedQuickUnion) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *unionFindWeightedQuickUnion) union(p, q int) {
	pid := uf.find(p)
	qid := uf.find(q)
	if pid == qid {
		return
	}

	if uf.sizes[pid] < uf.sizes[qid] {
		// merging p to q
		uf.data[pid] = qid
		uf.sizes[qid] += uf.sizes[pid]
	} else {
		// merging q to p
		uf.data[qid] = pid
		uf.sizes[pid] += uf.sizes[qid]
	}
	uf.n -= 1
}