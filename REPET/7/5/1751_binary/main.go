package main

import (
	"fmt"
	"log"
	"strconv"
	"unsafe"
)

/*
Преобразовать тип данных можно несколькими способами, например:
float32(s)
*(*float32)(unsafe.Pointer(&s))
*/

// — Функция должна корректно конвертировать двоичную строку в число с плавающей точкой типа float32
// — Функция должна вернуть число 0.15625 для строки “00111110001000000000000000000000”
func binaryStringToFloat(binary string) float32 {
	//strconv.ParseInt(строка, система счисления, тип данных)
	s, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("int64: %v \n", s)

	number := *(*uint32)(unsafe.Pointer(&s)) // преобразование в тип uint32(s) покажет тот же результат
	fmt.Printf("uint32(s): %v, *(*uint32)(unsafe.Pointer(&s)): %v\n", uint32(s), number)

	// преобразование float32(s) покажет другой результат
	fmt.Printf("float32(s): %v\n", float32(s))

	floatNumber := *(*float32)(unsafe.Pointer(&s))
	return floatNumber
}

func main() {
	b := binaryStringToFloat("00111110001000000000000000000000")
	fmt.Printf("*(*float32)(unsafe.Pointer(&s)): %v\n", b)

}
