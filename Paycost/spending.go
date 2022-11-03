package Paycost

//Accommodation 判断是否住宿
/**
p map[string][]string	map变量名，存放用户的基本信息
key string	用户基本信息索引
return float64 住宿费
*/
func Accommodation(p map[string][]string, key string) float64 {
	var accommodation float64
	if p[key][4] == "Y" { //住宿费
		accommodation = 300
	} else {
		accommodation = 0
	}
	return accommodation
}
