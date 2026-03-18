// https://leetcode.com/problems/maximum-points-activated-with-one-addition/
package main

type disjointsets struct {
	parent []int
	size   []int
}

func (dsu *disjointsets) find(i int) int {
	if i == dsu.parent[i] {
		return i
	}
	dsu.parent[i] = dsu.find(dsu.parent[i])
	return dsu.parent[i]
}

func (dsu *disjointsets) join(a, b int) {
	a, b = dsu.find(a), dsu.find(b)
	if a == b {
		return
	}
	if dsu.size[a] < dsu.size[b] {
		a, b = b, a
	}
	dsu.parent[b] = a
	dsu.size[a] = dsu.size[a] + dsu.size[b]
}

func newDSU(n int) disjointsets {
	parent, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i], size[i] = i, 1
	}
	return disjointsets{
		parent: parent,
		size:   size,
	}
}

func maxActivated(points [][]int) int {
	dsu, xs, ys := newDSU(len(points)), make(map[int]int), make(map[int]int)

	// connect all points with equal x values to the same component.
	// same goes for y.
	for i, p := range points {
		xv, ok := xs[p[0]]
		if ok {
			dsu.join(xv, i)
		} else {
			xs[p[0]] = i // first occurrence of x
		}

		yv, ok := ys[p[1]]
		if ok {
			dsu.join(yv, i)
		} else {
			ys[p[1]] = i // first occurrence of y
		}
	}

	mx1, mx2 := 0, 0
	for i, p := range dsu.parent {
		if i != p {
			continue
		}

		// component root (i == p)
		sz := dsu.size[i]
		if sz > mx1 {
			mx1, mx2 = sz, mx1
		} else if sz > mx2 {
			mx2 = sz
		}
	}
	return mx1 + mx2 + 1
}
