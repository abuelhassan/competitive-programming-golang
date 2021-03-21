package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(mat [][]int) int {
	ans := countL(mat)
	for i := 0; i < 3; i++ {
		mat = rotate(mat)
		ans += countL(mat)
	}
	return ans
}

func rotate(mat [][]int) [][]int {
	r, c := len(mat), len(mat[0])
	rot := make([][]int, c)
	for i := 0; i < c; i++ {
		rot[i] = make([]int, r)
		for j := 0; j < r; j++ {
			rot[i][j] = mat[j][c-i-1]
		}
	}
	return rot
}

// countL counts L shapes and mirrored (on the y axis) L shapes.
func countL(mat [][]int) int {
	r, c := len(mat), len(mat[0])

	cntRight := make([][]int, r)
	for i := 0; i < r; i++ {
		cntRight[i] = make([]int, c)
		for j, cnt := c-1, 0; j >= 0; j-- {
			if mat[i][j] == 0 {
				cnt = 0
				continue
			}
			cnt++
			cntRight[i][j] = cnt
		}
	}

	cntLeft := make([][]int, r)
	for i := 0; i < r; i++ {
		cntLeft[i] = make([]int, c)
		for j, cnt := 0, 0; j < c; j++ {
			if mat[i][j] == 0 {
				cnt = 0
				continue
			}
			cnt++
			cntLeft[i][j] = cnt
		}
	}

	ans := 0
	for j := 0; j < c; j++ {
		for i, cnt := 0, 0; i < r; i++ {
			if mat[i][j] == 0 {
				cnt = 0
				continue
			}
			cnt++
			if cnt < 4 {
				continue
			}
			vert, horR, horL := cnt/2-1, cntRight[i][j]-1, cntLeft[i][j]-1
			if horR > 0 {
				ans += min(vert, horR)
			}
			if horL > 0 {
				ans += min(vert, horL)
			}
		}
	}
	return ans
}

func main() {
	io := newBuffIO()
	defer io.Flush()

	var (
		cases int
		r     int
		c     int
		mat   [][]int
	)

	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		io.Read(&r, &c)
		mat = make([][]int, r)
		for i := 0; i < r; i++ {
			mat[i] = make([]int, c)
			for j := 0; j < c; j++ {
				io.Read(&mat[i][j])
			}
		}
		io.Write("Case #%d: %d\n", caseID, solve(mat))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type buffIO struct {
	r *bufio.Reader
	w *bufio.Writer
}

func newBuffIO() buffIO {
	return buffIO{
		r: bufio.NewReader(os.Stdin),
		w: bufio.NewWriter(os.Stdout),
	}
}

func (io *buffIO) Read(args ...interface{}) {
	_, _ = fmt.Fscan(io.r, args...)
}

func (io *buffIO) Write(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(io.w, format, args...)
}

func (io *buffIO) Flush() {
	_ = io.w.Flush()
}
