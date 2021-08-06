package main

import (
	"fmt"
	"sort"
)

// p is the parent node.
// That works fine because there are no repeated edges.
// In case of repeated edges, edge id should be used instead.
func dfs(u, p int, timer *int, inTime, low []int, adj [][]int, bridges *[][]int) {
	inTime[u] = *timer
	low[u] = *timer
	*timer++
	for _, v := range adj[u] {
		if v == p {
			continue
		}
		if inTime[v] != -1 { // visited
			low[u] = min(low[u], inTime[v])
			continue
		}
		dfs(v, u, timer, inTime, low, adj, bridges)
		low[u] = min(low[u], low[v])
		if low[v] > inTime[u] { // there is no back edges from v to u or any of u's ancestors.
			*bridges = append(*bridges, []int{u, v})
		}
	}
}

// The graph is guaranteed to be connected.
// There are no self-loops. And there are no repeated connections.
func criticalConnections(n int, connections [][]int) [][]int {
	adj := make([][]int, n)
	for _, edge := range connections {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	inTime, low := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		inTime[i], low[i] = -1, -1
	}

	bridges, timer := make([][]int, 0), 0
	dfs(0, -1, &timer, inTime, low, adj, &bridges)

	return bridges
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	tt := []struct {
		n     int
		edges [][]int
		out   [][]int
	}{
		{
			n:     4,
			edges: [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}},
			out:   [][]int{{1, 3}},
		},
	}

	equal := func(a, b [][]int) bool {
		// size matters!
		if len(a) != len(b) {
			return false
		}

		for i := 0; i < len(a); i++ {
			// pair (x, y) is the same as pair (y, x)
			sort.Ints(a[i])
			sort.Ints(b[i])
		}

		// order doesn't matter
		var ptr *[][]int
		less := func(i, j int) bool {
			for k := 0; k < len((*ptr)[i]) && k < len((*ptr)[j]); k++ {
				if (*ptr)[i][k] != (*ptr)[j][k] {
					return (*ptr)[i][k] < (*ptr)[j][k]
				}
			}
			return len((*ptr)[i]) < len((*ptr)[j])
		}
		ptr = &a
		sort.Slice(a, less)
		ptr = &b
		sort.Slice(b, less)

		// check equality
		for i := 0; i < len(a); i++ {
			if len(a[i]) != len(b[i]) {
				return false
			}
			for j := 0; j < len(a[i]); j++ {
				if a[i][j] != b[i][j] {
					return false
				}
			}
		}
		return true
	}

	for _, t := range tt {
		if out := criticalConnections(t.n, t.edges); !equal(out, t.out) {
			panic(fmt.Sprintf("\noutput: %v\nexpected: %v\n", out, t.out))
		}
	}
	fmt.Println("Todo bien!")
}
