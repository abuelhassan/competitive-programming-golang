package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(p int, stacks [][]int) int {
	n, m := len(stacks), len(stacks[0])
	dp := newDP(2, p+1)
	for j, sm := 0, 0; j <= p; j++ {
		dp.set(0, j, sm)
		if j < m {
			sm += stacks[0][j]
		}
	}
	for i := 1; i < n; i++ {
		dp.clear(i)
		for j, sm := 0, 0; j <= m; j++ {
			for k := 0; j+k <= p; k++ {
				dp.set(i, j+k, max(dp.get(i, j+k), sm+dp.get(i-1, k)))
			}
			if j < m {
				sm += stacks[i][j]
			}
		}
	}
	return dp.get(n-1, p)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var n, k, p int
		io.Read(&n, &k, &p)
		stacks := make([][]int, n)
		for i := 0; i < n; i++ {
			stacks[i] = make([]int, k)
			for j := 0; j < k; j++ {
				io.Read(&stacks[i][j])
			}
		}
		io.Write("Case #%d: %d\n", caseID, solve(p, stacks))
	}
}

type dpIface interface {
	get(i, j int) int
	set(i, j int, val int)
	clear(i int)
}

func newDP(h, w int) dpIface {
	m := make([][]int, h)
	for i := 0; i < h; i++ {
		m[i] = make([]int, w)
	}
	return &memo{m: m}
}

type memo struct {
	m [][]int
}

func (m memo) get(i, j int) int {
	return m.m[i%len(m.m)][j]
}

func (m *memo) set(i, j int, val int) {
	m.m[i%len(m.m)][j] = val
}

func (m *memo) clear(i int) {
	m.m[i%len(m.m)] = make([]int, len(m.m[0]))
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
