package types

type Modifiers struct {
	Id   string
	Name string
}

type ItemNBT struct {
	Count    int   `json:"Count" nbt:"Count"`
	Tag      Tag   `json:"tag" nbt:"tag"`
	Damage   int64 `json:"Damage" nbt:"Damage"`
	Price    float64
	Modified Modifiers
}

type Tag struct {
	Count           int64           `json:"Count" nbt:"Count"`
	Display         Display         `json:"display" nbt:"display"`
	ExtraAttributes ExtraAttributes `json:"ExtraAttributes" nbt:"ExtraAttributes"`
}

type Display struct {
	Lore  []string `json:"Lore" nbt:"Lore"`
	Color int64    `json:"color" nbt:"color"`
	Name  string   `json:"Name" nbt:"Name"`
}

type OpalGem struct {
	Quality string `json:"quality" nbt:"quality"`
}

type ExtraAttributes struct {
	Modifier               string                 `json:"modifier" nbt:"modifier"`
	ID                     string                 `json:"id" nbt:"id"`
	Enchantments           map[string]int64       `json:"enchantments" nbt:"enchantments"`
	UpgradeLevel           int64                  `json:"upgrade_level" nbt:"upgrade_level"`
	Attributes             map[string]int64       `json:"attributes" nbt:"attributes"`
	HotPotatoCount         int64                  `json:"hot_potato_count" nbt:"hot_potato_count"`
	Runes                  map[string]int64       `json:"runes" nbt:"runes"`
	Gems                   map[string]interface{} `json:"gems" nbt:"gems"`
	DungeonItemLevel       int64                  `json:"dungeon_item_level" nbt:"dungeon_item_level"`
	PetInfo                string                 `json:"petInfo" nbt:"petInfo"`
	RarityUpgrades         int64                  `json:"rarity_upgrades" nbt:"rarity_upgrades"`
	WoodSingularityCount   int64                  `json:"wood_singularity_count" nbt:"wood_singularity_count"`
	TunedTransmission      int64                  `json:"tuned_transmission" nbt:"tuned_transmission"`
	PowerAbilityScroll     string                 `json:"power_ability_scroll" nbt:"power_ability_scroll"`
	TalismanEnrichment     string                 `json:"talisman_enrichment" nbt:"talisman_enrichment"`
	ArtOfWarCount          int64                  `json:"art_of_war_count" nbt:"art_of_war_count"`
	Skin                   string                 `json:"skin" nbt:"skin"`
	WinningBid             int64                  `json:"winning_bid" nbt:"winning_bid"`
	DrillPartUpgradeModule string                 `json:"drill_part_upgrade_module" nbt:"drill_part_upgrade_module"`
	DrillPartEngine        string                 `json:"drill_part_engine" nbt:"drill_part_engine"`
	DrillPartFuelTank      string                 `json:"drill_part_fuel_tank" nbt:"drill_part_fuel_tank"`
	GemstoneSlots          int64                  `json:"gemstone_slots" nbt:"gemstone_slots"`
	ManaDisintegratorCount int64                  `json:"mana_disintegrator_count" nbt:"mana_disintegrator_count"`
	Ethermerge             int64                  `json:"ethermerge" nbt:"ethermerge"`
	AbilityScroll          []string               `json:"ability_scroll" nbt:"ability_scroll"`
	DyeItem                string                 `json:"dye_item" nbt:"dye_item"`
	FarmingForDummiesCount int64                  `json:"farming_for_dummies_count" nbt:"farming_for_dummies_count"`
	//EmanKills                          int64             `json:"eman_kills"`
	//BossTier                           int64             `json:"boss_tier"`
	//AnvilUses                          int64             `json:"anvil_uses"`
	//RaiderKills                        int64             `json:"raider_kills"`
	//BlocksWalked                       int64             `json:"blocks_walked"`
	//SwordKills                         int64             `json:"sword_kills"`
	//DrillFuel                          int64             `json:"drill_fuel"`
	//ItemDurability                     int64             `json:"item_durability"`
	//BuilderSWandData                   []int64           `json:"builder's_wand_data"`
	//ZombieKills                        int64             `json:"zombie_kills"`
	//BloodGodKills                      int64             `json:"blood_god_kills"`
	//PotionLevel                        int64             `json:"potion_level"`
	//Potion                             string            `json:"potion"`
	//Effects                            []Effect          `json:"effects"`
	//PotionType                         PotionType        `json:"potion_type"`
	//RepellingColor                     string            `json:"repelling_color"`
	//RepRadius                          int64             `json:"rep_radius"`
	//RepParticles                       int64             `json:"rep_particles"`
	//ExpertiseKills                     int64             `json:"expertise_kills"`
	//DungeonSkillReq                    DungeonSkillReq   `json:"dungeon_skill_req"`
	//Color                              string            `json:"color"`
	//BaseStatBoostPercentage            int64             `json:"baseStatBoostPercentage"`
	//OriginTag                          string            `json:"originTag"`
	//UUID                               string            `json:"uuid"`
	//ItemTier                           int64             `json:"item_tier"`
	//Timestamp                          string            `json:"timestamp"`
	//BossID                             string            `json:"bossId"`
	//DungeonItem                        int64             `json:"dungeon_item"`
	//SpawnedFor                         string            `json:"spawnedFor"`
	//CompactBlocks                      int64             `json:"compact_blocks"`
	//SpiderKills                        int64             `json:"spider_kills"`
	//YearObtained                       int64             `json:"yearObtained"`
	//NecromancerSouls                   []NecromancerSoul `json:"necromancer_souls"`
	//IsShiny                            int64             `json:"is_shiny"`
	//HotPotatoBonus                     string            `json:"hotPotatoBonus"`
	//NewYearsCake                       int64             `json:"new_years_cake"`
	//YogsKilled                         int64             `json:"yogsKilled"`
	//FishesCaught                       int64             `json:"fishes_caught"`
	//Ammo                               int64             `json:"ammo"`
	//HecatombSRuns                      int64             `json:"hecatomb_s_runs"`
	//Enhanced                           int64             `json:"enhanced"`
	//Extended                           int64             `json:"extended"`
	//ThunderCharge                      int64             `json:"thunder_charge"`
	//StoredDrillFuel                    int64             `json:"stored_drill_fuel"`
	//UltimateSoulEaterData              float64           `json:"ultimateSoulEaterData"`
	//PartyHatYear                       int64             `json:"party_hat_year"`
	//PartyHatColor                      PartyHatColor     `json:"party_hat_color"`
	//ChampionCombatXP                   float64           `json:"champion_combat_xp"`
	//BackpackColor                      string            `json:"backpack_color"`
	//BowKills                           int64             `json:"bow_kills"`
	//LeaderVotes                        int64             `json:"leaderVotes"`
	//LeaderPosition                     int64             `json:"leaderPosition"`
	//Exe                                int64             `json:"EXE"`
	//BlocksBroken                       int64             `json:"blocksBroken"`
	//StatsBook                          int64             `json:"stats_book"`
	//Health                             int64             `json:"health"`
	//ShopDungeonFloorCompletionRequired int64             `json:"shop_dungeon_floor_completion_required"`
	//TrainingWeightsHeldTime            int64             `json:"trainingWeightsHeldTime"`
	//GildedGiftedCoins                  int64             `json:"gilded_gifted_coins"`
	//DonatedMuseum                      int64             `json:"donated_museum"`
	//MagmaCubesKilled                   int64             `json:"magmaCubesKilled"`
	//WishingCompassUses                 int64             `json:"wishing_compass_uses"`
	//Splash                             int64             `json:"splash"`
	//GhastBlaster                       int64             `json:"ghast_blaster"`
	//RecallPotionBiome                  string            `json:"recall_potion_biome"`
	//RecallPotionLocation               string            `json:"recall_potion_location"`
	//RanchersSpeed                      int64             `json:"ranchers_speed"`
	//PotionName                         PotionName        `json:"potion_name"`
	//BlazeConsumer                      int64             `json:"blaze_consumer"`
	//DungeonPotion                      int64             `json:"dungeon_potion"`
	//CakeOwner                          string            `json:"cake_owner"`
	//SoulDurability                     int64             `json:"soul_durability"`
	//FungiCutterMode                    string            `json:"fungi_cutter_mode"`
	//WAI                                int64             `json:"WAI"`
	//Glowing                            int64             `json:"glowing"`
	//Edition                            int64             `json:"edition"`
	//TdAttuneMode                       int64             `json:"td_attune_mode"`
	//MagmaCubeAbsorber                  int64             `json:"magma_cube_absorber"`
	//FarmedCultivating                  int64             `json:"farmed_cultivating"`
	//BottleOfJyrreSeconds               int64             `json:"bottle_of_jyrre_seconds"`
	//Mixins                             []string          `json:"mixins"`
	//Radius                             int64             `json:"radius"`
	//Spray                              string            `json:"spray"`
	//TrapsDefused                       int64             `json:"trapsDefused"`
	//DungeonPaperID                     string            `json:"dungeon_paper_id"`
	//ZombiesKilled                      int64             `json:"zombiesKilled"`
	//NewYearCakeBagData                 []int64           `json:"new_year_cake_bag_data"`
	//MaxedStats                         int64             `json:"maxed_stats"`
	//TuningForkTuning                   int64             `json:"tuning_fork_tuning"`
	//SmallBackpackData                  []int64           `json:"small_backpack_data"`
}
