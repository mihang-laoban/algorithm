package cocurrency

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Foo struct {
	startTwo   chan bool
	startThree chan bool
	over       chan bool
}

func (f *Foo) one() {
	fmt.Println("one")
	f.startTwo <- true
}

func (f *Foo) two() {
	<-f.startTwo
	fmt.Println("two")
	f.startThree <- true
}

func (f *Foo) three() {
	<-f.startThree
	fmt.Println("three")
	f.over <- true
}

func Test(t *testing.T) {
	f := &Foo{}
	f.startTwo = make(chan bool)
	f.startThree = make(chan bool)
	f.over = make(chan bool)
	go f.one()
	go f.two()
	go f.three()
	<-f.over
}

func TestCounterThreadSafe(t *testing.T) {
	//var mut sync.Mutex
	//count := 0
	var count rune
	for i := 0; i < 5000; i++ {
		go func() {
			atomic.AddInt32(&count, 2)
			//mut.Lock()
			//count++
			//mut.Unlock()
		}()
	}
	time.Sleep(1000 * time.Millisecond)
	t.Log(count)
}

func TestCounterThreadSaf(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.RWMutex
	var count rune
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() { mut.Unlock() }()
			mut.Lock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(count)
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 50)
	fmt.Println("task is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func TestAsynService(t *testing.T) {
	retCH := AsynService()
	otherTask()
	fmt.Println(<-retCH)
	time.Sleep(100 * time.Millisecond)
}

func AsynService() chan string {
	retCH := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCH <- ret
		fmt.Println("service exited")
	}()
	return retCH
}

func TestContextTimeOut(t *testing.T) {
	// 超时控制
	fmt.Println(jobTimeOut())
}

func jobTimeOut() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Millisecond * 1500)
		done <- struct{}{}
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("timeout")
	}
}

func job() int {
	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(5)
	fmt.Println(res)
	time.Sleep(time.Second * time.Duration(res))
	return res
}

func TestGetFaster(t *testing.T) {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go func() {
			c <- job()
		}()
	}
	fmt.Println("Fastest: ", <-c, "second")
}
