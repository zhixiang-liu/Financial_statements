package Wage

import "Financial_statements/utils"

// GrossSalary 计算税前工资
/**
  param p map[string][]string	map变量名，存放用户的基本信息
  param key string			用户基本信息索引
  return wage float64			税前工资
*/
func GrossSalary(p map[string][]string, key string) (wage float64) {
	var allowance int64
	var wage1 int64
	years := utils.Calc(p[key][3])          //计算工龄
	classfees := utils.ClassFees(p[key][2]) //计算课时费
	key = p[key][1]                         //将职位信息赋给x
	if key == "A" {                         //A为行政
		allowance = years * 500              //津贴
		wage1 = 4500 + allowance + classfees //工资
	} else if key == "T" { //T为教师
		allowance = years * 500
		wage1 = 5000 + allowance + classfees
	} else if key == "D" { //D为教导主任
		allowance = years * 800
		wage1 = 5000 + allowance + classfees
	}
	wage = float64(wage1)
	return wage
}
