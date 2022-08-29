package types

type AuctionsEndedPage struct {
	Success     bool            `json:"success"`
	LastUpdated int             `json:"last_updated"`
	Auctions    []AuctionsEnded `json:"auctions"`
}

type AuctionsEnded struct {
	AuctionId     string  `json:"auction_id"`
	Seller        string  `json:"seller"`
	SellerProfile string  `json:"seller_profile"`
	Buyer         string  `json:"buyer"`
	Timestamp     int64   `json:"timestamp"`
	Price         float64 `json:"price"`
	Bin           bool    `json:"bin"`
	ItemBytes     string  `json:"item_bytes"`
}

type FormattedAuction struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Seller         string  `json:"seller"`
	Ending         int64   `json:"ending"`
	Count          int     `json:"count"`
	Recombobulated bool    `json:"recombobulated"`
	EstimatedPrice float64 `json:"estimatedPrice"`
}
