package utils

import (
	"fmt"
	"time"
)

func StringsTimeUnix(s string) int64 {

	loc, _ := time.LoadLocation("Local")                 //获取本地时区
	t, err := time.ParseInLocation("2006-01-02", s, loc) //将字符串转化成时间
	if err != nil {                                      //错误处理
		println("错误出在StringsTimeUnix")
		fmt.Printf(err.Error())
		return 0
	}
	return t.Unix() //返回处理好的秒

}
