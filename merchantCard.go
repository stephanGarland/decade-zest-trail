package main

import "fmt"

func deck() {
	merchantDeck := []MerchantCard{
		{
			Card: Card{
				Id:   0,
				Name: "0-1G1Y",
			},
			OutputCubes: Cube{
				yellowQty: 1,
				greenQty:  1,
			},
		},
		{
			Card: Card{
				Id:   1,
				Name: "2G-2B",
			},
			InputCubes: Cube{
				greenQty: 2,
			},
			OutputCubes: Cube{
				brownQty: 2,
			},
		},
	}
	fmt.Println(merchantDeck)
}
