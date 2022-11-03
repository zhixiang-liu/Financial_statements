package utils

import (
	"time"
)

//Calc 计算入职年数
/**
t string 入职时间(含年月日)
return years int64 入职年数
*/
func Calc(t string) (years int64) {

	currentTime := time.Now().Unix() //返回当前时间
	inTheTime := StringsTimeUnix(t)  //返回入职时间
	years = (currentTime - inTheTime) / 31536000
	return years
}
