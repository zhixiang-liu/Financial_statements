package Wage

//Insurance 计算五险一金
/**
parameter wage float64	税前工资
return	五险一金
*/
func Insurance(wage float64) float64 {

	return wage * 0.15
}
