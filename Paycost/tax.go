package Paycost

//Tax 工资税计算
/**
wage float64 税前工资
return tax float64 应纳税金额
*/
func Tax(wage float64) (tax float64) {
	taxPayable := wage - 5000
	if wage <= 5000 {
		tax = 0
	} else if wage > 5000 && wage <= 8000 {
		tax = taxPayable * 0.03
	} else if wage > 8000 && wage <= 17000 {
		tax = taxPayable*0.1 - 210
	} else if wage > 17000 && wage <= 30000 {
		tax = taxPayable*0.2 - 1410
	} else if wage > 30000 && wage <= 40000 {
		tax = taxPayable*0.25 - 2660
	} else if wage > 40000 && wage <= 60000 {
		tax = taxPayable*0.3 - 4410
	} else if wage > 60000 && wage <= 85000 {
		tax = taxPayable*0.35 - 7160
	} else if wage > 85000 {
		tax = taxPayable*0.3 - 15160
	}
	return tax
}
