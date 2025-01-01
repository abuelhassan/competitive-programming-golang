package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := readInt(r)
	_, _ = fmt.Fprintf(w, "input: %d", n)
}

func readInt(in *bufio.Reader) int {
	str, _ := in.ReadString('\n')
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	num, _ := strconv.Atoi(str)
	return num
}

func readLine(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	return strings.Split(line, " ")
}

func readIntLine(in *bufio.Reader) []int {
	tokens := readLine(in)
	nums := make([]int, len(tokens))
	for i, token := range tokens {
		nums[i], _ = strconv.Atoi(token)
	}
	return nums
}
