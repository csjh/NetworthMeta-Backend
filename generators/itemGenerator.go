package generators

import (
	"context"
	"fmt"
	"github.com/csjh/NetworthMeta-Backend/db"
	"github.com/csjh/NetworthMeta-Backend/models"
	"github.com/csjh/NetworthMeta-Backend/types"
	"github.com/csjh/NetworthMeta-Backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math"
	"regexp"
	"strings"
	"time"
)

var prices = map[string]float64{}

const (
	Silex                   float64 = 0.7
	EnchantedBooks          float64 = 0.85
	WorseEnchantedBooks     float64 = 0.35
	EvenWorseEnchantedBooks float64 = 0.2
	Runes                   float64 = 0.5
	Attributes              float64 = 0.5
	FumingPotatoBooks       float64 = 0.6
	Enrichments             float64 = 0.75
	ArtOfWar                float64 = 0.6
	FarmingForDummies       float64 = 0.5
	GemstoneChamber         float64 = 0.9
	Dyes                    float64 = 0.9
)

func RetrievePrices() {
	client := db.Connect()
	defer db.Disconnect(client)

	maroDatabase := client.Database("networthmeta")

	auctions := maroDatabase.Collection("auctions")
	bazaar := maroDatabase.Collection("bazaar")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get all auction items
	cur, err := auctions.Find(ctx, bson.D{})

	var auctionItems []models.Auction
	err = cur.All(ctx, &auctionItems)
	if err != nil {
		panic(err)
		return
	}

	for _, item := range auctionItems {
		prices[strings.ToLower(item.ID)] = item.Price
		prices["Recombobulated "+strings.ToLower(item.ID)] = item.RecombobulatedPrice
	}

	// Get all bazaar items
	cur, err = bazaar.Find(ctx, bson.D{})

	var bazaarItems []models.Bazaar
	err = cur.All(ctx, &bazaarItems)
	if err != nil {
		panic(err)
		return
	}

	for _, product := range bazaarItems {
		prices[strings.ToLower(product.ID)] = product.Price
	}

	prices[""] = 0
	prices["efficiency_10"] = prices["sil_ex"] * Silex
	prices["salty_placeholder"] = prices["blue_ice_hunk"] * 64
	prices["treacherous_placeholder"] = prices["ice_hunk"] * 64
}

func ParseItem(item types.ItemNBT) types.ItemNBT {
	ExtraAttributes := item.Tag.ExtraAttributes

	petInfo := ExtraAttributes.PetInfo
	if petInfo != "" {
		return getPetPrice(item, prices)
	}

	if ExtraAttributes.ID == "" {
		return item
	}

	itemName := utils.GetRawLore(item.Tag.Display.Name)
	itemId := strings.ToLower(ExtraAttributes.ID)

	count := item.Tag.Count
	if count == 0 {
		count = int64(item.Count)
	}

	var price float64
	if ExtraAttributes.RarityUpgrades > 0 && (ExtraAttributes.Enchantments != nil || utils.Talismans[itemId]) {
		price = prices["Recombobulated "+itemId] + prices[itemId]*float64(count-1)
	} else {
		price = prices[itemId] * float64(count)
	}

	if utils.NameLookup[itemName] != "" {
		item.Price = price
		item.Modified = types.Modifiers{Id: itemId, Name: itemName}
		return item
	}
	if strings.Contains(itemId, "midas") {
		price = math.Max(float64(ExtraAttributes.WinningBid), price)
	}

	if itemId == "enchanted_book" {
		price = 0
		itemName = "God Book"
		for enchant, level := range ExtraAttributes.Enchantments {
			price += prices[fmt.Sprintf(`%s_%d`, enchant, level)]
			if len(ExtraAttributes.Enchantments) == 1 {
				itemName = cases.Title(language.English).String(fmt.Sprintf(`%s %d`, enchant, level))
			}
		}
	} else {
		for enchant, level := range ExtraAttributes.Enchantments {
			priceMultiplier := EnchantedBooks
			if enchant == "ultimate_soul_eater" || enchant == "overload" || enchant == "big_brain" || enchant == "ultimate_fatal_tempo" || enchant == "ultimate_inferno" {
				priceMultiplier = WorseEnchantedBooks
			} else if enchant == "counter_strike" {
				priceMultiplier = EvenWorseEnchantedBooks
			}
			price += prices[fmt.Sprintf(`%s_%d`, enchant, level)] * priceMultiplier
		}
	}

	if itemId == "attribute_shard" {
		price = 0
		for attribute, level := range ExtraAttributes.Attributes {
			price = prices[fmt.Sprintf(`%s_%d`, attribute, level)]
		}
	} else {
		for attribute, level := range ExtraAttributes.Attributes {
			price += prices[fmt.Sprintf(`%s_%d`, attribute, level)] * Attributes
		}
	}

	if itemId == "rune" {
		price = 0
		for runeName, level := range ExtraAttributes.Runes {
			price = prices[fmt.Sprintf(`%s_%d`, runeName, level)]
		}
	} else {
		for runeName, level := range ExtraAttributes.Runes {
			price += prices[fmt.Sprintf(`%s_%d`, runeName, level)] * Runes
		}
	}

	/*
	   todo:
	       necromancer_souls?
	       magmaCubesKilled? kinda worthless
	       captured_player? kinda worthless and jumps around
	       raider_kills? kinda worthless and specific
	       item_tier? i think dungeon floor, worth considering, wouldn't be too hard?
	       eman_kills to FINAL_DESTINATION prefix, zombie_kills to REVENANT, REAPER prefix, spider_kills to TARANTULA prefix, RECLUSE_FANG
	       blood_god_kills? shouldn't be too hard
	       party_hat_color? again shouldn't be too hard
	       ----
	       Add sometihng in mod for cake bag, add cakes to prices
	*/

	for _, gem := range utils.ParseItemGems(ExtraAttributes.Gems) {
		price += prices[strings.ToLower(fmt.Sprintf(`%s_%s_gem`, gem.Tier, gem.GemType))]
	}

	if !utils.Talismans[itemId] {
		price += prices[utils.Reforges[ExtraAttributes.Modifier]]
	}

	if (ExtraAttributes.DungeonItemLevel != 0 && ExtraAttributes.DungeonItemLevel > 5) || (ExtraAttributes.UpgradeLevel != 0 && ExtraAttributes.UpgradeLevel > 5) {
		starsUsed := ExtraAttributes.DungeonItemLevel - 5
		if ExtraAttributes.UpgradeLevel != 0 {
			starsUsed = ExtraAttributes.UpgradeLevel - 5
		}

		for star := 0; star < int(starsUsed); star++ {
			price += prices[utils.MasterStars[star]]
		}
	}

	for _, scroll := range ExtraAttributes.AbilityScroll {
		price += prices[strings.ToLower(scroll)]
	}

	price += float64(ExtraAttributes.GemstoneSlots) * prices["gemstone_chamber"] * GemstoneChamber
	price += prices[ExtraAttributes.DrillPartEngine]
	price += prices[ExtraAttributes.DrillPartFuelTank]
	price += prices[ExtraAttributes.DrillPartUpgradeModule]

	hotPotatoes := math.Min(float64(ExtraAttributes.HotPotatoCount), 10)
	price += prices["hot_potato_book"] * hotPotatoes
	fumingPotatoes := ExtraAttributes.HotPotatoCount - int64(hotPotatoes)
	price += prices["fuming_potato_book"] * float64(fumingPotatoes) * FumingPotatoBooks

	price += prices[strings.ToLower(ExtraAttributes.DyeItem)] * Dyes
	price += prices[strings.ToLower(ExtraAttributes.Skin)]
	price += prices[strings.ToLower(fmt.Sprintf(`TALISMAN_ENRICHMENT_%s`, ExtraAttributes.TalismanEnrichment))] * Enrichments
	price += prices[strings.ToLower(ExtraAttributes.PowerAbilityScroll)]

	price += prices["etherwarp_conduit"] * float64(ExtraAttributes.Ethermerge)
	price += prices["wood_singularity"] * float64(ExtraAttributes.WoodSingularityCount)
	price += prices["the_art_of_war"] * float64(ExtraAttributes.ArtOfWarCount) * ArtOfWar
	price += prices["transmission_tuner"] * float64(ExtraAttributes.TunedTransmission)
	price += prices["farming_for_dummies"] * float64(ExtraAttributes.FarmingForDummiesCount) * FarmingForDummies
	price += prices["mana_disintegrator"] * float64(ExtraAttributes.ManaDisintegratorCount)

	item.Price = price
	item.Modified = types.Modifiers{Id: itemId, Name: itemName}
	return item
}

func parseItems(items []types.ItemNBT) []types.ItemNBT {
	parsedItems := make([]types.ItemNBT, len(items))
	for i, item := range items {
		parsedItems[i] = ParseItem(item)
	}
	return parsedItems
}

var normalStorage = regexp.MustCompile(`(inventory|Backpack|Ender Chest|Wardrobe|Accessory Bag|Auctions|Personal Vault|Museum|Bazaar|Pets)`)

func GetItems(profile types.Profile) types.Output {
	output := types.Output{}
	output.MiscStorage = make(map[string][]types.ItemNBT)

	if profile.MiscStorageCache["sacks"] != nil {
		for _, item := range profile.MiscStorageCache["sacks"] {
			itemId := utils.NameLookup[item.Tag.Display.Name]
			count := item.Tag.Count
			sackPrice := prices[strings.ToLower(itemId)]

			if sackPrice != 0 {
				output.Sacks.Total += sackPrice * float64(count)
				output.Sacks.Items = append(output.Sacks.Items, types.SackItem{
					Name:  item.Tag.Display.Name,
					Count: int(count),
					Price: sackPrice * float64(count),
				})
			}
		}
	}

	for _, chest := range profile.ChestCache {
		output.Chests = append(output.Chests, parseItems(chest)...)
	}

	for key, storage := range profile.MiscStorageCache {
		storageName := normalStorage.FindString(key)
		if storageName != "" {
			output.MiscStorage[storageName] = append(
				output.MiscStorage[storageName],
				parseItems(storage)...,
			)
		}
	}

	for _, coins := range profile.MiscCoins {
		output.Coins += coins
	}

	return output
}
