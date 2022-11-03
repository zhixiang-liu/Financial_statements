package utils

import (
	"fmt"
	"strconv"
)

//ClassFees 计算课时费
/**
num string 当前map
return classfees int64
*/
func ClassFees(num string) (classfees int64) {
	if num == "0" {
		return 0
	} else {
		num1, err := strconv.Atoi(num) //将string类型转化为int类型
		if err != nil {
			println("错误出在ClassFees")
			fmt.Printf(err.Error()) //错误处理
		}
		classfees = int64(num1) * 40 //计算课时费
		return classfees
	}
}
