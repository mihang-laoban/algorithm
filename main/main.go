package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
* Implement method/function with name 'solve' below.
* The function accepts following as parameters.
*  1. ar is of type int32.
* return int32.
 */
func solve(ar []int32) int32 {
	memo := map[int32]int{}

	for _, v := range ar {
		memo[v]++
	}
	if memo[0] < memo[1] {
		return int32(memo[0]) * 2
	}
	return int32(memo[1]) * 2
}

func main() {
	fmt.Println(123)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
