package generators

import (
	"github.com/csjh/NetworthMeta-Backend/types"
	"math"
)

func isItemRecombobulated(item types.ItemNBT) bool {
	return item.Tag.ExtraAttributes.RarityUpgrades != 0
}

func remove(s []types.CalculatedItem, i int) []types.CalculatedItem {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func GetNetworth(data types.Output) types.Networth {
	output := types.Networth{}
	output.Categories = make(map[string]types.Category)

	for key, category := range data.MiscStorage {
		categoryValue := types.Category{}

		for _, item := range category {
			if item.Price == 0 || math.IsNaN(item.Price) {
				continue
			}

			categoryValue.Total += item.Price

			categoryValue.Items = append(categoryValue.Items, types.CalculatedItem{
				Id:     item.Modified.Id,
				Name:   item.Modified.Name,
				Price:  item.Price,
				Recomb: isItemRecombobulated(item),
				Count:  item.Count,
			})
		}

		if len(categoryValue.Items) > 0 {
			output.Categories[key] = categoryValue
		}
	}

	chestCategory := types.Category{}

	for _, chest := range data.Chests {
		if chest.Price == 0 || math.IsNaN(chest.Price) {
			continue
		}

		chestCategory.Total += chest.Price

		chestCategory.Items = append(chestCategory.Items, types.CalculatedItem{
			Id:     chest.Modified.Id,
			Name:   chest.Modified.Name,
			Price:  chest.Price,
			Recomb: isItemRecombobulated(chest),
			Count:  chest.Count,
		})
	}

	if len(chestCategory.Items) > 0 {
		for i, potentialItem := range chestCategory.Items {
			if potentialItem.Price == 0 || math.IsNaN(potentialItem.Price) {
				chestCategory.Items = remove(chestCategory.Items, i)
			}
		}
		output.Categories["chests"] = chestCategory
	}

	output.Sacks.Total = data.Sacks.Total
	for _, potentialItem := range data.Sacks.Items {
		if potentialItem.Price == 0 || math.IsNaN(potentialItem.Price) {
			continue
		}
		output.Sacks.Items = append(output.Sacks.Items, potentialItem)
	}

	for _, value := range output.Categories {
		output.Networth += value.Total
	}
	output.Networth += output.Sacks.Total + data.Coins

	return output
}
