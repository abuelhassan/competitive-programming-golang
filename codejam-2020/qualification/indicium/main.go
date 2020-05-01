package main

import (
	"bufio"
	"fmt"
	"os"
)

// Backtracking generating arrays of length n that sums up to k,
// where every element is an integer from 0 to n-1.
func generateDiagonals(n, k int, diag []int, in <-chan bool, out chan<- []int) bool {
	if n == len(diag) && k == 0 {
		out <- diag
		select {
		case getNext := <-in:
			return !getNext // returning false, means more combinations are needed.
		}
	}
	if n == len(diag) || k < 0 || k > (n-len(diag))*(n-1) { // dead end
		return false
	}
	for i := 0; i < n; i++ {
		diag = append(diag, i)
		if generateDiagonals(n, k-i, diag, in, out) {
			return true // no more combinations are needed.
		}
		diag = diag[:len(diag)-1]
	}
	if len(diag) == 0 { // no more available combinations.
		out <- nil
	}
	return false
}

func solve(n, k int) (string, [][]int) {
	getNext := make(chan bool)
	diagonals := make(chan []int)
	defer func() {
		close(getNext)
		close(diagonals)
	}()

	go generateDiagonals(n, k-n, []int{}, getNext, diagonals) // diagonals generator.

	for {
		select {
		case diag := <-diagonals:
			if diag == nil { // no more combinations available.
				return "IMPOSSIBLE", nil
			}
			ok, mat := getMat(diag)
			if ok {
				getNext <- false // stop the backtracking go-routine.
				return "POSSIBLE", mat
			}
			getNext <- true
		}
	}
}

func getMat(diag []int) (bool, [][]int) {
	n := len(diag)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mat[i][j] = -1
		}
	}
	for r := 0; r < n; r++ {
		adj := getAdjList(r, diag, mat)
		ok, bipB := maxBipartiteMatching(adj)
		if !ok {
			return false, nil
		}
		for i := 0; i < n; i++ {
			mat[r][bipB[i]] = i
		}
	}
	return true, mat
}

func getAdjList(r int, diag []int, mat [][]int) [][]int {
	n := len(mat)
	adj := make([][]int, n)
	for c := 0; c < n; c++ {
		if c == r {
			// For cells on the diagonal, give them only one possible edge.
			// So that, they are forced to use that edge.
			adj[c] = append(adj[c], diag[r])
			continue
		}
		vis := map[int]bool{
			diag[r]: true,
			diag[c]: true,
		}
		for i := 0; i < r; i++ {
			vis[mat[i][c]] = true
		}
		for i := 0; i < n; i++ {
			if !vis[i] {
				adj[c] = append(adj[c], i)
			}
		}
	}
	return adj
}

func maxBipartiteMatching(adj [][]int) (bool, []int) {
	n := len(adj)
	bipB := make([]int, n)
	for i := 0; i < n; i++ {
		bipB[i] = -1
	}
	for i := 0; i < n; i++ {
		if !bm(i, adj, make(map[int]bool), bipB) {
			return false, nil
		}
	}
	return true, bipB
}

func bm(u int, adj [][]int, vis map[int]bool, bipB []int) bool {
	for _, v := range adj[u] {
		if vis[v] {
			continue
		}
		vis[v] = true
		if bipB[v] == -1 || bm(bipB[v], adj, vis, bipB) {
			bipB[v] = u
			return true
		}
	}
	return false
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var n, k int
		io.Read(&n, &k)
		possible, mat := solve(n, k)
		io.Write("Case #%d: %s\n", caseID, possible)
		for _, r := range mat {
			for _, c := range r {
				io.Write("%d ", c+1)
			}
			io.Write("\n")
		}
	}
}

type fastIO struct {
	r *bufio.Reader
	w *bufio.Writer
}

func newFastIO() fastIO {
	return fastIO{
		r: bufio.NewReader(os.Stdin),
		w: bufio.NewWriter(os.Stdout),
	}
}

func (io *fastIO) Read(args ...interface{}) {
	_, _ = fmt.Fscan(io.r, args...)
}

func (io *fastIO) Write(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(io.w, format, args...)
}

func (io *fastIO) Flush() {
	_ = io.w.Flush()
}
