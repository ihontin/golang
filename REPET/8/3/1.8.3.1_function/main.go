package main

import "fmt"

func UserInfo(name []int) int {
	var allCity = 0
	for i := 0; i < len(name); i++ {
		if name[i]%2 == 0 {
			allCity += name[i]
		} else {
			allCity -= name[i]
		}

	}
	return allCity
}

func main() {
	fmt.Println(UserInfo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))

}
