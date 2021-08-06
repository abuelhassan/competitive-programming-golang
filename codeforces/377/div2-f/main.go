package main

import (
	"bufio"
	"fmt"
	"os"
)

// visiting by edge is required, because of the orientation step.
func getBridges(u int, timer *int, tin, low []int, adj [][][2]int, edges [][2]int, vis []bool, bridges map[int]struct{}) {
	tin[u], low[u] = *timer, *timer
	*timer++
	for _, e := range adj[u] {
		v, edgeID := e[0], e[1]
		if vis[edgeID] {
			continue
		}
		vis[edgeID] = true
		if u != edges[edgeID][0] { // orientation step (bridges should be re-oriented later)
			edges[edgeID][0], edges[edgeID][1] = u, v
		}
		if tin[v] != 0 { // visited node
			low[u] = min(low[u], tin[v])
			continue
		}
		getBridges(v, timer, tin, low, adj, edges, vis, bridges)
		low[u] = min(low[u], low[v])
		if low[v] > tin[u] { // there is no back edges from v to u or any of u's ancestors.
			bridges[edgeID] = struct{}{}
		}
	}
}

// gets component size without using any of the bridges.
func componentSizes(u int, adj [][][2]int, vis []bool, bridges map[int]struct{}) int {
	vis[u] = true
	q, sz := []int{u}, 0
	for len(q) > 0 {
		u, q, sz = q[0], q[1:], sz+1
		for _, e := range adj[u] {
			v, edgeID := e[0], e[1]
			if _, isBridge := bridges[edgeID]; isBridge || vis[v] {
				continue
			}
			vis[v] = true
			q = append(q, v)
		}
	}
	return sz
}

func reorientBridges(u int, adj [][][2]int, vis []bool, edges [][2]int, bridges map[int]struct{}) {
	vis[u] = true
	for _, e := range adj[u] {
		v, edgeID := e[0], e[1]
		if vis[v] {
			continue
		}
		if _, isBridge := bridges[edgeID]; isBridge && edges[edgeID][0] != v {
			edges[edgeID][0], edges[edgeID][1] = v, u
		}
		reorientBridges(v, adj, vis, edges, bridges)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	_, _ = fmt.Fscan(r, &n, &m)

	edges, adj := make([][2]int, m), make([][][2]int, n)
	for i := 0; i < m; i++ {
		_, _ = fmt.Fscan(r, &edges[i][0], &edges[i][1])
		edges[i][0], edges[i][1] = edges[i][0]-1, edges[i][1]-1
		adj[edges[i][0]] = append(adj[edges[i][0]], [2]int{edges[i][1], i})
		adj[edges[i][1]] = append(adj[edges[i][1]], [2]int{edges[i][0], i})
	}

	// Get bridges and orient SCCs
	var (
		timer    = 1
		tin, low = make([]int, n), make([]int, n)
		vis      = make([]bool, m+1)
		bridges  = make(map[int]struct{})
	)
	getBridges(0, &timer, tin, low, adj, edges, vis, bridges)

	// Get a node to the maximum component
	mxNode, mxCompSize := -1, 0
	for i := 0; i < n; i++ {
		vis[i] = false
	}
	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		sz := componentSizes(i, adj, vis, bridges)
		if sz > mxCompSize {
			mxCompSize, mxNode = sz, i
		}
	}

	// Orient bridges toward the largest component.
	for i := 0; i < n; i++ {
		vis[i] = false
	}
	reorientBridges(mxNode, adj, vis, edges, bridges)

	_, _ = fmt.Fprintln(w, mxCompSize)
	for _, edge := range edges {
		_, _ = fmt.Fprintf(w, "%d %d\n", edge[0]+1, edge[1]+1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
