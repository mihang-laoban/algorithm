package alg

import (
	linkedList2 "dp/alg/linkedList"
	. "dp/ds/linkedList"
	"fmt"
	"github.com/pkg/errors"
	"testing"
	"time"
)

func TestAlgorithmComponentsBreakLab(t *testing.T) {
}

func linkedList() {
	head := ArrayToLinkedList([]int{1, 2, 3, 4})
	cur := linkedList2.ReverseListR(head)

	fmt.Println(LinkedListToArray(cur))
}

//
func isPositive(a int) (bool, error) {
	if a == 0 {
		return false, errors.New("undefined")
	}
	return a > 0, nil
}

type MyError struct {
	code int
	msg  string
	data interface{}
}

func (err *MyError) Error() string {
	return fmt.Sprintf("%d - %s - %s", err.code, err.msg, err.data)
}

func throwTest() error {
	return &MyError{404, "page cannot be found", ""}
}

type temp interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := err.(temp)
	return ok && te.Temporary()
}

func TestErrSample(t *testing.T) {
	if v, err := isPositive(0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
	err := throwTest()
	if err, ok := err.(*MyError); ok {
		errors.Wrap(err, "open fail")
	}

}

func TestPanicHandling(t *testing.T) {
	Go(func() {
		panic("一路向西")
	})
	time.Sleep(5 * time.Second)
}

// 全局sync包，处理挂掉的野生goroutine
func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		x()
	}()
}
