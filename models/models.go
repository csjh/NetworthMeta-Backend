package models

import (
	"github.com/csjh/NetworthMeta-Backend/types"
)

type Bazaar struct {
	ID    string  `json:"id,omitempty" bson:"id"`
	Price float64 `json:"price,omitempty" bson:"price"`
}

type Auction struct {
	ID                     string                   `json:"id,omitempty" bson:"id"`
	Auctions               []types.FormattedAuction `json:"past_auctions,omitempty" bson:"past_auctions"`
	RecombobulatedAuctions []types.FormattedAuction `json:"past_recombobulated_auctions,omitempty" bson:"past_recombobulated_auctions"`
	Price                  float64                  `json:"price,omitempty" bson:"price"`
	RecombobulatedPrice    float64                  `json:"recombobulated_price,omitempty" bson:"recombobulated_price"`
}
