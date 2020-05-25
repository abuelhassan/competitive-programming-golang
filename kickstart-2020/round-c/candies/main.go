package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var n, m int
		io.Read(&n, &m)
		pfx, mpfx := newBIT(n), newBIT(n)
		for i, x := 1, 0; i <= n; i++ {
			io.Read(&x)
			if i%2 == 0 {
				x = -x
			}
			pfx.update(i, x)
			mpfx.update(i, x*i)
		}
		ans := 0
		var qType string
		for i, a, b := 0, 0, 0; i < m; i++ {
			io.Read(&qType, &a, &b)
			switch qType {
			case "U":
				if a%2 == 0 {
					b = -b
				}
				pfx.update(a, b-pfx.getRange(a, a))
				mpfx.update(a, (b*a)-mpfx.getRange(a, a))
			case "Q":
				v := mpfx.getRange(a, b) - (a-1)*(pfx.getRange(a, b))
				if a%2 == 0 {
					v = -v
				}
				ans += v
			}
		}
		io.Write("Case #%d: %d\n", caseID, ans)
	}
}

func newBIT(n int) bit {
	return bit{make([]int, n)}
}

type bit struct {
	bit []int
}

// 1-based index
func (b *bit) update(i, v int) {
	for i <= len(b.bit) {
		b.bit[i-1] += v
		i += i & -i
	}
}

// 1-based index
func (b bit) get(i int) int {
	pfx := 0
	for i > 0 {
		pfx += b.bit[i-1]
		i -= i & -i
	}
	return pfx
}

// 1-based indices
func (b bit) getRange(i, j int) int {
	return b.get(j) - b.get(i-1)
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
