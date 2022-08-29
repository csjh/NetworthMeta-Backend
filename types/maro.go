package types

type Player map[string]Profile

type Profile struct {
	ChestCache       map[string][]ItemNBT `json:"chestCache"`
	MiscStorageCache map[string][]ItemNBT `json:"miscStorageCache"`
	MiscCoins        map[string]float64   `json:"miscCoins"`
}

type Pet struct {
	Price     float64 `json:"price"`
	Type      string  `json:"type"`
	Active    bool    `json:"active"`
	Exp       float64 `json:"exp"`
	Tier      string  `json:"tier"`
	HideInfo  bool    `json:"hideInfo"`
	HeldItem  string  `json:"heldItem"`
	CandyUsed int     `json:"candyUsed"`
	Uuid      string  `json:"uuid"`
	Skin      string  `json:"skin"`
}

type CalculatedItem struct {
	Id     string
	Name   string
	Price  float64
	Recomb bool
	Count  int
}

type SackItem struct {
	Name  string
	Count int
	Price float64
}

type Category struct {
	Total float64
	Items []CalculatedItem
}

type Networth struct {
	Categories map[string]Category
	Sacks      struct {
		Total float64
		Items []SackItem
	}
	Networth float64
}

type Output struct {
	Sacks struct {
		Total float64
		Items []SackItem
	}
	MiscStorage map[string][]ItemNBT
	Chests      []ItemNBT
	Coins       float64
}
