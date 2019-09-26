package ch1

type unionFindQuickFind struct {
	data []int
	n int
}

func NewUnionFindQuickFind(n int) unionFindQuickFind {
	d := make([]int, n)
	for i := range d {
		d[i] = i
	}

	return unionFindQuickFind {
		data: d,
		n: n,
	}
}

func (uf *unionFindQuickFind) find(p int) int {
	return uf.data[p]
}

func (uf *unionFindQuickFind) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *unionFindQuickFind) count() int {
	return uf.n
}

func (uf *unionFindQuickFind) union(p, q int) {
	pid := uf.find(p)
	qid := uf.find(q)
	if pid == qid {
		return
	}

	for i := range uf.data {
		if uf.data[i] == pid {
			uf.data[i] = qid
		}
	}
	uf.n -= 1
}
