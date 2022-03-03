package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Printf("asdf1")
	target := big.NewInt(1)
	fmt.Printf("%v\n", target)
	target.Lsh(target, uint(256-24))
	fmt.Printf("%v\n", target)
	fmt.Printf("asdf")
	fmt.Printf("asdf2")
	fmt.Println("asdf")
}
