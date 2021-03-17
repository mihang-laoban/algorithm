package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Param map[string]interface{}

type Show struct {
	Param
}

func TestInterview(t *testing.T) {
	s := new(Show)
	s.Param = map[string]interface{}{}
	s.Param["RMB"] = 10000
	Jay(student{Name: "tar"})
	Jay(&student{Name: "target"})
}

type student struct {
	Name string
}

func Jay(v interface{}) {
	switch msg := v.(type) {
	case *student:
		fmt.Println(msg.Name)
	case student:
		fmt.Println(msg.Name)
	}
}

func Test(t *testing.T) {
	js := `{"name":"asd"}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("people:", p)

	o := &People{"Asdf"}
	fmt.Println(o.String())
}

type People struct {
	Name string `json:"name"`
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p.Name)
}

func TestFind(t *testing.T) {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				close(ch)
				return
			}
			fmt.Println("a:", a)
		}
	}()
	fmt.Println("ok")
	time.Sleep(5 * time.Second)
}

func TestFib(t *testing.T) {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(6))
}

func TestF(t *testing.T) {
	n := 6

	a, b := 1, 1
	for i := 3; i < n+1; i++ {
		c := a + b
		a = b
		b = c
	}
	fmt.Println(b)
}
