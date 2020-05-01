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
		var n int
		io.Read(&n)
		arr := make([]int, n)
		peaks := 0
		for i := 0; i < n; i++ {
			io.Read(&arr[i])
			if i > 1 && arr[i-1] > arr[i-2] && arr[i-1] > arr[i] {
				peaks++
			}
		}
		io.Write("Case #%d: %d\n", caseID, peaks)
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
