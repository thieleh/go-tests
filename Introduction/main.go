package main

import (
	"fmt"
	"tests-introduction/address"
)

func main() {
	typeOfAddress := address.TypeOfAddress("Rod Random")
	fmt.Println(typeOfAddress)
}
