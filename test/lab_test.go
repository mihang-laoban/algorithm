package test

import (
	"dp/tools"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
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

type Animal struct {
}

func (m *Animal) Eat() {
	fmt.Println("eat")
}

func TestCallByReflect(t *testing.T) {
	animal := Animal{}
	value := reflect.TypeOf(&animal)
	f, _ := value.MethodByName("Eat")
	fmt.Println(f)
}

var count int64 = 2

func fish(wg *sync.WaitGroup, counter int64, fishCh, catCh chan struct{}) {
	for {
		if counter > count {
			wg.Done()
			return
		}
		<-fishCh
		fmt.Println("fish")
		atomic.AddInt64(&counter, 1)
		catCh <- struct{}{}
	}
}

func cat(wg *sync.WaitGroup, counter int64, catCh, dogCh chan struct{}) {
	for {
		if counter > count {
			wg.Done()
			return
		}
		<-catCh
		fmt.Println("cat")
		atomic.AddInt64(&counter, 1)
		dogCh <- struct{}{}
	}
}

func dog(wg *sync.WaitGroup, counter int64, dogCh, catCh chan struct{}) {
	for {
		if counter > count {
			wg.Done()
			return
		}
		<-dogCh
		fmt.Println("dog")
		atomic.AddInt64(&counter, 1)
		catCh <- struct{}{}
	}
}

func TestGo(t *testing.T) {
	var (
		wg          sync.WaitGroup
		dogCounter  int64
		fishCounter int64
		catCounter  int64
	)
	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)
	catCh := make(chan struct{}, 1)
	wg.Add(3)
	go dog(&wg, dogCounter, dogCh, fishCh)
	go fish(&wg, fishCounter, fishCh, catCh)
	go cat(&wg, catCounter, catCh, dogCh)
	dogCh <- struct{}{}
	wg.Wait()
}

func TestSize(t *testing.T) {
	//fmt.Println(unsafe.Sizeof())
	var n1 int8 = 127
	var n2 int16 = 10
	var n3 int32 = 10
	var n4 int64 = 10
	var n5 int = 10
	var n6 float32 = 10
	var n7 float64 = 10
	var n8 bool = true
	var n9 byte = 10
	var n10 rune = 10
	fmt.Printf("%T: %d\n", n1, unsafe.Sizeof(n1))
	fmt.Printf("%T: %d\n", n2, unsafe.Sizeof(n2))
	fmt.Printf("%T: %d\n", n3, unsafe.Sizeof(n3))
	fmt.Printf("%T: %d\n", n4, unsafe.Sizeof(n4))
	fmt.Printf("%T: %d\n", n5, unsafe.Sizeof(n5))
	fmt.Printf("%T: %d\n", n6, unsafe.Sizeof(n6))
	fmt.Printf("%T: %d\n", n7, unsafe.Sizeof(n7))
	fmt.Printf("%T: %d\n", n8, unsafe.Sizeof(n8))
	fmt.Printf("%T: %d\n", n9, unsafe.Sizeof(n9))
	fmt.Printf("%T: %d\n", n10, unsafe.Sizeof(n10))

	fmt.Println(tools.Bit2Int(tools.Int2bit(3)))
	fmt.Println(tools.Int2bit(256))
	fmt.Println(256 * 256)
}
