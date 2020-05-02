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
		var n, d int
		io.Read(&n, &d)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			io.Read(&arr[i])
		}
		for i := len(arr) - 1; i >= 0; i-- {
			d = (d / arr[i]) * arr[i]
		}
		io.Write("Case #%d: %d\n", caseID, d)
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
