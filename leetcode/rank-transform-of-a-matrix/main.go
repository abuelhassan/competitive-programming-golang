// https://leetcode.com/problems/rank-transform-of-a-matrix/
package main

import (
	"fmt"
	"sort"
)

type disjointSets struct {
	parents []int
	ranks   []int
	mxVals  []int
}

func (dsu *disjointSets) find(i int) int {
	if dsu.parents[i] == i {
		return i
	}
	dsu.parents[i] = dsu.find(dsu.parents[i])
	return dsu.parents[i]
}
func (dsu *disjointSets) join(a, b, v int) {
	a, b = dsu.find(a), dsu.find(b)
	if a == b {
		return // already connected
	}
	if dsu.ranks[a] < dsu.ranks[b] {
		a, b = b, a
	}
	if dsu.ranks[a] == dsu.ranks[b] {
		dsu.ranks[a]++
	}
	dsu.parents[b] = a
	dsu.mxVals[a] = max(dsu.mxVals[a], dsu.mxVals[b], v)
}

func newDSU(n int) disjointSets {
	parents, ranks, mxVals := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parents[i], ranks[i], mxVals[i] = i, 0, 0
	}
	return disjointSets{
		parents: parents,
		ranks:   ranks,
		mxVals:  mxVals,
	}
}

func matrixRankTransform(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	vals, mp := make([]int, 0), make(map[int][][2]int)
	for i := 0; i < n; i++ {
		for j, v := range matrix[i] {
			if _, ok := mp[v]; !ok {
				vals = append(vals, v)
			}
			mp[v] = append(mp[v], [2]int{i, j})
		}
	}
	sort.Ints(vals)

	mxRow, mxCol := make([]int, n), make([]int, m)
	for _, v := range vals {
		dsu := newDSU(n + m)
		for _, cell := range mp[v] {
			dsu.join(cell[0], n+cell[1], max(mxRow[cell[0]], mxCol[cell[1]]))
		}
		for _, cell := range mp[v] {
			val, i, j := dsu.mxVals[dsu.find(cell[0])]+1, cell[0], cell[1]
			matrix[i][j], mxRow[i], mxCol[j] = val, val, val
		}
	}
	return matrix
}

func main() {
	tt := []struct {
		in  [][]int
		out [][]int
	}{
		{
			in: [][]int{
				{-5, -4, -2},
				{1, -1, -2},
				{-2, -1, -3},
			},
			out: [][]int{
				{1, 2, 3},
				{5, 4, 3},
				{2, 4, 1},
			},
		},
	}

	equal := func(a, b [][]int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if len(a[i]) != len(b[i]) {
				return false
			}
			for j := 0; j < len(a[i]); j++ {
				if a[i][j] != b[i][j] {
					return false
				}
			}
		}
		return true
	}

	for _, t := range tt {
		if out := matrixRankTransform(t.in); !equal(out, t.out) {
			panic(fmt.Sprintf("\noutput: %v\nexpected: %v\n", out, t.out))
		}
	}
	fmt.Println("Todo bien!")
}
