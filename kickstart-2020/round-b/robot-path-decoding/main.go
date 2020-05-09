package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(str string) (int, int) {
	const mod = 1000000000
	mp := map[rune][]int{
		'S': {0, 1},
		'N': {0, mod - 1},
		'E': {1, 0},
		'W': {mod - 1, 0},
	}

	mult := func(a, b int) int {
		return int((int64(a) * int64(b)) % mod)
	}
	add := func(a, b int) int {
		return (a + b) % mod
	}

	st := newStack(0, 0, 1)
	for _, c := range str {
		dir, isDir := mp[c]
		switch {
		case isDir:
			ox, oy, m := st.pop()
			st.push(add(ox, dir[0]), add(oy, dir[1]), m)
		case c == ')':
			ox, oy, om := st.pop()
			tx, ty, tm := st.pop()
			st.push(add(tx, mult(ox, om)), add(ty, mult(oy, om)), tm)
		case c >= '1' && c <= '9':
			st.push(0, 0, int(c-'0'))
		}
	}
	return st.st[0][0], st.st[0][1]
}

func newStack(x, y, m int) stack {
	return stack{[][]int{{x, y, m}}}
}

type stack struct {
	st [][]int
}

func (s *stack) push(x, y, m int) {
	s.st = append(s.st, []int{x, y, m})
}

func (s *stack) pop() (int, int, int) {
	i := len(s.st) - 1
	x, y, m := s.st[i][0], s.st[i][1], s.st[i][2]
	s.st = s.st[:i] // O(1)
	return x, y, m
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var str string
		io.Read(&str)
		x, y := solve(str)
		io.Write("Case #%d: %d %d\n", caseID, x+1, y+1)
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
