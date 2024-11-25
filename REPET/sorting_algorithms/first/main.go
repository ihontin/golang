package main

import "fmt"

//сортировка вставками
func insertSort(co []int) []int {
	for i := 0; i < len(co); i++ {
		for j := i; j > 0 && co[j-1] > co[j]; j-- {
			co[j], co[j-1] = co[j-1], co[j]
		}
	}
	return co
}

// сортировка выбором

func main() {
	fmt.Println(insertSort([]int{21, 4, 89, 38, 11}))
}
