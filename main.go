package main

import (
	"fmt"
	"math/big"
)

func main() {
	target := big.NewInt(1)
	fmt.Printf("%v\n", target)
	target.Lsh(target, uint(256-24))
	fmt.Printf("%v\n", target)
}
