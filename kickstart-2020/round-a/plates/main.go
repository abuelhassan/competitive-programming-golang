package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(p int, stacks [][]int) interface{} {
	n, m := len(stacks), len(stacks[0])
	dp := newDP(2, p+1, 0)
	for i := 0; i < n; i++ {
		for j := 0; j <= p; j++ {
			dp.set(i, j, 0)
		}
		for j, sm := 0, 0; j <= m; j++ {
			for k := 0; j+k <= p; k++ {
				best := sm
				if i > 0 {
					best += dp.get(i-1, k).(int)
				}
				x := dp.get(i, j+k).(int)
				if best < x {
					best = x
				}
				dp.set(i, j+k, best)
			}
			if j < m {
				sm += stacks[i][j]
			}
		}
	}
	return dp.get(n-1, p)
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
	get(i, j int) interface{}
	set(i, j int, val interface{})
}

func newDP(h, w int, val interface{}) dpIface {
	m := make([][]interface{}, h)
	for i := 0; i < h; i++ {
		m[i] = make([]interface{}, w)
		for j := 0; j < w; j++ {
			m[i][j] = val
		}
	}
	return &memo{m: m}
}

type memo struct {
	m [][]interface{}
}

func (m memo) get(i, j int) interface{} {
	return m.m[i%len(m.m)][j]
}

func (m *memo) set(i, j int, val interface{}) {
	m.m[i%len(m.m)][j] = val
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
