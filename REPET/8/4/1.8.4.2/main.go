package main

import "fmt"

// FindMaxAndMin которая принимает неопределенное количество целочисленных аргументов
// и возвращает два значения: максимальное и минимальное значение из переданных аргументов.
func FindMaxAndMin(args ...int) (int, int) {
	if len(args) == 0 {
		return 0, 0
	}
	var min, max = args[0], args[0]
	for _, num := range args {
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}
	return min, max
}

func main() {
	min, max := FindMaxAndMin(0, -1)
	fmt.Printf("min = %d, max = %d", min, max)
}
