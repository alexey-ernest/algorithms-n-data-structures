package ch4

type depthFirstPaths struct{
	adj [][]int
	source int
	marked []bool
	edgeTo []int
}

func NewDepthFirstPaths(adj [][]int, source int) *depthFirstPaths {
	marked := make([]bool, len(adj))
	edgeTo := make([]int, len(adj))
	for i := range edgeTo {
		edgeTo[i] = -1
	}

	dfp := &depthFirstPaths{
		adj: adj,
		source: source,
		marked: marked,
		edgeTo: edgeTo,
	}

	dfp.dfs(source)

	return dfp
}

func (this *depthFirstPaths) dfs(v int) {
	this.marked[v] = true
	for _, w := range this.adj[v] {
		if this.marked[w] {
			continue
		}
		this.edgeTo[w] = v
		this.dfs(w)
	}
}

func (this *depthFirstPaths) HasPathTo(v int) bool {
	return this.marked[v]
}

func (this *depthFirstPaths) PathTo(v int) []int {
	if v < 0 {
		panic("v should be >= 0")
	}

	path := make([]int, 0)
	for w := v; w >= 0; w = this.edgeTo[w] {
		path = append(path, w)
	}

	res := make([]int, len(path))
	for i := range path {
		res[len(res)-i-1] = path[i]
	}
	return res
}
