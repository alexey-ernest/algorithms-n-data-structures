package ch4

func IsTwoColorable(adj [][]int) bool {
	marked := make([]bool, len(adj))
	color := make([]bool, len(adj))
	istwocolorable := true

	var dfs func(int)
	dfs = func (v int) {
		marked[v] = true
		for _, w := range adj[v] {
			if !marked[w] {
				color[w] = !color[v]
				dfs(w)
			} else if color[w] == color[v] {
				istwocolorable = false
			}
		}
	}

	for i := range adj {
		if marked[i] {
			continue
		}
		dfs(i)
	}

	return istwocolorable
}