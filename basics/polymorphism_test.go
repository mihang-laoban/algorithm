package basics

import (
	"fmt"
	"testing"
)

type Obj interface {
	Name()
	Age()
}

type one struct {
	name string
	age  int
}

func (this *one) Name() {
	fmt.Println(this.name)
}

func (this *one) Age() {
	fmt.Println(this.age)
}

type two struct {
	name string
	age  int
}

func (this *two) Name() {
	fmt.Println(this.name)
}

func (this *two) Age() {
	fmt.Println(this.age)
}

type three struct {
	two
}

func printObj(obj Obj) {
	obj.Age()
	obj.Name()
}

func TestPhmorph(t *testing.T) {
	th := &three{}
	th.name = "three"
	th.age = 3
	test := []interface{}{
		&one{name: "one", age: 1},
		&two{name: "two", age: 2},
		th,
	}
	//printObj(&one{name: "one", age: 1})
	//printObj(&two{name: "two", age: 2})
	//printObj(test[0].(*one))
	for i := 0; i < len(test); i++ {
		switch test[i].(type) {
		case *one:
			printObj(test[i].(*one))
		case *two:
			printObj(test[i].(*two))
		case *three:
			printObj(test[i].(*three))
		}
	}
}
