package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solve(w, h, l, u, r, d int, pre *preCalcs) float64 {
	pre.preCalc(w + h)
	getProb := func(n, k, lm int) float64 {
		if k == lm {
			return 0
		}
		var p float64 = 0
		for ; k <= n; k++ {
			p += math.Pow(2, pre.pre[n]-pre.pre[k]-pre.pre[n-k]-float64(n))
		}
		return p
	}
	return getProb(l+d-2, d, h) + getProb(r+u-2, r, w)
}

type preCalcs struct {
	pre []float64
}

func newPreCalcs() preCalcs {
	pre := make([]float64, 1, 200000)
	pre[0] = 0
	return preCalcs{pre}
}

func (l *preCalcs) preCalc(n int) {
	for i := len(l.pre); i < n; i++ {
		l.pre = append(l.pre, l.pre[i-1]+math.Log2(float64(i)))
	}
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var w, h, l, u, r, d int
		io.Read(&w, &h, &l, &u, &r, &d)
		pre := newPreCalcs()
		io.Write("Case #%d: %f\n", caseID, solve(w, h, l, u, r, d, &pre))
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
