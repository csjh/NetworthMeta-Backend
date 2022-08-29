package types

type BazaarEntry struct {
	ProductId   string  `json:"product_id"`
	SellSummary []Order `json:"sell_summary"`
	BuySummary  []Order `json:"buy_summary"`
	QuickStatus struct {
		ProductId      string  `json:"productId"`
		SellPrice      float64 `json:"sellPrice"`
		SellVolume     int64   `json:"sellVolume"`
		SellMovingWeek int64   `json:"sellMovingWeek"`
		SellOrders     int64   `json:"sellOrders"`
		BuyPrice       float64 `json:"buyPrice"`
		BuyVolume      int64   `json:"buyVolume"`
		BuyMovingWeek  int64   `json:"buyMovingWeek"`
		BuyOrders      int64   `json:"buyOrders"`
	} `json:"quick_status"`
}

type Order struct {
	Amount       int     `json:"amount"`
	PricePerUnit float64 `json:"pricePerUnit"`
	Orders       int     `json:"orders"`
}

type Bazaar struct {
	Success     bool                   `json:"success"`
	LastUpdated int64                  `json:"lastUpdated"`
	Products    map[string]BazaarEntry `json:"products"`
}
