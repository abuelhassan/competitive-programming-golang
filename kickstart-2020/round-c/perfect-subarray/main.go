package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	io := newFastIO()
	defer io.Flush()

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var n, a, ans, mxSm, sm int
		sumCnt := make(map[int]int)
		io.Read(&n)
		for i := 0; i < n; i++ {
			io.Read(&a)
			sumCnt[sm] += 1
			sm += a
			mxSm += abs(a)
			for j := 0; j*j <= mxSm; j++ {
				ans += sumCnt[sm-(j*j)]
			}
		}
		io.Write("Case #%d: %d\n", caseID, ans)
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
