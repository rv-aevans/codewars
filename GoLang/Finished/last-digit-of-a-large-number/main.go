package main

import (
	"fmt"
	"math/big"
)

func main() {
	// fmt.Println(LastDigit("552614859038286924289967931601074326821347888193597702180855998140460661877030689097413855173895670915021503134270420498345228497541089617687338500195877220254307466814180579109977672374655625992810851781350660840820640344681303916", "1376118368364296115916170111338760549242646603"))
	fmt.Println(LastDigit("4", "124"))
}

func LastDigit(n1, n2 string) int {
	d1, d2 := new(big.Int), new(big.Int)
	d1, _ = d1.SetString(n1, 0)
	d2, _ = d2.SetString(n2, 0)
	ten := big.NewInt(10)
	d1.Exp(d1, d2, ten)
	return int(d1.Int64())

	// fmt.Println(d1.Int64())

	// s := fmt.Sprintf("%d", d1.Int64())[:11]

	// fmt.Println(s)

	// m3, _ := strconv.Atoi(string(s[10]))
	// m4, _ := strconv.Atoi(string(s[9]))

	// if m3 >= 5 {
	// 	m4++
	// }

	// fmt.Println(m4)

	// if n2 == "0" {
	// 	return 1
	// }

	// m1, _ := strconv.Atoi(n1)
	// m2, _ := strconv.Atoi(n2)

	// if len(n1) < 2 || len(n2) < 2 {
	// 	return int(math.Pow(float64(m1), float64(m2))) % 10
	// }

	// d1, d2 := new(big.Int), new(big.Int)
	// d1, _ = d1.SetString(n1, 10)
	// d2, _ = d2.SetString(n2, 10)
	// one := big.NewInt(1)
	// four := big.NewInt(4)
	// ten := big.NewInt(10)
	// d2.Exp(d2, one, four)
	// if d2.Int64() == 0 {
	// 	d1.Exp(d1, d2, ten).Int64()
	// 	return int(math.Abs(float64(d1.Int64())))
	// } else {
	// 	return int(math.Abs(float64(d2.Int64())))
	// }
	// return 0
}
