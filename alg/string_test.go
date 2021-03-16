package alg

import (
	"bytes"
	"dp/tools"
	"fmt"
	"strings"
	"testing"
	"time"
)

//allocs/op 表示每个操作（单次迭代）发生了多少个不同的内存分配。
//B/op每个操作分配了多少字节。

// 所以他们的区别就在于 bytes.Buffer 是重新申请了一块空间，存放生成的string变量， 而strings.Builder直接将底层的[]byte转换成了string类型返回了回来，去掉了申请空间的操作。

func TestBytesBuffer(t *testing.T) {
	loop := 100000
	var s string
	s1 := "hello"
	s2 := "world"
	var start time.Time
	//加号+连接
	start = time.Now()
	for i := 0; i < loop; i++ {
		s1 += s2
	}
	fmt.Println("+连接方法:", time.Since(start))
	//append连接
	s1 = "hello"
	s2 = "world"
	start = time.Now()
	for i := 0; i < loop; i++ {
		s = string(append([]byte(s1), s2...))
	}
	fmt.Println("append方法:", time.Since(start))
	//Join方法连接
	v := []string{"hello", "world"}
	start = time.Now()
	for i := 0; i < loop; i++ {
		s = strings.Join(v, "")
	}
	fmt.Println("strings.Join方法:", time.Since(start))
	//bytes.writestring方法
	start = time.Now()
	for i := 0; i < loop; i++ {
		var buf bytes.Buffer
		buf.WriteString("hello")
		buf.WriteString("world")
		buf.String()
	}
	fmt.Println("bytes.writestring方法:", time.Since(start))
	//fmt方法连接
	start = time.Now()
	for i := 0; i < loop; i++ {
		s = fmt.Sprintf("%s%s", "hello", "world")
	}
	fmt.Println("fmt方法:", time.Since(start))
	fmt.Println(s)
}

func BenchmarkMyStringTest(b *testing.B) {
	a := "I am a"
	c := "good student"
	var buffer bytes.Buffer
	buffer.WriteString(a)
	buffer.WriteString(c)
	resBuffer := buffer.String()
	fmt.Println(resBuffer)

	var builder strings.Builder
	builder.WriteString(a)
	builder.WriteString(c)
	resBuilder := buffer.String()
	fmt.Println(resBuilder)
}

func TestString(t *testing.T) {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.ContainsAny("seafood", "z"))
	fmt.Println(strings.ContainsAny("seafood", "zd"))
	fmt.Println(strings.ContainsAny("seafood", ""))
	fmt.Println(strings.ContainsRune("我是中国", '我'))
	fmt.Println(strings.Count("aabc", "a"))
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println(strings.Fields(" foo bar  baz   "))

	avail := map[int32]bool{
		'*':  true,
		'|':  true,
		'\t': true,
	}
	var ans strings.Builder
	for _, record := range []string{" aaa*1892*122", "aaa\taa\t", "124|939|22"} {
		tmp := strings.FieldsFunc(record, func(ch rune) bool {
			switch {
			case avail[ch]:
				return true
			}
			return false
		})
		ans.WriteString(strings.Join(tmp, " "))
	}
	res := ans.String()
	fmt.Println(res)
	fmt.Println(tools.ReverseStr(res))

	//======================================================================================================

	fmt.Println("")
	fmt.Println("HasPrefix 函数的用法")
	fmt.Println(strings.HasPrefix("NLT_abc", "NLT")) //前缀是以NLT开头的

	fmt.Println("")
	fmt.Println("HasSuffix 函数的用法")
	fmt.Println(strings.HasSuffix("NLT_abc", "abc")) //后缀是以NLT开头的

	fmt.Println("")
	fmt.Println("Index 函数的用法")
	fmt.Println(strings.Index("NLT_abc", "abc")) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.Index("NLT_abc", "aaa")) // 在存在返回 -1
	fmt.Println(strings.Index("我是中国人", "中"))     // 在存在返回 6

	fmt.Println("")
	fmt.Println("IndexAny 函数的用法")
	fmt.Println(strings.IndexAny("我是中国人", "中")) // 在存在返回 6
	fmt.Println(strings.IndexAny("我是中国人", "和")) // 在存在返回 -1

	fmt.Println("")
	fmt.Println("Index 函数的用法")
	fmt.Println(strings.IndexRune("NLT_abc", 'b')) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.IndexRune("NLT_abc", 's')) // 在存在返回 -1
	fmt.Println(strings.IndexRune("我是中国人", '中'))   // 在存在返回 6

	fmt.Println("")
	fmt.Println("Join 函数的用法")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // 返回字符串：foo, bar, baz

	fmt.Println("")
	fmt.Println("LastIndex 函数的用法")
	fmt.Println(strings.LastIndex("go gopher", "go")) // 3

	fmt.Println("")
	fmt.Println("LastIndexAny 函数的用法")
	fmt.Println(strings.LastIndexAny("go gopher", "go")) // 4
	fmt.Println(strings.LastIndexAny("我是中国人", "中"))      // 6

	fmt.Println("")
	fmt.Println("Map 函数的用法")
	rot13 := func(r rune) rune {
		switch {
		case r == ' ':
			return '-'
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

	fmt.Println("")
	fmt.Println("Repeat 函数的用法")
	fmt.Println("ba" + strings.Repeat("na", 2)) //banana

	fmt.Println("")
	fmt.Println("Replace 函数的用法")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	fmt.Println("")
	fmt.Println("Split 函数的用法")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split("xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Println("")
	fmt.Println("SplitAfter 函数的用法")
	fmt.Printf("%q\n", strings.SplitAfter("/home/m_ta/src", "/")) //["/""home/""m_ta/""src"]

	fmt.Println("")
	fmt.Println("SplitAfterN 函数的用法")
	fmt.Printf("%q\n", strings.SplitAfterN("/home/m_ta/src", "/", 2))  //["/""home/m_ta/src"]
	fmt.Printf("%q\n", strings.SplitAfterN("#home#m_ta#src", "#", -1)) //["/""home/""m_ta/""src"]

	fmt.Println("")
	fmt.Println("SplitN 函数的用法")
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 1))

	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 2))  //["/""home/""m_ta/""src"]
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", -1)) //["""home""m_ta""src"]
	fmt.Printf("%q\n", strings.SplitN("home,m_ta,src", ",", 2))   //["/""home/""m_ta/""src"]

	fmt.Printf("%q\n", strings.SplitN("#home#m_ta#src", "#", -1)) //["/""home/""m_ta/""src"]

	fmt.Println("")
	fmt.Println("Title 函数的用法") //这个函数，还真不知道有什么用
	fmt.Println(strings.Title("her royal highness"))

	fmt.Println("")
	fmt.Println("ToLower 函数的用法")
	fmt.Println(strings.ToLower("Gopher")) //gopher

	fmt.Println("")
	fmt.Println("ToLowerSpecial 函数的用法")

	fmt.Println("")
	fmt.Println("ToTitle 函数的用法")
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("loud 中国"))

	fmt.Println("")
	fmt.Println("Replace 函数的用法")
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", 2)) // aBaACEDF
	//第四个参数小于0，表示所有的都替换， 可以看下golang的文档
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", -1)) // aBaaCEDF

	fmt.Println("")
	fmt.Println("ToUpper 函数的用法")
	fmt.Println(strings.ToUpper("Gopher")) //GOPHER

	fmt.Println("")
	fmt.Println("Trim  函数的用法")
	fmt.Printf("[%q]", strings.Trim("!!! Achtung !!! ", "! ")) // ["Achtung"]

	fmt.Println("")
	fmt.Println("TrimLeft 函数的用法")
	fmt.Printf("[%q]", strings.TrimLeft("!!! Achtung !!! ", "! ")) // ["Achtung !!! "]

	fmt.Println("")
	fmt.Println("TrimSpace 函数的用法")
	fmt.Println(strings.TrimSpace("\t\n a lone gopher \n\t\r\n")) // a lone gopher
}

func BenchmarkBuffer(b *testing.B) {
	res := BufferVal()
	fmt.Println(res.String())
}

func BenchmarkBuilder(b *testing.B) {

}

func TestBufferAndBuilder(t *testing.T) {
	start := time.Now().UnixNano()
	bufferVal := BufferVal()
	fmt.Println(bufferVal.String())
	end := time.Now().UnixNano()
	fmt.Printf("[%s] takes: %d\n", "buffer value", end-start)

	start = time.Now().UnixNano()
	bufferRef := BufferRef()
	fmt.Println(bufferRef)
	end = time.Now().UnixNano()
	fmt.Printf("[%s] takes: %d\n", "buffer ref", end-start)

	start = time.Now().UnixNano()
	builderVal := BuilderVal()
	fmt.Println(builderVal.String())
	end = time.Now().UnixNano()
	fmt.Printf("[%s] takes: %d\n", "build value", end-start)

	start = time.Now().UnixNano()
	builderRef := BuilderRef()
	fmt.Println(builderRef)
	end = time.Now().UnixNano()
	fmt.Printf("[%s] takes: %d\n", "build ref", end-start)
}

func BufferVal() bytes.Buffer {
	var buffer bytes.Buffer
	for i := 0; i < 10; i++ {
		buffer.WriteString("a")
	}
	return buffer
}

func BufferRef() *bytes.Buffer {
	buffer := &bytes.Buffer{}
	for i := 0; i < 10; i++ {
		buffer.WriteString("a")
	}
	return buffer
}

func BuilderVal() strings.Builder {
	builder := strings.Builder{}
	for i := 0; i < 10; i++ {
		builder.WriteString("a")
	}
	return builder
}

func BuilderRef() *strings.Builder {
	builder := &strings.Builder{}
	for i := 0; i < 10; i++ {
		builder.WriteString("a")
	}
	return builder
}
