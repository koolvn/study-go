package main

import (
	"fmt"
	"math/big"
)

func main() {
	pow36_17 := new(big.Int).Exp(big.NewInt(36), big.NewInt(17), nil)

	for x := int64(0); x < 1000; x++ {
		z := big.NewInt(0)
		pow6_x := new(big.Int).Exp(big.NewInt(6), big.NewInt(x), nil)
		z.Sub(pow36_17, pow6_x)
		z.Add(z, big.NewInt(71))
		sum := big.NewInt(0)
		for z.Cmp(big.NewInt(0)) > 0 {
			rem := new(big.Int).Mod(z, big.NewInt(6))
			sum.Add(sum, rem)
			z.Div(z, big.NewInt(6))
		}

		if sum.Cmp(big.NewInt(61)) == 0 {
			fmt.Printf("Found x: %d with answer: %s\n", x, sum)
			break
		}
	}
}
