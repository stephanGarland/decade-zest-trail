package main

func scoreDeck() []ScoreCard {
	scoreDeck := []ScoreCard{
		{
			Card: Card{
				Id:   0,
				Name: "2G3B18",
			},
			Cubes: Cube{
				greenQty: 2,
				brownQty: 3,
			},
			Value: 18,
		},
	}
	return scoreDeck
}
