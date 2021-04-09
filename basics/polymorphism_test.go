package basics

import (
	"fmt"
	"reflect"
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

func TestPolymorphism(t *testing.T) {
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

//=============================================================================================

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("write a book")
}

func TestBook(t *testing.T) {
	b := &Book{}
	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	// r: pair <type: Book, value:book{}>
	w = r.(Writer) // 此处断言成功是因为w,r的concrete type一致
	w.WriteBook()
}

//=============================================================================================

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", this)
}

func TestReflect(t *testing.T) {
	user := User{1, "AceId", 18}
	DoFieldAndMethod(user)
}

func DoFieldAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType.Name())

	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		fmt.Println()
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

}

//=============================================================================================

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagInfo, " doc:", tagDoc)
	}
}

func TestTag(t *testing.T) {
	var re resume
	findTag(&re)
}

//=============================================================================================
