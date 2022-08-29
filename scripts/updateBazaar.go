package scripts

import (
	"context"
	"encoding/json"
	"github.com/csjh/NetworthMeta-Backend/db"
	"github.com/csjh/NetworthMeta-Backend/models"
	"github.com/csjh/NetworthMeta-Backend/types"
	"github.com/csjh/NetworthMeta-Backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"io/ioutil"
	"math"
	"net/http"
)

var bazaar types.Bazaar
var lastModified string
var oneHourBazaar = map[string][]float64{}

func ShowBazaar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(oneHourBazaar)
	if err != nil {
		return
	}
}

func ReloadBazaar() {
	resp, err := http.Get("https://api.hypixel.net/skyblock/bazaar")
	if err != nil {
		panic(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return
	}

	var temporaryBazaar types.Bazaar
	err = json.Unmarshal(body, &temporaryBazaar)
	if err != nil {
		panic(err)
		return
	}

	if temporaryBazaar.Success {
		bazaar = temporaryBazaar
		lastModified = resp.Header.Get("Last-Modified")
		if lastModified == "" {
			lastModified = "Thu, 01 Jan 1970 00:00:00 GMT"
		}
	}
}

func UpdateOneHourBazaar() {
	for item, product := range bazaar.Products {
		if len(product.BuySummary) != 0 {
			oneHourBazaar[item] = append(oneHourBazaar[item], product.BuySummary[0].PricePerUnit)
			choppingIndex := int(math.Max(float64(len(oneHourBazaar[item])-360), 0))
			oneHourBazaar[item] = oneHourBazaar[item][choppingIndex:]
		}
	}
}

func UpdateMongoBazaar() {

	client := db.Connect()
	defer db.Disconnect(client)

	collection := client.Database("networthmeta").Collection("bazaar")

	var writeModels []mongo.WriteModel
	for item, values := range oneHourBazaar {
		writeModels = append(writeModels, mongo.NewReplaceOneModel().SetFilter(bson.D{{"id", item}}).SetReplacement(models.Bazaar{ID: item, Price: utils.Median(values)}).SetUpsert(true))
	}
	_, err := collection.BulkWrite(context.TODO(), writeModels, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return
	}

}
