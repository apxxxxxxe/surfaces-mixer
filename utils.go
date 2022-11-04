package main

import (
	"fmt"
	"os"
	"strings"
)

func isFileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func countDigit(n int) int {
	res := 0
	for n > 0 {
		n /= 10
		res++
	}
	return res
}

func printPartNum(n, i int) string {
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", countDigit(n)), i)
}

func addIndents(s string, n int) string {
	res := ""
	for _, l := range strings.Split(strings.Trim(s, "\n"), "\n") {
		res += strings.Repeat(" ", n) + l + "\n"
	}
	return res
}
