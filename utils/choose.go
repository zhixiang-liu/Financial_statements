package utils

// Choose 选择map索引
/**
param num int	功能数字，范围1-10
return no string	map索引
*/
func Choose(num int) (no string) {
	switch num {
	case 1:
		no = "1谢希仁"
	case 2:
		no = "2吴功宜"
	case 3:
		no = "3陈鸣"
	case 4:
		no = "4吴英"
	case 5:
		no = "5徐恪"
	case 6:
		no = "6徐明伟"
	case 7:
		no = "7陈威兵"
	case 8:
		no = "8张传福"
	case 9:
		no = "9张刚林"
	case 10:
		no = "10赵立英"
	}
	return no
}
