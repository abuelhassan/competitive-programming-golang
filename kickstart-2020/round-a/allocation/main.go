package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solve(arr []int, m int) int {
	sort.Ints(arr)
	ans := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > m {
			break
		}
		m -= arr[i]
		ans++
	}
	return ans
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var n, m int
		io.Read(&n, &m)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			io.Read(&arr[i])
		}
		io.Write("Case #%d: %d\n", caseID, solve(arr, m))
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
