// https://leetcode.com/problems/maximum-genetic-difference-query/
//
// sub-problem:
// Given a list of integers (values) and another list of integers (queries).
// For each query X, get the number from the array that maximizes the XOR with X.
//
// In order to maximize the XOR with X, get the number that if XORed with X maximizes the most significant bit.
// In case of ties, try maximizing the next bit, and so on...
// Saving all the numbers in a trie, in their binary representation
// from the most significant bit to the least significant bit is the way to go!

package main

import (
	"fmt"
)

const bitsLen = 18 // max value is 2e5

type (
	query struct {
		idx int
		val int
	}
	trieNode struct {
		cnt      int
		children [2]*trieNode
	}
)

func update(root *trieNode, val, inc int) {
	for i := bitsLen - 1; i >= 0; i-- {
		bit := (val >> i) & 1
		if root.children[bit] == nil {
			root.children[bit] = &trieNode{}
		}
		root.children[bit].cnt += inc
		root = root.children[bit]
	}
}

func maxXOR(root *trieNode, val int) int {
	ans := 0
	for i := bitsLen - 1; i >= 0; i-- {
		bit := (val >> i) & 1
		// invert bit if possible.
		if root.children[1-bit] != nil && root.children[1-bit].cnt > 0 {
			bit = 1 - bit
			ans = ans | (1 << i)
		}
		root = root.children[bit]
	}
	return ans
}

func dfs(curNode int, adj [][]int, queries [][]query, tr *trieNode, ans []int) {
	update(tr, curNode, 1)
	for _, q := range queries[curNode] {
		ans[q.idx] = maxXOR(tr, q.val)
	}
	for _, v := range adj[curNode] {
		dfs(v, adj, queries, tr, ans)
	}
	update(tr, curNode, -1)
}

func maxGeneticDifference(parents []int, queries [][]int) []int {
	adj, root := make([][]int, len(parents)), -1
	for u, p := range parents {
		if p == -1 {
			root = u
		} else {
			adj[p] = append(adj[p], u)
		}
	}

	offline := make([][]query, len(parents))
	for i, q := range queries {
		offline[q[0]] = append(offline[q[0]], query{idx: i, val: q[1]})
	}

	ans := make([]int, len(queries))
	dfs(root, adj, offline, &trieNode{}, ans)
	return ans
}

func main() {
	tt := []struct {
		parents []int
		queries [][]int
		out     []int
	}{
		{
			parents: []int{-1, 0, 1, 1},
			queries: [][]int{{0, 2}, {3, 2}, {2, 5}},
			out:     []int{2, 3, 7},
		},
		{
			parents: []int{3, 7, -1, 2, 0, 7, 0, 2},
			queries: [][]int{{4, 6}, {1, 15}, {0, 5}},
			out:     []int{6, 14, 7},
		},
	}

	equal := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	for _, t := range tt {
		if out := maxGeneticDifference(t.parents, t.queries); !equal(out, t.out) {
			panic(fmt.Sprintf("\noutput: %v\nexpected: %v\n", out, t.out))
		}
	}
	fmt.Println("Todo bien!")
}
