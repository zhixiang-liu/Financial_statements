package Wage

import (
	"Financial_statements/Paycost"
	"Financial_statements/information"
	"Financial_statements/utils"
	"fmt"
)

func Operation() {
	var num int
	for num = 1; num <= 10; num++ {
		key := utils.Choose(num)                                               //选择计算数据
		grossSalary := GrossSalary(information.M1(), key)                      //税前工资
		insurance := Insurance(grossSalary)                                    //五险一金
		tax := Paycost.Tax(grossSalary - insurance)                            //工资税金
		dues := Paycost.Dues(grossSalary-tax-insurance, information.M1(), key) //党费缴纳
		accommdation := Paycost.Accommodation(information.M1(), key)           //住宿
		spending := tax + accommdation + dues + insurance                      //综合支出
		realWages := grossSalary - spending                                    //实发工资
		fmt.Printf("%v的税前工资为:%.2f,五险一金:%.2f,交税:%.2f,党费:%.2f,实发工资:%.2f\n", key, grossSalary, insurance, tax, dues, realWages)
	}

}
