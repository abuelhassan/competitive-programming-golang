package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(mat [][]int) (square, rows, cols int) {
	n := len(mat)
	hayDuplicates := func(i, j, di, dj int) bool {
		vis := make(map[int]bool)
		for i < n && j < n {
			if vis[mat[i][j]] {
				return true
			}
			vis[mat[i][j]] = true
			i += di
			j += dj
		}
		return false
	}
	for i := 0; i < len(mat); i++ {
		square += mat[i][i]
		if hayDuplicates(i, 0, 0, 1) {
			rows++
		}
		if hayDuplicates(0, i, 1, 0) {
			cols++
		}
	}
	return
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var n int
		io.Read(&n)
		mat := make([][]int, n)
		for i := 0; i < n; i++ {
			mat[i] = make([]int, n)
			for j := 0; j < n; j++ {
				io.Read(&mat[i][j])
			}
		}
		square, rows, cols := solve(mat)
		io.Write("Case #%d: %d %d %d\n", caseID, square, rows, cols)
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
