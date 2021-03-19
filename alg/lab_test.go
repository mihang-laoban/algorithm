package alg

import (
	"bytes"
	"context"
	linkedList2 "dp/alg/linkedList"
	. "dp/ds/linkedList"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"
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

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)

	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}

func TestConcurrency(t *testing.T) {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "test")
	_ = tr.Event(context.Background(), "test")
	_ = tr.Event(context.Background(), "test")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}

var errMy = errors.New("my error is here")

func TestWrapErr(t *testing.T) {
	err := service()
	fmt.Printf("main: %+v\n", err)
}

func service() error {
	return biz()
}

func biz() error {
	return dao()
}

func dao() error {
	return errors.Wrap(errMy, "dao failed")
}

func TestSize(t *testing.T) {
	fmt.Println(unsafe.Sizeof(struct {
		i8  int8
		i16 int16
		i32 int32
	}{}))
	fmt.Println(unsafe.Sizeof(struct {
		i8  int64
		i32 int32
		i16 int16
	}{}))
}

// fmt.Sprintf
func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < 10; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}

// add
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < 10; j++ {
			str = str + string(j)
		}
	}
	b.StopTimer()
}

// bytes.Buffer
func BenchmarkStringBuffer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		for j := 0; j < 10; j++ {
			buffer.WriteString(strconv.Itoa(j))
		}
		_ = buffer.String()
	}
	b.StopTimer()
}

// strings.Builder
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 10; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}

func TestRemoveDuplicate(t *testing.T) {
	str := "asbb j   pq[we"
	fmt.Println(str)
	fmt.Println(io.EOF)
}
