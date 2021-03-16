package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLine() string {
	r, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return r
}

func Start() {
	fmt.Println(ReadLine())
}

func main() {
	Start()
}
