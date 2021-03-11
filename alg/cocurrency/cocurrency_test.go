package cocurrency

import (
	"fmt"
	"testing"
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
