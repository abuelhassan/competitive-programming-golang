// https://leetcode.com/problems/delete-duplicate-folders-in-system/

package main

import (
	"fmt"
	"sort"
)

const (
	pathSeparator int64 = '/'
	pathEnd       int64 = '#'
	hashMod       int64 = 2000000011
	hashBase      int64 = 128
)

type (
	node struct {
		val      string
		hash     hash
		children children
	}

	// ordered children. because maps have non-deterministic order.
	children struct {
		indices map[string]int
		nodes   []*node
	}

	// hash for concatenated paths for all children
	hash struct {
		len int
		val int64
	}
)

func newNode(val string) *node {
	return &node{
		val: val,
		children: children{
			indices: make(map[string]int),
			nodes:   make([]*node, 0),
		},
	}
}

// returns (b^p) % mod
func fastPower(b int64, p int, mod int64) int64 {
	if p == 0 {
		return 1
	}
	if p == 1 {
		return b
	}
	res := fastPower(b, p/2, mod) * 2
	if p%2 == 1 {
		res += b
	}
	return res % mod
}

func concatHash(h1 hash, h2 hash) hash {
	return hash{
		val: (h1.val*fastPower(hashBase, h2.len, hashMod) + h2.val) % hashMod,
		len: h1.len + h2.len,
	}
}

// paths should be sorted in descending order.
func makeTree(paths [][]string) *node {
	root := newNode("")
	for i, path := range paths {
		if i > 0 && len(paths[i]) < len(paths[i-1]) {
			continue // this path is already added, by one of its children.
		}
		it := root
		for _, s := range path {
			if idx, ok := it.children.indices[s]; ok {
				it = it.children.nodes[idx]
				continue
			}
			idx := len(it.children.nodes)
			it.children.nodes = append(it.children.nodes, newNode(s))
			it.children.indices[s] = idx
			it = it.children.nodes[idx]
		}
	}
	return root
}

// generates the hashes and construct a count for all hashes
func generateHashes(tree *node, hashCnt map[int64]int) hash {
	if len(tree.children.nodes) == 0 {
		return hash{val: pathEnd, len: 1}
	}
	tree.hash = hash{val: 0, len: 0}
	for _, ch := range tree.children.nodes {
		tree.hash = concatHash(tree.hash, hash{val: pathSeparator, len: 1})
		for _, c := range ch.val {
			tree.hash = concatHash(tree.hash, hash{val: int64(c), len: 1})
		}
		tree.hash = concatHash(tree.hash, generateHashes(ch, hashCnt))
	}
	hashCnt[tree.hash.val]++
	return tree.hash
}

func getUniqueFiles(tree *node, hashCnt map[int64]int, cur []string, ans *[][]string) {
	cur = append(cur, tree.val)
	*ans = append(*ans, append([]string{}, cur...)) // append a deep copy of cur
	for _, ch := range tree.children.nodes {
		if hashCnt[ch.hash.val] == 1 || len(ch.children.nodes) == 0 {
			getUniqueFiles(ch, hashCnt, cur, ans)
		}
	}
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	// sort in descending order
	sort.Slice(paths, func(i, j int) bool {
		i, j = j, i // descending order
		for k := 0; k < len(paths[i]) && k < len(paths[j]); k++ {
			if paths[i][k] != paths[j][k] {
				return paths[i][k] < paths[j][k]
			}
		}
		return len(paths[i]) < len(paths[j])
	})

	hashCnt := make(map[int64]int)
	tree := makeTree(paths)
	generateHashes(tree, hashCnt)

	ans := make([][]string, 0)
	for _, ch := range tree.children.nodes {
		if hashCnt[ch.hash.val] == 1 || len(ch.children.nodes) == 0 {
			getUniqueFiles(ch, hashCnt, []string{}, &ans)
		}
	}
	return ans
}

func main() {
	tt := []struct {
		in  [][]string
		out [][]string
	}{
		{
			in: [][]string{
				{"a"},
				{"b"},
				{"a", "x"},
				{"b", "x"},
				{"a", "x", "y"},
				{"b", "x", "y"},
				{"a", "x", "z"},
				{"b", "x", "z"},
			},
			out: [][]string{},
		},
		{
			in: [][]string{
				{"a"},
				{"a", "x"},
				{"a", "x", "y"},
				{"a", "x", "z"},
				{"b"},
				{"b", "x"},
				{"b", "x", "y"},
				{"b", "x", "z"},
				{"c"},
				{"c", "w"},
				{"c", "w", "y"},
				{"c", "w", "z"},
				{"d"},
				{"d", "e"},
				{"d", "e", "f"},
				{"d", "e", "g"},
				{"d", "e", "g", "y"},
				{"d", "e", "g", "z"},
			},
			out: [][]string{
				{"c"},
				{"d"},
				{"d", "e"},
				{"d", "e", "f"},
			},
		},
	}

	equal := func(a, b [][]string) bool {
		// size matters!
		if len(a) != len(b) {
			return false
		}

		// order doesn't matter
		var ptr *[][]string
		less := func(i, j int) bool {
			for k := 0; k < len((*ptr)[i]) && k < len((*ptr)[j]); k++ {
				if (*ptr)[i][k] != (*ptr)[j][k] {
					return (*ptr)[i][k] < (*ptr)[j][k]
				}
			}
			return len((*ptr)[i]) < len((*ptr)[j])
		}
		ptr = &a
		sort.Slice(a, less)
		ptr = &b
		sort.Slice(b, less)

		// check equality
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
		if out := deleteDuplicateFolder(t.in); !equal(out, t.out) {
			panic(fmt.Sprintf("\noutput: %v\nexpected: %v\n", out, t.out))
		}
	}
	fmt.Println("Todo bien!")
}
