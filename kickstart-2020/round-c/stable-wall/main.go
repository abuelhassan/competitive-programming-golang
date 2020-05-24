package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(shape [][]rune) string {
	n, m, ans := len(shape), len(shape[0]), ""

	// 0 -> not found, 1 -> found but not safe, 2 -> safe
	getState := func(c rune) int {
		state := 0
		indices := make([][]int, 0)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if shape[i][j] == c {
					state = 2
					indices = append(indices, []int{i, j})
					if i != n-1 && shape[i+1][j] != c && shape[i+1][j] != '-' {
						return 1
					}
				}
			}
		}
		for _, idx := range indices {
			shape[idx[0]][idx[1]] = '-'
		}
		return state
	}

	for {
		state := 0
		for c := 'A'; c <= 'Z'; c++ {
			st := getState(c)
			switch st {
			case 1:
				if state == 0 {
					state = 1
				}
			case 2:
				ans = ans + string(c)
				state = 2
			}
		}
		if state == 0 {
			break
		}
		if state == 1 {
			return "-1"
		}
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
		shape := make([][]rune, n)
		for i := 0; i < n; i++ {
			var s string
			io.Read(&s)
			shape[i] = []rune(s)
		}
		io.Write("Case #%d: %s\n", caseID, solve(shape))
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
