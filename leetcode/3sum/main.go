// https://leetcode.com/problems/3sum/
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sm := nums[i] + nums[j] + nums[k]
			if sm < 0 {
				j++
				continue
			}
			if sm > 0 {
				k--
				continue
			}
			triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
			for j < k && nums[j] == nums[j+1] {
				j++
			}
			for j < k && nums[k] == nums[k-1] {
				k--
			}
			j, k = j+1, k-1
		}
	}
	return triplets
}

func main() {
	tt := []struct {
		inp []int
		out [][]int
	}{
		{
			inp: []int{-1, 0, 1, 2, -1, -4},
			out: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			inp: []int{0, 0, 0, 0},
			out: [][]int{{0, 0, 0}},
		},
	}

	eq := func(a [][]int, b [][]int) bool {
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
		if out := threeSum(t.inp); !eq(out, t.out) {
			panic(fmt.Sprintf("\ninput: %v\noutput: %v\nexpected: %v\n", t.inp, out, t.out))
		}
	}
	fmt.Println("Todo bien!")
}
