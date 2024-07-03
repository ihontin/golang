package main

import "fmt"

func ReverseString(s *string) {
	var newS string
	for _, runStr := range *s {
		newS += string(runStr)
	}
	*s = newS
}

func main() {
	myString := "Mama mila ramu"
	fmt.Println(myString)
	ReverseString(&myString)
	fmt.Println(myString)

}
