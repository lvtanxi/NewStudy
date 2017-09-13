//时间
package main

import (
	"time"
	"fmt"
)

func main() {
	//定义一个FORMAT
	DATE_FORMAT := "2006-01-02 15:04:05"
	//获取当前时间
	t := time.Now()
	fmt.Println(t.Unix())
	//获取年月日
	year, month, day := t.Date()
	dateStr := fmt.Sprintf("%d-%d-%d 00:00:00", year, month, day)
	fmt.Println(dateStr)
	//转换日期格式
	fmt.Println(t.Format(DATE_FORMAT)) //我不知道
	fmt.Println(t.Format(time.RFC3339Nano))

	//时间计算
	//获取明天时间
	tomorrow :=t.Add(24*time.Hour)
	fmt.Println("明天是",tomorrow.Format(DATE_FORMAT))
	//时间戳转时间
	getTime :=time.Unix(1505290407,0)
	fmt.Println(getTime.Format(DATE_FORMAT))
	//时间比较
	fmt.Println(t.After(getTime))



	/*fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	fmt.Println(t.Format(time.RFC3339Nano))*/
}
