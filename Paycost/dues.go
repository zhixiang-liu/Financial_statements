package Paycost

//Dues 计算党费
/**
afterTaxWages float64	税后工资
p map[string][]string	map变量名，存放用户的基本信息
key string	用户基本信息索引
*/
func Dues(afterTaxWages float64, p map[string][]string, key string) (dues float64) {
	if p[key][5] == "Y" { //判断是否是党员
		dues = PayStandard(afterTaxWages)
	} else if p[key][5] == "N" {
		dues = 0
	}
	return dues
}

// PayStandard 	缴纳标准
/**
afterTaxWages float64	税前工资
return dues float64 缴纳金额
*/
func PayStandard(afterTaxWages float64) (dues float64) { //缴纳标准
	if afterTaxWages <= 3000 && afterTaxWages > 0 {
		dues = afterTaxWages * 0.005
	} else if afterTaxWages > 3000 && afterTaxWages <= 5000 {
		dues = afterTaxWages * 0.01
	} else if afterTaxWages > 5000 && afterTaxWages <= 10000 {
		dues = afterTaxWages * 0.015
	} else if afterTaxWages > 10000 {
		dues = afterTaxWages * 0.02
	}
	return dues
}
