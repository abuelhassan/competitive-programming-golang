package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s string, k int) int {
	score := 0
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			score++
		}
		i++
		j--
	}
	if score > k {
		return score - k
	}
	return k - score
}

func main() {
	io := newBuffIO()
	defer io.Flush()

	var (
		cases int
		k     int
		str   string
	)

	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		io.Read(&k, &k, &str)
		io.Write("Case #%d: %d\n", caseID, solve(str, k))
	}
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
