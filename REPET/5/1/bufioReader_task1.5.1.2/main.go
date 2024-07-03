package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1 Add должна принимать два целых числа и возвращать указатель на результат их сложения.
func Add(a, b int) *int {
	res := a + b
	return &res
}

// 2 Max должна принимать срез целых чисел и возвращать указатель на максимальное число в этом срезе.
func Maxnew(s []int) *int {
	res := 0
	for i := range s {
		if s[i] > s[res] {
			res = i
		}
	}
	return &s[res]
}

// 3 IsPrime должна принимать целое число и возвращать указатель на логическое значение, которое указывает, является ли число простым.
func IsPrime(num int) *bool {
	var res = true
	if num < 2 {
		res = false
		return &res
	} else if num == 2 {
		return &res
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			res = false
			break
		}
	}
	return &res
}

// 4
// ConcatenateStrings должна принимать срез строк и возвращать указатель на строку, которая является результатом их конкатенации.
func ConcatenateStrings(s []string) *string {
	res := strings.Join(s, "")
	return &res
}

func main() {
	// 1
	//var a, b int
	// fmt.Scanln(&a, &b)
	// n := Add(a, b)
	// fmt.Println(*n)

	// 2
	// запись в одну строку: bufio.NewReader(os.Stdin).ReadString('\n')
	reader := bufio.NewReader(os.Stdin)  // создаем экземпляр bufio.Reader который может читать из файла и консоли,
	text, err := reader.ReadString('\n') // считывает до обнаруженния символа '\n'

	if err != nil {
		fmt.Errorf("ReadString unsucsessful: %v", err) //плейсхолдер %v универсальный
	}
	text = strings.TrimSuffix(text, "\n") // убрать символ с конца
	text = strings.TrimSpace(text)        // убрать пробелы с вначале и конце
	sliceText := strings.Split(text, " ") // делит строку на элементы, в данном случае по пробелам
	fmt.Printf("sliceText: %v, %T \n", sliceText, sliceText)

	var sliceInteger = make([]int, len(sliceText))

	for i, num := range sliceText {
		convert, err := strconv.Atoi(num)
		sliceInteger[i] = convert

		if err != nil {
			fmt.Errorf("strconv.Atoi unsucsessful: %v", err)
		}
	}

	fmt.Printf(" sliceInteger: %v, %T, max integer: %d\n", sliceInteger, sliceInteger, *Maxnew(sliceInteger))

	//3
	isItPrime := IsPrime(23)
	fmt.Println("Is 23 prime: ", *isItPrime)

	//4
	strFromSlice := *ConcatenateStrings(sliceText)
	fmt.Printf("string from slice: %s, Tipe: %T\n", strFromSlice, strFromSlice)
}
