package ch4

func IsAcyclic(adj [][]int) bool {
	marked := make([]bool, len(adj))
	acyclic := true

	var dfs func(int, int)
	dfs = func (v int, u int) {
		marked[v] = true
		for _, w := range adj[v] {
			if !marked[w] {
				dfs(w, v)
			} else if w != u {
				acyclic = false
			}
		}
	}

	for i := range adj {
		if marked[i] {
			continue
		}
		dfs(i, i)
	}

	return acyclic
}