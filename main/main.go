package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("./main/test")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	delimiter := regexp.MustCompile(` `)

	ar_size, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	var ar []int32
	ar_items := delimiter.Split(strings.TrimSpace(readLine(reader)), -1)

	for i := 0; i < int(ar_size); i++ {
		ar_Item_Temp, err := strconv.ParseInt(ar_items[i], 10, 64)
		checkError(err)
		ar_Item := int32(ar_Item_Temp)
		ar = append(ar, ar_Item)
	}

	output := solve(ar)
	fmt.Fprint(writer, output)
	writer.Flush()
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
