// https://leetcode.com/problems/critical-connections-in-a-network
package main

import "fmt"

type stack []int

// Using a stack with ascending values. O(N) solution.
func largestRectangleArea(heights []int) int {
	n, ans, st := len(heights), 0, stack{-1}
	for i := 0; i < n; i++ {
		for st.top() != -1 && heights[st.top()] >= heights[i] {
			cur := st.pop()
			area := (i - st.top() - 1) * heights[cur] // width * height
			if area > ans {
				ans = area
			}
		}
		st.push(i)
	}
	for st.top() != -1 {
		cur := st.pop()
		area := (n - st.top() - 1) * heights[cur] // width * height
		if area > ans {
			ans = area
		}
	}
	return ans
}

func (st *stack) push(a int) {
	*st = append(*st, a)
}
func (st *stack) pop() int {
	elem := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return elem
}
func (st *stack) top() int {
	return (*st)[len(*st)-1]
}

func main() {
	tt := []struct {
		in  []int
		out int
	}{
		{
			in:  []int{2, 1, 5, 6, 2, 3},
			out: 10,
		},
		{
			in:  []int{2, 4, 5, 5, 4, 4, 3},
			out: 20,
		},
		{
			in:  []int{2, 4, 5, 5, 3, 3, 3, 1},
			out: 18,
		},
	}
	for _, t := range tt {
		if out := largestRectangleArea(t.in); out != t.out {
			panic(fmt.Sprintf("\ninput: %v\noutput: %v\nexpected: %v\n", t.in, out, t.out))
		}
	}
	fmt.Println("Todo bien!")
}
