package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a, b string) string {
	num1 := new(big.Int)
	num1.SetString(a, 2)
	num2 := new(big.Int)
	num2.SetString(b, 2)

	return num2.Add(num1, num2).Text(2)
}
