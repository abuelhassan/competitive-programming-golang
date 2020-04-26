package main

import (
	"bufio"
	"fmt"
	"os"
)

func solveSmall(n int) (moves [][]int) {
	for i := 0; i < n; i++ {
		moves = append(moves, []int{i + 1, 1})
	}
	return moves
}

func solve(n int) (moves [][]int) {
	if n <= 500 {
		return solveSmall(n)
	}
	n = n - 30
	ones := 30
	for i, j, dj := 1, 1, 1; ones > 0; i++ {
		func() {
			defer func() {
				if j != 1 {
					j++
				}
				n >>= 1
			}()
			if n&1 == 0 {
				ones--
				moves = append(moves, []int{i, j})
				return
			}
			for j >= 1 && j <= i {
				moves = append(moves, []int{i, j})
				j += dj
			}
			j -= dj
			if i != 1 {
				dj *= -1
			}
		}()
	}
	return moves
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var n int
		io.Read(&n)
		moves := solve(n)
		io.Write("Case #%d:\n", caseID)
		for _, m := range moves {
			io.Write("%d %d\n", m[0], m[1])
		}
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
