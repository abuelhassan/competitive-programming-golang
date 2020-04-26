package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(pats []string) string {
	pfx, in, sfx := "", "", ""
	for _, pat := range pats {
		fIdx := strings.IndexByte(pat, '*')
		curPfx := pat[:fIdx]
		if len(curPfx) > len(pfx) {
			curPfx, pfx = pfx, curPfx
		}
		if !strings.HasPrefix(pfx, curPfx) {
			return "*"
		}

		lIdx := strings.LastIndexByte(pat, '*')
		curSfx := pat[lIdx+1:]
		if len(curSfx) > len(sfx) {
			curSfx, sfx = sfx, curSfx
		}
		if !strings.HasSuffix(sfx, curSfx) {
			return "*"
		}

		for j := fIdx + 1; j < lIdx; j++ {
			if pat[j] != '*' {
				in = in + string(pat[j])
			}
		}
	}
	return pfx + in + sfx
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var n int
		io.Read(&n)
		pats := make([]string, n)
		for i := 0; i < n; i++ {
			io.Read(&pats[i])
		}
		io.Write("Case #%d: %s\n", caseID, solve(pats))
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
