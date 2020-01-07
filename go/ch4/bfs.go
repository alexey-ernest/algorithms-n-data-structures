package ch4

type breadthFirstPaths struct {
	adj [][]int
	source int
	marked []bool
	edgeTo []int
}

func NewBreadthFirstPaths(adj [][]int, source int) *breadthFirstPaths {
	marked := make([]bool, len(adj))
	edgeTo := make([]int, len(adj))
	for i := range edgeTo {
		edgeTo[i] = -1
	}

	bfs := &breadthFirstPaths{
		adj: adj,
		source: source,
		marked: marked,
		edgeTo: edgeTo,
	}

	bfs.bfs(source)

	return bfs
}

func (this *breadthFirstPaths) bfs(v int) {
	q := &Queue{}
	q.Enqueue(v)
	this.marked[v] = true
	
	for !q.IsEmpty() {
		v := q.Dequeue()
		for _, w := range this.adj[v] {
			if this.marked[w] {
				continue
			}
			this.edgeTo[w] = v
			this.marked[w] = true
			q.Enqueue(w)
		}
	}
}

func (this *breadthFirstPaths) HasPathTo(v int) bool {
	return this.marked[v]
}

func (this *breadthFirstPaths) PathTo(v int) []int {
	path := make([]int, 0)
	for w := v; w >= 0; w = this.edgeTo[w] {
		path = append(path, w)
	}

	res := make([]int, len(path))
	for i := range path {
		res[len(path)-i-1] = path[i]
	}
	return res
}