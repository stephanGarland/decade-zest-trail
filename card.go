package main

type Card struct {
	Id int `json:"id"`
}

type CaravanCard struct {
	Card  Card
	Cubes Cube
	Limit int
	Order int
}

type PointCard struct {
	Card  Card
	Cubes Cube
	Value int
	cubes string `json:"-"`
}

type MerchantCard struct {
	Card         Card
	InputCubes   Cube
	OutputCubes  Cube
	UpgradeCubes UpgradeCubes
}

type Cube struct {
	YellowQty     int `json:"yellowQty,omitempty"`
	RedQty        int `json:"redQty,omitempty"`
	GreenQty      int `json:"greenQty,omitempty"`
	BrownQty      int `json:"brownQty,omitempty"`
	FromYellowQty int `json:"fromYellowQty,omitempty"`
	FromRedQty    int `json:"fromRedQty,omitempty"`
	FromGreenQty  int `json:"fromgreenQty,omitempty"`
	FromBrownQty  int `json:"fromBrownQty,omitempty"`
	ToYellowQty   int `json:"toYellowQty,omitempty"`
	ToRedQty      int `json:"toRedQty,omitempty"`
	ToGreenQty    int `json:"toGreenQty,omitempty"`
	ToBrownQty    int `json:"toBrownQty,omitempty"`
}

type Coin struct {
	SilverQty int
	GoldQty   int
}

type UpgradeCubes struct {
	UpgradeAmount int
}
