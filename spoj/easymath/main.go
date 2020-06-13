package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int64) int64 {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int64) int64 {
	return a * (b / gcd(a, b))
}

func solve(n int64, factors []int64) int64 {
	// The Inclusion-Exclusion Principle
	ans := n
	for i := 1; i < 1<<uint(len(factors)); i++ {
		popcount := 0
		l := int64(1)
		for j, f := range factors {
			if i>>uint(j)&1 == 1 {
				popcount++
				l = lcm(l, f)
			}
		}
		if popcount&1 == 1 {
			l = -l
		}
		ans += n / l
	}
	return ans
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var t int
	io.Read(&t)
	for c := 0; c < t; c++ {
		var n, m, a, d int64
		io.Read(&n, &m, &a, &d)
		factors := []int64{a, a + d, a + (d * 2), a + (d * 3), a + (d * 4)}
		io.Write("%d\n", solve(m, factors)-solve(n-1, factors))
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
