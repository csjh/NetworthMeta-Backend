package main

import (
	"compress/zlib"
	"encoding/json"
	"fmt"
	"github.com/csjh/NetworthMeta-Backend/generators"
	"github.com/csjh/NetworthMeta-Backend/scripts"
	"github.com/csjh/NetworthMeta-Backend/types"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"math"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileHeader := r.MultipartForm.File["file"]
	if len(fileHeader) != 1 {
		http.Error(w, "Expected a single file", http.StatusBadRequest)
		return
	}

	file, err := fileHeader[0].Open()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reader, err := zlib.NewReader(file)
	result, err := ioutil.ReadAll(reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var player types.Player
	err = json.Unmarshal(result, &player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profileNetworths := make(map[string]float64)
	for profile, contents := range player {
		items := generators.GetItems(contents)
		networth := generators.GetNetworth(items).Networth
		profileNetworths[profile] = math.Round(networth)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]map[string]float64{"data": profileNetworths})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

const (
	everyFiveMinutes = "0 */5 * * * *"
	everyTenSeconds  = "*/10 * * * * *"
	everyOneMinute   = "0 */1 * * * *"
)

func main() {
	c := cron.New(cron.WithSeconds())

	scripts.ReloadBazaar()
	_, err := c.AddFunc(everyTenSeconds, scripts.ReloadBazaar)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished reload bazaar")

	scripts.UpdateOneHourBazaar()
	_, err = c.AddFunc(everyTenSeconds, scripts.UpdateOneHourBazaar)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished update one hour bazaar")

	scripts.ReloadAuctionData()
	_, err = c.AddFunc(everyOneMinute, scripts.ReloadAuctionData)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished reload auction data")

	generators.RetrievePrices()
	_, err = c.AddFunc(everyFiveMinutes, generators.RetrievePrices)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished retrieve prices")

	scripts.UpdateMongoBazaar()
	_, err = c.AddFunc(everyFiveMinutes, scripts.UpdateMongoBazaar)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished update mongo bazaar")
	c.Start()

	http.HandleFunc("/networth", handler)
	http.HandleFunc("/auctions", scripts.ShowAuctions)
	http.HandleFunc("/bazaar", scripts.ShowBazaar)
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err)
	}
}
