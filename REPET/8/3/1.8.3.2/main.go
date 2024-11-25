package main

import "fmt"

//Функция должна объединять строки через разделитель в порядке их следования и возвращать строку в формате
//“even: some string, odd: some string”,где even - количество строк с четным индексом,
//odd - количество строк с нечетным индексом. Индексация начинается с 0.

func ConcStr(sep string, manyStr ...string) string {
	var firstEv, firstOd = true, true
	ev, od := "even: ", "odd: "
	for i := 0; i < len(manyStr); i++ {
		if (i+2)%2 == 0 {
			if firstEv {
				ev += manyStr[i]
				firstEv = false
			}
			ev += sep + manyStr[i]
			continue
		}
		if firstOd {
			od += manyStr[i]
			firstOd = false
		}
		od += sep + manyStr[i]

	}
	return ev + ", " + od
}

func main() {
	fmt.Println(ConcStr("*-*", "q", "E", "q", "E", "q", "E", "q", "E", "q", "E", "q",
		"E", "q", "E", "q", "E", "q", "E", "q", "E", "q", "E", "q"))
}
