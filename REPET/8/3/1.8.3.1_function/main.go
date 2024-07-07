package main

import "fmt"

func UserInfo(name string, age int, cities ...string) string {
	var allCity string
	for i, cit := range cities {
		if i != 0 {
			allCity += ", "
		}
		allCity += cit
	}
	return fmt.Sprintf("Имя: %s, возраст: %d, города: %v", name, age, allCity)
}

func main() {
	fmt.Println(UserInfo("Griga", 2, "Benglad", "Urban", "Impal"))

}
