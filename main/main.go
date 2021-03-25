package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLine() string {
	r, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return r
}

func ReadInt() (res [][]int) {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		curNums := []int{}
		str := reader.Text()
		trimmed := strings.TrimSpace(str)
		strArr := strings.Split(trimmed, " ")
		for _, value := range strArr {
			curInt, _ := strconv.Atoi(value)
			curNums = append(curNums, curInt)
		}
		res = append(res, curNums)
		fmt.Println(res)
	}
	return
}

func ReadInt2() (res []int) {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		str := reader.Text()
		trimmed := strings.TrimSpace(str)
		strArr := strings.Split(trimmed, " ")
		for _, value := range strArr {
			curInt, _ := strconv.Atoi(value)
			res = append(res, curInt)
		}
	}
	return
}

func Start() {
	//fmt.Println(ReadLine())
	ReadInt2()
}

func main() {
	Start()
}
