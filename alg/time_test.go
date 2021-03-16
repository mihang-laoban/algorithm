package alg

import (
	"dp/tools"
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// unix to string
	res := tools.Unix2TimeStr(time.Now().Unix())
	fmt.Println(res)
	// string to unix
	fmt.Println(tools.TimeStr2Unix(res))
}

func timeSample() {
	t := time.Now().Unix()      //外部传入的时间戳（秒为单位），必须为int64类型
	t1 := "2019-01-08 13:50:30" //外部传入的时间字符串

	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型 链接符号可变
	timeTemplate2 := "2006-01-02"          //其他类型 链接符号可变
	timeTemplate3 := "15:04:05"            //其他类型 链接符号可变

	// ======= 将时间戳格式化为日期字符串 =======
	fmt.Println("unix to string:")
	res := time.Unix(t, 0).Format(timeTemplate1)
	fmt.Println(res)                                   //输出：2019-01-08 13:50:30
	fmt.Println(time.Unix(t, 0).Format(timeTemplate2)) //输出：2019-01-08
	fmt.Println(time.Unix(t, 0).Format(timeTemplate3)) //输出：13:50:30
	fmt.Println()
	fmt.Println("string to unix:")
	// ======= 将时间字符串转换为时间戳 =======
	stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local)   //使用parseInLocation将字符串格式化返回本地时区时间
	stamp2, _ := time.ParseInLocation(timeTemplate1, res, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	fmt.Println(stamp.Unix())                                         //输出：1546926630
	fmt.Println(stamp2.Unix())                                        //输出：1546926630
}
