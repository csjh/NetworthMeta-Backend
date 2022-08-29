package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"github.com/Tnze/go-mc/nbt"
	"github.com/csjh/NetworthMeta-Backend/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

var m1 = regexp.MustCompile("§.")

func GetRawLore(text string) string {
	return m1.ReplaceAllString(text, "")
}

func Capitalize(str string) string {
	words := strings.Replace(str, "_", " ", -1)
	return cases.Title(language.English).String(words)
}

type itemGem struct {
	GemType string
	Tier    string
}

func ParseItemGems(gems map[string]interface{}) []itemGem {
	var parsed []itemGem

	for key, value := range gems {
		slotType := strings.Split(key, "_")[0]

		if Slots["special"][slotType] {
			gemType, gemTypeIsString := gems[key+"_gem"].(string)
			gemTier, gemTierIsString := value.(string)

			if gemTypeIsString && gemTierIsString {
				parsed = append(parsed, itemGem{GemType: gemType, Tier: gemTier})
			} else {
				if opalStruct, ok := gems[key].(types.OpalGem); ok {
					parsed = append(parsed, itemGem{GemType: gemType, Tier: opalStruct.Quality})
				}
			}
		} else if Slots["normal"][slotType] {
			gemTier, gemTierIsString := value.(string)
			if gemTierIsString {
				parsed = append(parsed, itemGem{GemType: slotType, Tier: gemTier})
			} else {
				if gemTier, ok := value.(types.OpalGem); ok {
					parsed = append(parsed, itemGem{GemType: slotType, Tier: gemTier.Quality})
				}
			}
		}
	}

	return parsed
}

func DecodeNBT(data string) (types.ItemNBT, error) {
	// Decode data from base64
	buffer, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return types.ItemNBT{}, err
	}

	// Decompress data
	gr, err := gzip.NewReader(bytes.NewBuffer(buffer))
	if err != nil {
		return types.ItemNBT{}, err
	}
	defer func(gr *gzip.Reader) {
		err := gr.Close()
		if err != nil {
			panic(err)
		}
	}(gr)
	buffer, err = ioutil.ReadAll(gr)
	if err != nil {
		return types.ItemNBT{}, err
	}
	var ungzipped bytes.Buffer
	_, err = ungzipped.Write(buffer)
	if err != nil {
		return types.ItemNBT{}, err
	}

	// Parse data
	var item map[string][]types.ItemNBT
	err = nbt.Unmarshal(ungzipped.Bytes(), &item)
	if err != nil {
		return types.ItemNBT{}, err
	}

	// Return the data I want
	return item["i"][0], nil
}

func Median(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	values := make([]float64, len(arr))
	copy(values, arr)
	sort.Float64s(values)

	half := len(values) / 2
	if len(values)%2 == 1 {
		return values[half]
	}

	return (values[half-1] + values[half]) / 2.0
}
