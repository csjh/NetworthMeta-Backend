package generators

import (
	"encoding/json"
	"fmt"
	"github.com/csjh/NetworthMeta-Backend/types"
	"github.com/csjh/NetworthMeta-Backend/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func getPricesFromDb(pet types.Pet, db map[string]float64) (lvl1 float64, lvl100 float64, lvl200 float64) {
	formattedPet := cases.Lower(language.English).String(fmt.Sprintf("%s_%s", pet.Type, pet.Tier))
	lvl1 = db[fmt.Sprintf("lvl_1_%s", formattedPet)]
	lvl100 = db[fmt.Sprintf("lvl_100_%s", formattedPet)]
	lvl200 = db[fmt.Sprintf("lvl_200_%s", formattedPet)]

	return
}

func CalculateSkillLevel(pet types.Pet) (xpMax int, level int) {
	maxLevel := 100
	if pet.Type == "GOLDEN_DRAGON" {
		maxLevel = 200
	}
	rarityOffset := utils.PetRarityOffset[pet.Tier]
	levels := utils.PetLevels[rarityOffset : int(rarityOffset)+maxLevel]

	level = 1
	totalExperience := 0

	for i := 0; i < maxLevel; i++ {
		totalExperience += levels[i]
		if float64(totalExperience) > pet.Exp {
			totalExperience -= levels[i]
			break
		}
		level++
	}
	if level > maxLevel {
		level = maxLevel
	}

	xpMax = 0
	for _, v := range levels {
		xpMax += v
	}

	return
}

func getPetPrice(pet types.ItemNBT, db map[string]float64) types.ItemNBT {
	var petInfo types.Pet
	err := json.Unmarshal([]byte(pet.Tag.ExtraAttributes.PetInfo), &petInfo)
	if err != nil {
		return pet
	}
	xpMax, level := CalculateSkillLevel(petInfo)
	lvl1, lvl100, lvl200 := getPricesFromDb(petInfo, db)

	special := ""
	if petInfo.Skin != "" {
		special = "✦"
	}
	pet.Modified = types.Modifiers{
		Id:   petInfo.Type,
		Name: fmt.Sprintf(`[Lvl %d] %s %s`, level, utils.Capitalize(fmt.Sprintf(`%s %s`, petInfo.Tier, petInfo.Type)), special),
	}

	if lvl1 == 0 || lvl100 == 0 {
		pet.Price = 0
		return pet
	}

	price := lvl100
	if lvl200 != 0 {
		price = lvl200
	}

	if level < 100 && xpMax != 0 {
		baseFormula := (lvl100 - lvl1) / float64(xpMax)
		price = baseFormula*petInfo.Exp + lvl1
	}

	if level > 100 && level < 200 {
		level %= 100
		baseFormula := (lvl200 - lvl100) / 100

		price = baseFormula*float64(level) + lvl100
	}

	if petInfo.HeldItem != "" && level != 200 {
		heldItem := db[strings.ToLower(petInfo.HeldItem)]

		if heldItem != 0 {
			price += heldItem
		}
	}

	if petInfo.CandyUsed > 0 && petInfo.Type != "ENDER_DRAGON" {
		price *= 0.65
	}

	if petInfo.Skin != "" {
		price += db[strings.ToLower(fmt.Sprintf(`PET_SKIN_%s`, petInfo.Skin))]
	}

	pet.Price = price
	return pet
}
