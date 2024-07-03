package main

import (
	"fmt"
	"math"
)

// decimalPlaces количество сдвига под округление
// вернуть true, если округленные значения равны
// и вернуть разницу между округленными значениями.
func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	moove := math.Pow(10, float64(decimalPlaces)) //возведет 10 в степень decimalPlaces

	//сместит точку на значение moove, округлит и сместит точку обратно
	x := math.Round(a*moove) / moove
	y := math.Round(b*moove) / moove
	fmt.Println("Изначально:", a, b, "Округлено:", x, y)
	return x == y, math.Abs(x - y)
}

func main() {
	var a, b float64 = 0.258, 0.264
	var dec int = 2
	// fmt.Scan(&a)
	// fmt.Scan(&b)
	// fmt.Scan(&dec)
	fmt.Println(CompareRoundedValues(a, b, dec))
}
