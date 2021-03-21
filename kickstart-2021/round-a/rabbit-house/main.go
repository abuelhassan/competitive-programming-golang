package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func solve(mat [][]int) int {
	var (
		di = []int{0, 0, 1, -1}
		dj = []int{1, -1, 0, 0}
	)

	r, c := len(mat), len(mat[0])
	h := make(maxHeap, 0, r*c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			h.Push(node{
				p: mat[i][j],
				i: i,
				j: j,
			})
		}
	}
	heap.Init(&h)

	ans := 0
	for h.Len() > 0 {
		n := heap.Pop(&h).(node)
		if n.p != mat[n.i][n.j] { // visited
			continue
		}
		for d := 0; d < 4; d++ {
			ni := n.i + di[d]
			nj := n.j + dj[d]
			if ni < 0 || nj < 0 || ni >= r || nj >= c {
				continue
			}
			if mat[ni][nj] < n.p-1 {
				ans += n.p - 1 - mat[ni][nj]
				mat[ni][nj] = n.p - 1
				heap.Push(&h, node{
					p: n.p - 1,
					i: ni,
					j: nj,
				})
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

type (
	maxHeap []node
	node    struct {
		p int
		i int
		j int
	}
)

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].p > h[j].p }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(node))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
