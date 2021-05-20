package main

type Card struct {
	Id   int
	Name string
}

type CaravanCard struct {
	Card  Card
	Cubes Cube
	Limit int
	Order int
}

type ScoreCard struct {
	Card  Card
	Cubes Cube
	Value int
}

type MerchantCard struct {
	Card        Card
	InputCubes  Cube
	OutputCubes Cube
}

type MerchantDeck struct {
	MerchantCard MerchantCard
}

type Cube struct {
	yellowQty int
	redQty    int
	greenQty  int
	brownQty  int
}

type Coin struct {
	Silver struct {
		Qty int
	}

	Gold struct {
		Qty int
	}
}
