package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	_, _ = fmt.Fscan(r, &n)
	_, _ = fmt.Fprintf(w, "input: %d", n)
}
