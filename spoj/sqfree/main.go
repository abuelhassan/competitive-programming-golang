package main

import (
	"bufio"
	"fmt"
	"os"
)

func mobius(sz int) []int {
	notPrime := make([]bool, sz)
	primes := make([]int, 0)
	for i := 2; i < sz; i++ {
		if !notPrime[i] {
			primes = append(primes, i)
			for j := i * i; j < sz; j += i {
				notPrime[j] = true
			}
		}
	}
	mob := make([]int, sz)
	for i := 0; i < sz; i++ {
		mob[i] = 1
	}
	for _, p := range primes {
		for j := p * p; j < sz; j += p * p {
			mob[j] = 0
		}
	}
	for _, p := range primes {
		for j := p; j < sz; j += p {
			mob[j] = mob[j] * -1
		}
	}
	return mob
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	const mxSZ = 10000001
	mob := mobius(mxSZ)

	var t int
	_, _ = fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		var n int64
		_, _ = fmt.Fscan(r, &n)
		ans := n
		// The Inclusion-Exclusion Principle
		for i := int64(2); i*i <= n; i++ {
			ans += int64(mob[i]) * (n / (i * i))
		}
		_, _ = fmt.Fprintln(w, ans)
	}
}
