package ch4

type depthFirstPaths struct{
	adj [][]int
	source int
	marked []bool
	pathTo []int
}

func NewDepthFirstPaths(adj [][]int, source int) *depthFirstPaths {
	marked := make([]bool, len(adj))
	pathTo := make([]int, len(adj))
	for i := range pathTo {
		pathTo[i] = -1
	}

	dfp := &depthFirstPaths{
		adj: adj,
		source: source,
		marked: marked,
		pathTo: pathTo,
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
		this.pathTo[w] = v
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
	for w := v; w >= 0; w = this.pathTo[w] {
		path = append(path, w)
	}

	res := make([]int, len(path))
	for i := range path {
		res[len(res)-i-1] = path[i]
	}
	return res
}
