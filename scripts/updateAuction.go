package scripts

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/csjh/NetworthMeta-Backend/db"
	"github.com/csjh/NetworthMeta-Backend/generators"
	"github.com/csjh/NetworthMeta-Backend/models"
	"github.com/csjh/NetworthMeta-Backend/types"
	"github.com/csjh/NetworthMeta-Backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"regexp"
	"strings"
)

var auctions = map[string][]types.FormattedAuction{}

func ShowAuctions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(auctions)
	if err != nil {
		return
	}
}

func ReloadAuctionData() {
	auctionPage := getAuctionsEnded()
	processAuctions(auctionPage)
	updateAuctions()
}

func getAuctionsEnded() []types.AuctionsEnded {
	// fetch page
	response, err := http.Get("https://api.hypixel.net/skyblock/auctions_ended")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
		return nil
	}

	// marshall body into struct
	var auctionPage types.AuctionsEndedPage
	err = json.Unmarshal(body, &auctionPage)
	if err != nil {
		panic(err)
		return nil
	}
	return auctionPage.Auctions
}

var re = regexp.MustCompile(`§.`)

func getAttributes(base types.ItemNBT) (id string, name string, isRecombobulated bool) {
	item := base.Tag.ExtraAttributes
	itemId := item.ID
	itemName := re.ReplaceAllString(base.Tag.Display.Name, "")
	if itemId == "ENCHANTED_BOOK" && item.Enchantments != nil {
		for enchant, level := range item.Enchantments {
			if len(item.Enchantments) == 1 {
				itemId = fmt.Sprintf("%s_%d", enchant, level)
				itemName = utils.Capitalize(fmt.Sprintf("%s %d", enchant, level))
			}
		}
	}

	if itemId == "ATTRIBUTE_SHARD" && item.Attributes != nil {
		for attribute, level := range item.Attributes {
			if len(item.Attributes) == 1 {
				itemId = fmt.Sprintf("%s_%d", attribute, level)
				itemName = utils.Capitalize(fmt.Sprintf("%s %d", attribute, level))
			}
		}
	}

	if itemId == "RUNE" && item.Runes != nil {
		for runeName, level := range item.Runes {
			if len(item.Runes) == 1 {
				itemId = fmt.Sprintf("%s_%d", runeName, level)
				itemName = utils.Capitalize(fmt.Sprintf("%s %d", runeName, level))
			}
		}
	}

	if itemId == "PET" {
		var pet types.Pet
		err := json.Unmarshal([]byte(item.PetInfo), &pet)
		if err != nil {
			panic(err)
			return "", "", false
		}
		_, level := generators.CalculateSkillLevel(pet)
		if level == 1 || level == 100 || level == 200 {
			itemId = fmt.Sprintf("lvl_%d_%s_%s", level, pet.Tier, pet.Type)
			itemName = fmt.Sprintf("[Lvl %d] %s %s", level, utils.Capitalize(pet.Tier), utils.Capitalize(pet.Type))
		}
	}
	return itemId, itemName, item.RarityUpgrades != 0
}

func processAuctions(auctionsEnded []types.AuctionsEnded) {
	for _, auction := range auctionsEnded {
		item, err := utils.DecodeNBT(auction.ItemBytes)
		if err != nil {
			panic(err)
		}

		id, name, isRecombobulated := getAttributes(item)

		format := types.FormattedAuction{
			Id:             strings.ToUpper(id),
			Name:           utils.Capitalize(name),
			Price:          auction.Price,
			Seller:         auction.Seller,
			Ending:         auction.Timestamp,
			Count:          item.Count,
			Recombobulated: isRecombobulated,
			EstimatedPrice: generators.ParseItem(item).Price,
		}

		auctions[id] = append(auctions[id], format)
	}
}

func updateAuctions() {
	client := db.Connect()
	defer db.Disconnect(client)

	collection := client.Database("networthmeta").Collection("auctions")

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
		return
	}
	var rawDocuments []models.Auction
	err = cur.All(context.Background(), &rawDocuments)
	if err != nil && err != mongo.ErrNoDocuments {
		panic(err)
		return
	}
	documents := make(map[string]models.Auction)
	for _, document := range rawDocuments {
		documents[document.ID] = document
	}

	var writeModels []mongo.WriteModel
	for item, itemAuctions := range auctions {
		mongoData := documents[item]

		var recombobulatedAuctions, unrecombobulatedAuctions []types.FormattedAuction
		for _, auction := range itemAuctions {
			if auction.Recombobulated {
				recombobulatedAuctions = append(recombobulatedAuctions, auction)
			} else {
				unrecombobulatedAuctions = append(unrecombobulatedAuctions, auction)
			}
		}

		mongoData.ID = item

		mongoData.Auctions = append(mongoData.Auctions, unrecombobulatedAuctions...)
		minIdx := int(math.Max(float64(len(mongoData.Auctions))-10, 0))
		mongoData.Auctions = mongoData.Auctions[minIdx:]

		mongoData.RecombobulatedAuctions = append(mongoData.RecombobulatedAuctions, recombobulatedAuctions...)
		minIdx = int(math.Max(float64(len(mongoData.RecombobulatedAuctions))-10, 0))
		mongoData.RecombobulatedAuctions = mongoData.RecombobulatedAuctions[minIdx:]

		var pricesAdjustedByCount []float64
		for _, auction := range mongoData.Auctions {
			pricesAdjustedByCount = append(pricesAdjustedByCount, auction.Price/float64(auction.Count))
		}
		if len(pricesAdjustedByCount) > 0 {
			mongoData.Price = utils.Median(pricesAdjustedByCount)
		}

		pricesAdjustedByCount = []float64{}
		for _, auction := range mongoData.RecombobulatedAuctions {
			pricesAdjustedByCount = append(pricesAdjustedByCount, auction.Price/float64(auction.Count))
		}
		if len(pricesAdjustedByCount) > 0 {
			mongoData.RecombobulatedPrice = utils.Median(pricesAdjustedByCount)
		} else {
			mongoData.RecombobulatedPrice = mongoData.Price + 4000000
		}

		writeModels = append(writeModels, mongo.NewReplaceOneModel().SetUpsert(true).SetFilter(bson.M{"id": mongoData.ID}).SetReplacement(mongoData))
	}

	_, err = collection.BulkWrite(context.Background(), writeModels)
	if err != nil {
		panic(err)
		return
	}

	auctions = make(map[string][]types.FormattedAuction)
}
