package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(str string) string {
	res := ""
	bal := 0
	for _, c := range str {
		dig := int(c - '0')
		diff, par := 1, "("
		if dig < bal {
			diff, par = -1, ")"
		}
		for bal != dig {
			res += par
			bal += diff
		}
		res = res + string(c)
	}
	for bal > 0 {
		res = res + ")"
		bal--
	}
	return res
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var str string
		io.Read(&str)
		io.Write("Case #%d: %s\n", caseID, solve(str))
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
