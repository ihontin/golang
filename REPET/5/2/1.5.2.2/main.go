package main

import (
	"fmt"
	"math/big"
)

// Factorial должна вычислять факториал числа и возвращать результат.
func Factorial(n *int) int {
	var bigNum = big.NewInt(1)

	for i := 1; i <= *n; i++ {
		bigNum.Mul(bigNum, big.NewInt(int64(i)))
	}
	tryBest := bigNum.Int64()
	return int(tryBest)
}

// isPalindrome должна проверить, является ли строка палиндромом, и вернуть результат проверки.
func isPalindrome(s *string) bool {
	checkS := *s
	for i := 0; i < len(checkS)/2; i++ {
		if checkS[i] != checkS[len(checkS)-1-i] {
			return false
		}
	}
	return true
}

// CountOccurrences должна вернуть количество вхождений числа в срез.
func CountOccurrences(numbers *[]int, target *int) int {
	var count int
	for _, num := range *numbers {
		if num == *target {
			count++
		}
	}
	return count
}

// ReverseString должна развернуть строку и вернуть результат.
func ReverseString(str *string) string {
	var reversStr string
	for _, letter := range *str {
		reversStr = string(letter) + reversStr
	}
	return reversStr
}

const (
	first = iota + 1
	sec
	th
	fi
	six
	seven
	eh
)

func main() {
	var a = 6
	fmt.Println("Factorial 6: ", Factorial(&a))
	var strNew, taskFour = "fwo7!#!7owf", "something unusual"
	fmt.Println(len(strNew), isPalindrome(&strNew))
	fmt.Println(ReverseString(&taskFour))
	fmt.Println(seven, eh)

}
