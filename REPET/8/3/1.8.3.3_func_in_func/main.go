package main

import "fmt"

func PrintNums(randomNum ...int) {
	for _, num := range randomNum {
		fmt.Println(num)
	}

}

func main() {

	PrintNums(1, 2)

	step1 := func(x int) func(int) func(int) int {
		return func(y int) func(int) int {
			return func(z int) int {
				return x * y * z
			}
		}

	}
	var (
		a = 2
		b = 3
		c = 4
	)

	fmt.Println(step1(a)(b)(c))
	res := step1(b)
	fmt.Println(res(b)(a))

}
