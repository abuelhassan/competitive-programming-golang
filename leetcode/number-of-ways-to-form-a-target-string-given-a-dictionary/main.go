// https://leetcode.com/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/
package main

func dp(i int, j int, target string, freq []map[byte]int, memo [][]int) int {
	if i == len(target) {
		return 1
	}
	if j == len(freq) {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	pass := dp(i, j+1, target, freq, memo)
	pick := int64(dp(i+1, j+1, target, freq, memo)) * int64(freq[j][target[i]])
	memo[i][j] = (pass + int(pick%1000000007)) % 1000000007
	return memo[i][j]
}

func numWays(words []string, target string) int {
	freq := make([]map[byte]int, len(words[0]))
	for i := 0; i < len(words[0]); i++ {
		freq[i] = make(map[byte]int)
		for j := 0; j < len(words); j++ {
			freq[i][words[j][i]]++
		}
	}
	memo := make([][]int, len(target))
	for i := 0; i < len(target); i++ {
		memo[i] = make([]int, len(freq))
		for j := 0; j < len(freq); j++ {
			memo[i][j] = -1
		}
	}
	return dp(0, 0, target, freq, memo)
}
