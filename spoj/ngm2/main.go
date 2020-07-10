package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int64) int64 {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int64) int64 {
	return a * (b / gcd(a, b))
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int64
	var k int
	_, _ = fmt.Fscan(r, &n, &k)
	factors := make([]int, k)
	for i := 0; i < k; i++ {
		_, _ = fmt.Fscan(r, &factors[i])
	}

	// The Inclusion-Exclusion Principle
	ans := n
	for i := 1; i < (1 << uint(k)); i++ {
		popcount := 0
		l := int64(1)
		for j := 0; j < k; j++ {
			if (i>>uint(j))&1 == 1 {
				popcount++
				l = lcm(l, int64(factors[j]))
			}
		}
		if popcount&1 == 1 {
			l = -l
		}
		ans += n / l
	}
	_, _ = fmt.Fprintln(w, ans)
}
