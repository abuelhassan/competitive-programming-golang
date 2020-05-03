package main

import (
	"bufio"
	"fmt"
	"os"
)

type dpt float32 // dp type

func solve(w, h, l, u, r, d int) dpt {
	if (l == 1 && u == 1) || (r == w && d == h) {
		return 0
	}
	hole := rect{l - 1, u - 1, r - 1, d - 1}
	dp := newDP(w, 2)
	dp.set(h-1, w-1, 1)
	for i, j := h-1, w-2; i >= 0; i, j = i-1, w-1 {
		for ; j >= 0; j-- {
			if hole.contains(i, j) {
				dp.set(i, j, 0)
				continue
			}
			if i == h-1 {
				dp.set(i, j, dp.get(i, j+1))
				continue
			}
			if j == w-1 {
				dp.set(i, j, dp.get(i+1, j))
				continue
			}
			dp.set(i, j, (dp.get(i+1, j)+dp.get(i, j+1))/2)
		}
	}
	return dp.get(0, 0)
}

type rect struct {
	l, u, r, d int
}

func (r rect) contains(i, j int) bool {
	return i >= r.u && i <= r.d && j >= r.l && j <= r.r
}

type dpIface interface {
	get(i, j int) dpt
	set(i, j int, val dpt)
}

func newDP(w, h int) dpIface {
	m := make([][]dpt, h)
	for i := 0; i < h; i++ {
		m[i] = make([]dpt, w)
	}
	return &memo{m: m}
}

type memo struct {
	m [][]dpt
}

func (m memo) get(i, j int) dpt {
	return m.m[i%len(m.m)][j]
}

func (m *memo) set(i, j int, val dpt) {
	if i < 0 || j < 0 {
		return
	}
	m.m[i%len(m.m)][j] = val
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var w, h, l, u, r, d int
		io.Read(&w, &h, &l, &u, &r, &d)
		io.Write("Case #%d: %f\n", caseID, solve(w, h, l, u, r, d))
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
