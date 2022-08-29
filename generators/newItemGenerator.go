package generators

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/csjh/NetworthMeta-Backend/utils"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type itemEndpointOutput struct {
	LastUpdated int64                   `json:"lastUpdated"`
	Items       []itemEndpointStructure `json:"items"`
}

type itemEndpointStructure struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Material     string `json:"material"`
	Durability   int    `json:"durability"`
	NpcSellPrice int    `json:"npc_sell_price"`
	Skin         string `json:"skin"`
}

type ItemLookup struct {
	Name         string   `json:"name"`
	PicLocation  string   `json:"pic_location"`
	Related      []string `json:"related"`
	NpcSellPrice int      `json:"npc_sell_price"`
	Recipe       []string `json:"recipe"`
}

type skin struct {
	Textures struct {
		SKIN struct {
			Url string `json:"url"`
		} `json:"SKIN"`
	} `json:"textures"`
}

var response itemEndpointOutput
var lastUpdated int64 = 0

func ParseNewItem(ID string) (ItemLookup, error) {
	if time.Now().Unix()-lastUpdated > 3600 {
		fmt.Println("fetched")

		resp, err := http.Get("https://api.hypixel.net/resources/skyblock/items")
		if err != nil {
			return ItemLookup{}, err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
			}
		}(resp.Body)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ItemLookup{}, err
		}

		err = json.Unmarshal(body, &response)
		if err != nil {
			panic(err)
			return ItemLookup{}, err
		}
		response.LastUpdated = time.Now().Unix()
	}

	lookupData := ItemLookup{
		Name:         ID,
		PicLocation:  "0-0.png",
		Related:      []string{},
		NpcSellPrice: 0,
		Recipe:       []string{},
	}

	for _, item := range response.Items {
		if item.ID != ID {
			continue
		}
		lookupData.Name = item.Name
		lookupData.NpcSellPrice = item.NpcSellPrice
		lookupData.PicLocation = item.Skin
		if item.Skin == "" {
			lookupData.PicLocation = fmt.Sprintf(`%d-%d.png`, utils.Materials[item.Material], item.Durability)
		}
		if lookupData.PicLocation[0] == 'e' {
			var skinVar skin
			decodedLocation, err := base64.StdEncoding.DecodeString(lookupData.PicLocation)
			if err != nil {
				return ItemLookup{}, err
			}
			err = json.Unmarshal(decodedLocation, &skinVar)
			if err != nil {
				return ItemLookup{}, err
			}
			lookupData.PicLocation = strings.Split(skinVar.Textures.SKIN.Url, "/texture/")[1] + ".png"
		}
	}

	return lookupData, nil
}
