package time

import (
	"fmt"
	"testing"
	"time"
)
func TestDemo1(t *testing.T) {
	now:=time.Now()
	fmt.Printf("current time:%v\n",now)
	
	year:=now.Year()
	month:=now.Month()
	day:=now.Day()
	hour:=now.Hour()
	minute:=now.Minute()
	second:=now.Second()
	fmt.Println(year,month,day,hour,minute,second)
}

//获取timestamp
func TestDemo2(t *testing.T) {
	now:=time.Now()
	timestamp:=now.Unix()
	milin:=now.UnixMilli()
	micro:=now.UnixMicro()
	namiao:=now.UnixNano()
	fmt.Println(timestamp,milin,micro,namiao)
}

//timestamp 转时间
func TestDemo3(t *testing.T) {
	ret:=time.Unix(1654570899, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Month())
}

//当前时间的加减操作
func TestAdd(t *testing.T) {
	fmt.Println(time.Second)
	now:=time.Now()
	fmt.Println(now.Add(24*time.Hour))
	fmt.Println(now.Add(10*time.Minute))
}

//时间格式转换
func TestConvert(t *testing.T){
	now:=time.Now()
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	fmt.Println(now.Format("2006/01/02 03:04:05.000"))
}


//解析时间格式
func TestTimestamp(t *testing.T){
	timeObj,err:=time.Parse("2006-01-02","2022-06-07")
	if err != nil {
		fmt.Printf("parse time failed err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
}

