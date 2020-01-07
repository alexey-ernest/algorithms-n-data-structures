package ch4

type cc struct {
	adj [][]int
	id []int
	count int
	marked []bool
}

func NewConnectedComponents(adj [][]int) *cc {
	id := make([]int, len(adj))
	marked := make([]bool, len(adj))
	comp := &cc{
		adj: adj,
		id: id,
		marked: marked,
	}

	for i := 0; i < len(adj); i += 1 {
		if marked[i] {
			continue
		}

		comp.dfs(i)
		comp.count++
	}

	return comp
}

func (this *cc) dfs(v int) {
	this.marked[v] = true
	this.id[v] = this.count

	for _, w := range this.adj[v] {
		if this.marked[w] {
			continue
		}
		this.dfs(w)
	}
}

func (this *cc) Connected(v, w int) bool {
	return this.id[v] == this.id[w]
}

func (this *cc) Count() int {
	return this.count
}

