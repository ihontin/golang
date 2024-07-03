package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func genFakeData() string {
	name := fake.FullName()
	address := fake.StreetAddress()
	phone := fake.Phone()
	email := fake.EmailAddress()
	return fmt.Sprintf("Name: %v, Address: %v, Phone: %v, Email: %v", name, address, phone, email)
}

func main() {
	fmt.Println(genFakeData())
}
