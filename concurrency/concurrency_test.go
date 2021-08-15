package concurrency

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	//testSimple()
	testClose()
	//testMergeInput()¬
	//testQuit()
	//testPCB()
	//testTimeout()
	//testInAndOutChan()
	//testMaxNumControl()
	//testSignal()
	//testSynchronize()
}

func testSimple() {
	intChan := make(chan int)

	go func() {
		intChan <- 1
	}()

	value := <-intChan
	fmt.Println("value : ", value)
}

/*
测试关闭channel
*/
func testClose() {
	ch := make(chan int, 5)
	sign := make(chan int, 2)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		fmt.Println("the channel is closed")
		sign <- 0

	}()

	go func() {
		for {
			i, ok := <-ch
			fmt.Printf("%d, %v \n", i, ok)

			if !ok {
				break
			}

			time.Sleep(time.Second * 2)
		}
		sign <- 1
	}()

	<-sign
	<-sign
}

/**
将多个输入的channel进行合并成一个channel
*/
func testMergeInput() {
	input1 := make(chan int)
	input2 := make(chan int)
	output := make(chan int)

	go func(in1, in2 <-chan int, out chan<- int) {
		for {
			select {
			case v := <-in1:
				out <- v
			case v := <-in2:
				out <- v
			}
		}
	}(input1, input2, output)

	go func() {
		for i := 0; i < 10; i++ {
			input1 <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for i := 20; i < 30; i++ {
			input2 <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			select {
			case value := <-output:
				fmt.Println("输出：", value)
			}
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("主线程退出")
}

/*
测试channel用于通知中断退出的问题
*/
func testQuit() {
	g := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println(v)
			case <-quit:
				fmt.Println("B退出")
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		g <- i
	}
	quit <- true
	fmt.Println("testAB退出")
}

/**
生产者消费者问题
*/
func testPCB() {
	fmt.Println("test PCB")

	intchan := make(chan int)
	quitChan := make(chan bool)
	quitChan2 := make(chan bool)

	value := 0

	go func() {
		for i := 0; i < 3; i++ {

			value = value + 1
			intchan <- value

			fmt.Println("write finish, value ", value)

			time.Sleep(time.Second)
		}
		quitChan <- true
	}()
	go func() {
		for {
			select {
			case v := <-intchan:
				fmt.Println("read finish, value ", v)
			case <-quitChan:
				quitChan2 <- true
				return
			}
		}

	}()

	<-quitChan2
	fmt.Println("task is done ")
}

/*
检查channel读写超时，并做超时的处理
*/
func testTimeout() {
	g := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println(v)
			case <-time.After(time.Second * time.Duration(3)):
				quit <- true
				fmt.Println("超时，通知主线程退出")
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		g <- i
	}

	<-quit
	fmt.Println("收到退出通知，主线程退出")
}

/*
指定channel是输入还是输出型的，防止编写时写错误输入输出，指定了的话，可以在编译时期作错误的检查
*/
func testInAndOutChan() {
	ch := make(chan int)
	quit := make(chan bool)

	//输入型的chan是这种格式的：inChan chan<- int，如果换成输出型的，则编译时会报错
	go func(inChan chan<- int) {
		for i := 0; i < 10; i++ {
			inChan <- i
			time.Sleep(time.Millisecond * 500)
		}
		quit <- true
		quit <- true
	}(ch)

	go func(outChan <-chan int) {
		for {
			select {
			case v := <-outChan:
				fmt.Println("print out value : ", v)
			case <-quit:
				fmt.Println("收到退出通知，退出")
				return
			}
		}
	}(ch)

	<-quit
	fmt.Println("收到退出通知，主线程退出")
}

/*
测试通过channel来控制最大并发数，来处理事件
*/
func testMaxNumControl() {
	maxNum := 3
	limit := make(chan bool, maxNum)
	quit := make(chan bool)

	for i := 0; i < 100; i++ {
		fmt.Println("start worker : ", i)

		limit <- true

		go func(i int) {
			fmt.Println("do worker start: ", i)
			time.Sleep(time.Millisecond * 20)
			fmt.Println("do worker finish: ", i)

			<-limit

			if i == 99 {
				fmt.Println("完成任务")
				quit <- true
			}

		}(i)
	}

	<-quit
	fmt.Println("收到退出通知，主程序退出")
}

func testSignal() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		time.Sleep(time.Second * 2)

		number := 0
		for {
			number++
			println("number : ", number)
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("按Ctrl+C可退出程序")
	<-quit
	fmt.Println("主程序退出")
}

//同步控制模型，生产者模型
var lockChan = make(chan int, 1)
var remainMoney = 1000

func testSynchronize() {
	quit := make(chan bool, 2)

	go func() {
		for i := 0; i < 10; i++ {
			money := (rand.Intn(12) + 1) * 100
			go testSynchronize_expense(money)

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}

		quit <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			money := (rand.Intn(12) + 1) * 100
			go testSynchronize_gain(money)

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}

		quit <- true
	}()

	<-quit
	<-quit

	fmt.Println("主程序退出")
}

func testSynchronize_expense(money int) {
	lockChan <- 0

	if remainMoney >= money {
		srcRemainMoney := remainMoney
		remainMoney -= money
		fmt.Printf("原来有%d, 花了%d，剩余%d\n", srcRemainMoney, money, remainMoney)
	} else {
		fmt.Printf("想消费%d钱不够了, 只剩%d\n", money, remainMoney)
	}

	<-lockChan
}

func testSynchronize_gain(money int) {
	lockChan <- 0

	srcRemainMoney := remainMoney
	remainMoney += money
	fmt.Printf("原来有%d, 赚了%d，剩余%d\n", srcRemainMoney, money, remainMoney)

	<-lockChan
}
