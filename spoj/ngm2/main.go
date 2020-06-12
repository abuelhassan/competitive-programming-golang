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

func main() {
	io := newFastIO()
	defer io.Flush()

	var n int64
	var k int
	io.Read(&n, &k)
	factors := make([]int, k)
	for i := 0; i < k; i++ {
		io.Read(&factors[i])
	}

	// The Inclusion-Exclusion Principle
	ans := n
	for i := 1; i < (1 << uint(k)); i++ {
		popcount := 0
		l := int64(1)
		for j := 0; j < k; j++ {
			if (i>>uint(j))&1 == 1 {
				popcount++
				l = lcm(l, int64(factors[j]))
			}
		}
		if popcount&1 == 1 {
			l = -l
		}
		ans += n / l
	}
	io.Write("%d\n", ans)
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
