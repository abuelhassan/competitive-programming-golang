package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type task struct {
	st  int
	ed  int
	idx int
}

func solve(tasks []task) string {
	byStart := func(t1, t2 *task) bool {
		return t1.st < t2.st
	}
	sortType(byStart).sort(tasks)
	ans := make([]rune, len(tasks))
	cEnd, jEnd := 0, 0
	for _, t := range tasks {
		switch {
		case cEnd <= t.st:
			ans[t.idx], cEnd = 'C', t.ed
		case jEnd <= t.st:
			ans[t.idx], jEnd = 'J', t.ed
		default:
			return "IMPOSSIBLE"
		}
	}
	return string(ans)
}

func main() {
	io := newFastIO()
	defer io.Flush()

	var cases int
	io.Read(&cases)
	for caseID := 1; caseID <= cases; caseID++ {
		var n int
		io.Read(&n)
		tasks := make([]task, n)
		for i := 0; i < n; i++ {
			io.Read(&tasks[i].st)
			io.Read(&tasks[i].ed)
			tasks[i].idx = i
		}
		io.Write("Case #%d: %s\n", caseID, solve(tasks))
	}
}

type taskSorter struct {
	tasks []task
	by    func(t1, t2 *task) bool
}

func (s *taskSorter) Len() int {
	return len(s.tasks)
}

func (s *taskSorter) Swap(i, j int) {
	s.tasks[i], s.tasks[j] = s.tasks[j], s.tasks[i]
}

func (s *taskSorter) Less(i, j int) bool {
	return s.by(&s.tasks[i], &s.tasks[j])
}

type sortType func(t1, t2 *task) bool

func (b sortType) sort(tasks []task) {
	ts := &taskSorter{
		tasks: tasks,
		by:    b,
	}
	sort.Sort(ts)
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
