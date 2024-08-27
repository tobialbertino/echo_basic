package fruits

import "echo_basic/infra/fruitApi"

func RemoveSLiceFruitByID(s []fruitApi.DataFruit, idx int) []fruitApi.DataFruit {
	s[idx] = s[len(s)-1]
	return s[:len(s)-1]
}
