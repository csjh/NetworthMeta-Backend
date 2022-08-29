package utils

var MasterStars = []string{"first_master_star", "second_master_star", "third_master_star", "fourth_master_star", "fifth_master_star"}
var Talismans = map[string]bool{"wedding_ring_0": true, "wedding_ring_2": true, "wedding_ring_4": true, "wedding_ring_7": true, "wedding_ring_9": true, "campfire_talisman_1": true, "campfire_talisman_4": true, "campfire_talisman_8": true, "campfire_talisman_13": true, "campfire_talisman_21": true, "farming_talisman": false, "vaccine_talisman": true, "wood_talisman": true, "skeleton_talisman": false, "coin_talisman": true, "magnetic_talisman": false, "gravity_talisman": false, "village_talisman": false, "mine_talisman": false, "night_vision_charm": false, "lava_talisman": false, "scavenger_talisman": false, "fire_talisman": false, "piggy_bank": false, "cracked_piggy_bank": false, "broken_piggy_bank": false, "pigs_foot": false, "wolf_paw": false, "frozen_chicken": false, "fish_affinity_talisman": false, "farmer_orb": false, "haste_ring": false, "experience_artifact": false, "new_year_cake_bag": false, "day_crystal": true, "night_crystal": true, "feather_talisman": false, "feather_ring": false, "feather_artifact": false, "potion_affinity_talisman": false, "ring_potion_affinity": false, "artifact_potion_affinity": false, "healing_talisman": false, "healing_ring": false, "candy_talisman": false, "candy_ring": false, "candy_artifact": false, "melody_hair": true, "sea_creature_talisman": true, "sea_creature_ring": false, "sea_creature_artifact": false, "intimidation_talisman": false, "intimidation_ring": false, "intimidation_artifact": false, "wolf_talisman": false, "wolf_ring": false, "bat_talisman": false, "bat_ring": false, "bat_artifact": false, "devour_ring": false, "zombie_talisman": false, "zombie_ring": false, "zombie_artifact": false, "spider_talisman": false, "spider_ring": false, "spider_artifact": false, "ender_artifact": false, "tarantula_talisman": false, "survivor_cube": false, "wither_artifact": false, "red_claw_talisman": false, "red_claw_ring": false, "red_claw_artifact": false, "bait_ring": false, "shady_ring": false, "crooked_artifact": false, "seal_of_the_family": false, "hunter_talisman": false, "hunter_ring": false, "party_hat_crab": false, "potato_talisman": false, "emerald_ring": true, "personal_compactor_4000": true, "personal_compactor_5000": true, "personal_compactor_6000": true, "personal_compactor_7000": true, "personal_deletor_4000": true, "personal_deletor_5000": true, "personal_deletor_6000": true, "personal_deletor_7000": true, "speed_talisman": true, "speed_ring": true, "speed_artifact": true, "cat_talisman": true, "lynx_talisman": true, "cheetah_talisman": true, "scarf_studies": true, "scarf_thesis": true, "scarf_grimoire": true, "treasure_talisman": true, "treasure_ring": true, "treasure_artifact": true, "mineral_talisman": true, "beastmaster_crest_common": true, "beastmaster_crest_uncommon": true, "beastmaster_crest_rare": true, "beastmaster_crest_epic": true, "beastmaster_crest_legendary": true, "raggedy_shark_tooth_necklace": true, "dull_shark_tooth_necklace": true, "honed_shark_tooth_necklace": true, "sharp_shark_tooth_necklace": true, "razor_sharp_shark_tooth_necklace": true, "hegemony_artifact": true, "bits_talisman": true, "bat_person_talisman": true, "bat_person_ring": true, "bat_person_artifact": true, "candy_relic": true, "lucky_hoof": false, "eternal_hoof": false, "wither_relic": true, "catacombs_expert_ring": true, "auto_recombobulator": true, "jerry_talisman_green": true, "jerry_talisman_blue": true, "jerry_talisman_purple": true, "jerry_talisman_golden": true, "king_talisman": true, "titanium_talisman": true, "titanium_ring": true, "titanium_artifact": true, "titanium_relic": true, "reaper_orb": true, "dante_talisman": true, "spiked_atrocity": true, "blood_god_crest": true, "master_skull_tier_1": true, "master_skull_tier_2": true, "master_skull_tier_3": true, "master_skull_tier_4": true, "master_skull_tier_5": true, "master_skull_tier_6": true, "master_skull_tier_7": true, "soulflow_pile": true, "soulflow_battery": true, "soulflow_supercell": true, "pocket_espresso_machine": true, "handy_blood_chalice": true, "ender_relic": true, "jungle_amulet": true, "power_talisman": true, "power_ring": true, "power_artifact": true, "nether_artifact": true, "bingo_talisman": true, "bingo_ring": true}
var Reforges = map[string]string{"salty": "salty_placeholder", "treacherous": "treacherous_placeholder", "moil": "moil_log", "dirty": "dirt_bottle", "stiff": "hardened_wood", "aote_stone": "aote_stone", "waxed": "blaze_wax", "candied": "candy_corn", "perfect": "diamond_atom", "fleet": "diamonite", "fabled": "dragon_claw", "spiked": "dragon_scale", "warped": "endstone_geode", "coldfusion": "entropy_suppressor", "jaded": "jaderald", "jerry_stone": "jerry_stone", "magnetic": "lapis_crystal", "fortified": "meteor_shard", "gilded": "midas_jewel", "cubic": "molten_cube", "necrotic": "necromancer_brooch", "fruitful": "onyx", "precise": "optical_lens", "undead": "premium_flesh", "mithraic": "pure_mithril", "reinforced": "rare_diamond", "ridiculous": "red_nose", "loving": "red_scarf", "auspicious": "rock_gemstone", "headstrong": "salmon_opal", "strengthened": "searing_stone", "glistening": "shiny_prism", "spiritual": "spirit_decoy", "suspicious": "suspicious_vial", "ambered": "amber_material", "blessed": "blessed_fruit", "bulky": "bulky_stone", "submerged": "deep_sea_orb", "renowned": "dragon_horn", "giant": "giant_tooth", "bountiful": "golden_ball", "chomp": "kuudra_mandible", "lucky": "lucky_dice", "stellar": "petrified_starfall", "ancient": "precursor_gear", "refined": "refined_amber", "empowered": "sadan_brooch", "withered": "wither_blood", "heated": "hot_stuff", "silky": "luxurious_spool", "sweet": "rock_candy", "forceful": "acacia_birdhouse", "bloody": "beating_heart", "shaded": "dark_orb", "toil": "toil_log"}
var Slots = map[string]map[string]bool{
	"normal": {
		"JADE": true, "AMBER": true, "TOPAZ": true, "SAPPHIRE": true, "AMETHYST": true, "JASPER": true, "RUBY": true, "OPAL": true,
	},
	"special": {
		"UNIVERSAL": true, "COMBAT": true, "OFFENSIVE": true, "DEFENSIVE": true, "MINING": true,
	},
}
var PetLevels = []int{100, 110, 120, 130, 145, 160, 175, 190, 210, 230, 250, 275, 300, 330, 360, 400, 440, 490, 540, 600, 660, 730, 800, 880, 960, 1050, 1150, 1260, 1380, 1510, 1650, 1800, 1960, 2130, 2310, 2500, 2700, 2920, 3160, 3420, 3700, 4000, 4350, 4750, 5200, 5700, 6300, 7000, 7800, 8700, 9700, 10800, 12000, 13300, 14700, 16200, 17800, 19500, 21300, 23200, 25200, 27400, 29800, 32400, 35200, 38200, 41400, 44800, 48400, 52200, 56200, 60400, 64800, 69400, 74200, 79200, 84700, 90700, 97200, 104200, 111700, 119700, 128200, 137200, 146700, 156700, 167700, 179700, 192700, 206700, 221700, 237700, 254700, 272700, 291700, 311700, 333700, 357700, 383700, 411700, 441700, 476700, 516700, 561700, 611700, 666700, 726700, 791700, 861700, 936700, 1016700, 1101700, 1191700, 1286700, 1386700, 1496700, 1616700, 1746700, 1886700, 0, 1, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700, 1886700}
var PetRarityOffset = map[string]int8{"COMMON": 0, "UNCOMMON": 6, "RARE": 11, "EPIC": 16, "LEGENDARY": 20}
var Materials = map[string]int{"FIREWORK": 401, "SLIME_BLOCK": 165, "EYE_OF_ENDER": 381, "TRAP_DOOR": 96, "WOOD_BUTTON": 143, "REDSTONE_LAMP_OFF": 123, "SNOW_BALL": 332, "MYCEL": 110, "RAW_BEEF": 363, "SUGAR_CANE": 338, "WATER_LILY": 111, "RAW_CHICKEN": 365, "NETHER_STALK": 372, "GRILLED_PORK": 320, "SEEDS": 295, "SULPHUR": 289, "EXP_BOTTLE": 384, "ENDER_STONE": 121, "CARROT_STICK": 398, "SNOW_BLOCK": 80, "LOG_2": 162, "PORK": 319, "HUGE_MUSHROOM_1": 99, "HUGE_MUSHROOM_2": 100, "CARROT_ITEM": 391, "POTATO_ITEM": 392, "RAW_FISH": 349, "SKULL_ITEM": 397, "INK_SACK": 351, "ACACIA_DOOR": 430, "ACACIA_FENCE": 192, "ACACIA_FENCE_GATE": 187, "ACACIA_STAIRS": 163, "ACTIVATOR_RAIL": 157, "AIR": 0, "ANVIL": 145, "APPLE": 260, "ARMOR_STAND": 416, "ARROW": 262, "BAKED_POTATO": 393, "BANNER": 425, "BARRIER": 166, "BEACON": 138, "BED": 355, "BEDROCK": 7, "BEEF": 363, "BIRCH_DOOR": 428, "BIRCH_FENCE": 189, "BIRCH_FENCE_GATE": 184, "BIRCH_STAIRS": 135, "BLAZE_POWDER": 377, "BLAZE_ROD": 369, "BOAT": 333, "BONE": 352, "BOOK": 340, "BOOKSHELF": 47, "BOW": 261, "BOWL": 281, "BREAD": 297, "BREWING_STAND": 379, "BRICK": 336, "BRICK_BLOCK": 45, "BRICK_STAIRS": 108, "BROWN_MUSHROOM": 39, "BROWN_MUSHROOM_BLOCK": 99, "BUCKET": 325, "CACTUS": 81, "CAKE": 354, "CARPET": 171, "CARROT": 391, "CARROT_ON_A_STICK": 398, "CAULDRON": 380, "CHAINMAIL_BOOTS": 305, "CHAINMAIL_CHESTPLATE": 303, "CHAINMAIL_HELMET": 302, "CHAINMAIL_LEGGINGS": 304, "CHEST": 54, "CHEST_MINECART": 342, "CHICKEN": 365, "CLAY": 82, "CLAY_BALL": 337, "CLOCK": 347, "COAL": 263, "COAL_BLOCK": 173, "COAL_ORE": 16, "COBBLESTONE": 4, "COBBLESTONE_WALL": 139, "COMMAND_BLOCK": 137, "COMMAND_BLOCK_MINECART": 422, "COMPARATOR": 404, "COMPASS": 345, "COOKED_BEEF": 364, "COOKED_CHICKEN": 366, "COOKED_FISH": 350, "COOKED_MUTTON": 424, "COOKED_PORKCHOP": 320, "COOKED_RABBIT": 412, "COOKIE": 357, "CRAFTING_TABLE": 58, "DARK_OAK_DOOR": 431, "DARK_OAK_FENCE": 191, "DARK_OAK_FENCE_GATE": 186, "DARK_OAK_STAIRS": 164, "DAYLIGHT_DETECTOR": 151, "DEADBUSH": 32, "DETECTOR_RAIL": 28, "DIAMOND": 264, "DIAMOND_AXE": 279, "DIAMOND_BLOCK": 57, "DIAMOND_BOOTS": 313, "DIAMOND_CHESTPLATE": 311, "DIAMOND_HELMET": 310, "DIAMOND_HOE": 293, "DIAMOND_HORSE_ARMOR": 419, "DIAMOND_LEGGINGS": 312, "DIAMOND_ORE": 56, "DIAMOND_PICKAXE": 278, "DIAMOND_SHOVEL": 277, "DIAMOND_SWORD": 276, "DIRT": 3, "DISPENSER": 23, "DOUBLE_PLANT": 175, "DRAGON_EGG": 122, "DROPPER": 158, "DYE": 351, "EGG": 344, "EMERALD": 388, "EMERALD_BLOCK": 133, "EMERALD_ORE": 129, "ENCHANTED_BOOK": 403, "ENCHANTING_TABLE": 116, "END_PORTAL_FRAME": 120, "END_STONE": 121, "ENDER_CHEST": 130, "ENDER_EYE": 381, "ENDER_PEARL": 368, "EXPERIENCE_BOTTLE": 384, "FARMLAND": 60, "FEATHER": 288, "FENCE": 85, "FENCE_GATE": 107, "FERMENTED_SPIDER_EYE": 376, "FILLED_MAP": 358, "FIRE_CHARGE": 385, "FIREWORK_CHARGE": 402, "FIREWORKS": 401, "FISH": 349, "FISHING_ROD": 346, "FLINT": 318, "FLINT_AND_STEEL": 259, "FLOWER_POT": 390, "FURNACE": 61, "FURNACE_MINECART": 343, "GHAST_TEAR": 370, "GLASS": 20, "GLASS_BOTTLE": 374, "GLASS_PANE": 102, "GLOWSTONE": 89, "GLOWSTONE_DUST": 348, "GOLD_BLOCK": 41, "GOLD_INGOT": 266, "GOLD_NUGGET": 371, "GOLD_ORE": 14, "GOLDEN_APPLE": 322, "GOLDEN_AXE": 286, "GOLDEN_BOOTS": 317, "GOLDEN_CARROT": 396, "GOLDEN_CHESTPLATE": 315, "GOLDEN_HELMET": 314, "GOLDEN_HOE": 294, "GOLDEN_HORSE_ARMOR": 418, "GOLDEN_LEGGINGS": 316, "GOLDEN_PICKAXE": 285, "GOLDEN_RAIL": 27, "GOLDEN_SHOVEL": 284, "GOLDEN_SWORD": 283, "GRASS": 2, "GRAVEL": 13, "GUNPOWDER": 289, "HARDENED_CLAY": 172, "HAY_BLOCK": 170, "HEAVY_WEIGHTED_PRESSURE_PLATE": 148, "HOPPER": 154, "HOPPER_MINECART": 408, "ICE": 79, "IRON_AXE": 258, "IRON_BARS": 101, "IRON_BLOCK": 42, "IRON_BOOTS": 309, "IRON_CHESTPLATE": 307, "IRON_DOOR": 330, "IRON_HELMET": 306, "IRON_HOE": 292, "IRON_HORSE_ARMOR": 417, "IRON_INGOT": 265, "IRON_LEGGINGS": 308, "IRON_ORE": 15, "IRON_PICKAXE": 257, "IRON_SHOVEL": 256, "IRON_SWORD": 267, "IRON_TRAPDOOR": 167, "ITEM_FRAME": 389, "JUKEBOX": 84, "JUNGLE_DOOR": 429, "JUNGLE_FENCE": 190, "JUNGLE_FENCE_GATE": 185, "JUNGLE_STAIRS": 136, "LADDER": 65, "LAPIS_BLOCK": 22, "LAPIS_ORE": 21, "LAVA_BUCKET": 327, "LEAD": 420, "LEATHER": 334, "LEATHER_BOOTS": 301, "LEATHER_CHESTPLATE": 299, "LEATHER_HELMET": 298, "LEATHER_LEGGINGS": 300, "LEAVES": 18, "LEAVES2": 161, "LEVER": 69, "LIGHT_WEIGHTED_PRESSURE_PLATE": 147, "LIT_PUMPKIN": 91, "LOG": 17, "LOG2": 162, "MAGMA_CREAM": 378, "MAP": 395, "MELON": 360, "MELON_BLOCK": 103, "MELON_SEEDS": 362, "MILK_BUCKET": 335, "MINECART": 328, "MOB_SPAWNER": 52, "MONSTER_EGG": 383, "MOSSY_COBBLESTONE": 48, "MUSHROOM_STEW": 282, "MUTTON": 423, "MYCELIUM": 110, "NAME_TAG": 421, "NETHER_BRICK": 112, "NETHER_BRICK_FENCE": 113, "NETHER_BRICK_STAIRS": 114, "NETHER_STAR": 399, "NETHER_WART": 372, "NETHERBRICK": 405, "NETHERRACK": 87, "NOTEBLOCK": 25, "OAK_STAIRS": 53, "OBSIDIAN": 49, "PACKED_ICE": 174, "PAINTING": 321, "PAPER": 339, "PISTON": 33, "PLANKS": 5, "POISONOUS_POTATO": 394, "PORKCHOP": 319, "POTATO": 392, "POTION": 373, "PRISMARINE": 168, "PRISMARINE_CRYSTALS": 410, "PRISMARINE_SHARD": 409, "PUMPKIN": 86, "PUMPKIN_PIE": 400, "PUMPKIN_SEEDS": 361, "QUARTZ": 406, "QUARTZ_BLOCK": 155, "QUARTZ_ORE": 153, "QUARTZ_STAIRS": 156, "RABBIT": 411, "RABBIT_FOOT": 414, "RABBIT_HIDE": 415, "RABBIT_STEW": 413, "RAIL": 66, "RECORD_11": 2266, "RECORD_13": 2256, "RECORD_BLOCKS": 2258, "RECORD_CAT": 2257, "RECORD_CHIRP": 2259, "RECORD_FAR": 2260, "RECORD_MALL": 2261, "RECORD_MELLOHI": 2262, "RECORD_STAL": 2263, "RECORD_STRAD": 2264, "RECORD_WAIT": 2267, "RECORD_WARD": 2265, "RED_FLOWER": 38, "RED_MUSHROOM": 40, "RED_MUSHROOM_BLOCK": 100, "RED_SANDSTONE": 179, "RED_SANDSTONE_STAIRS": 180, "REDSTONE": 331, "REDSTONE_BLOCK": 152, "REDSTONE_LAMP": 123, "REDSTONE_ORE": 73, "REDSTONE_TORCH": 76, "REEDS": 338, "REPEATER": 356, "ROTTEN_FLESH": 367, "SADDLE": 329, "SAND": 12, "SANDSTONE": 24, "SANDSTONE_STAIRS": 128, "SAPLING": 6, "SEA_LANTERN": 169, "SHEARS": 359, "SIGN": 323, "SKULL": 397, "SLIME": 165, "SLIME_BALL": 341, "SNOW": 80, "SNOW_LAYER": 78, "SNOWBALL": 332, "SOUL_SAND": 88, "SPAWN_EGG": 383, "SPECKLED_MELON": 382, "SPIDER_EYE": 375, "SPONGE": 19, "SPRUCE_DOOR": 427, "SPRUCE_FENCE": 188, "SPRUCE_FENCE_GATE": 183, "SPRUCE_STAIRS": 134, "STAINED_GLASS": 95, "STAINED_GLASS_PANE": 160, "STAINED_HARDENED_CLAY": 159, "STICK": 280, "STICKY_PISTON": 29, "STONE": 1, "STONE_AXE": 275, "STONE_BRICK_STAIRS": 109, "STONE_BUTTON": 77, "STONE_HOE": 291, "STONE_PICKAXE": 274, "STONE_PRESSURE_PLATE": 70, "STONE_SHOVEL": 273, "STONE_SLAB": 44, "STONE_SLAB2": 182, "STONE_STAIRS": 67, "STONE_SWORD": 272, "STONEBRICK": 98, "STRING": 287, "SUGAR": 353, "TALLGRASS": 31, "TNT": 46, "TNT_MINECART": 407, "TORCH": 50, "TRAPDOOR": 96, "TRAPPED_CHEST": 146, "TRIPWIRE_HOOK": 131, "VINE": 106, "WATER_BUCKET": 326, "WATERLILY": 111, "WEB": 30, "WHEAT": 296, "WHEAT_SEEDS": 295, "WOODEN_AXE": 271, "WOODEN_BUTTON": 143, "WOODEN_DOOR": 324, "WOODEN_HOE": 290, "WOODEN_PICKAXE": 270, "WOODEN_PRESSURE_PLATE": 72, "WOODEN_SHOVEL": 269, "WOODEN_SLAB": 126, "WOODEN_SWORD": 268, "WOOL": 35, "WRITABLE_BOOK": 386, "WRITTEN_BOOK": 387, "YELLOW_FLOWER": 3}
var NameLookup = map[string]string{
	"Bazaar Cookie":                  "BAZAAR_COOKIE",
	"Enchanted Carrot On A Stick":    "ENCHANTED_CARROT_ON_A_STICK",
	"Corrupted Bait":                 "CORRUPTED_BAIT",
	"Hot Bait":                       "HOT_BAIT",
	"Magmafish":                      "MAGMA_FISH",
	"Silver Magmafish":               "MAGMA_FISH_SILVER",
	"Gold Magmafish":                 "MAGMA_FISH_GOLD",
	"Diamond Magmafish":              "MAGMA_FISH_DIAMOND",
	"Amalgamated Crimsonite":         "AMALGAMATED_CRIMSONITE",
	"Old Amalgamated Crimsonite":     "AMALGAMATED_CRIMSONITE_NEW",
	"Melon Block":                    "MELON_BLOCK",
	"Cocoa Beans":                    "INK_SACK:3",
	"Lapis Lazuli":                   "INK_SACK:4",
	"Tarantula Web":                  "TARANTULA_WEB",
	"Horns of Torment":               "HORNS_OF_TORMENT",
	"Green Jerry Box":                "JERRY_BOX_GREEN",
	"Dark Orb":                       "DARK_ORB",
	"Enchanted Golden Carrot":        "ENCHANTED_GOLDEN_CARROT",
	"Enchanted Melon":                "ENCHANTED_MELON",
	"Enchanted Mycelium Cube":        "ENCHANTED_MYCELIUM_CUBE",
	"Enchanted Blaze Rod":            "ENCHANTED_BLAZE_ROD",
	"Flawed Jasper Gemstone":         "FLAWED_JASPER_GEM",
	"Enchanted Brown Mushroom":       "ENCHANTED_BROWN_MUSHROOM",
	"Yellow Goblin Egg":              "GOBLIN_EGG_YELLOW",
	"Flawed Amber Gemstone":          "FLAWED_AMBER_GEM",
	"Enchanted Glistering Melon":     "ENCHANTED_GLISTERING_MELON",
	"Prismarine Shard":               "PRISMARINE_SHARD",
	"Protector Dragon Fragment":      "PROTECTOR_FRAGMENT",
	"Enchanted Melon Block":          "ENCHANTED_MELON_BLOCK",
	"Rock Gemstone":                  "ROCK_GEMSTONE",
	"Power Crystal":                  "POWER_CRYSTAL",
	"Diamond":                        "DIAMOND",
	"Shark Fin":                      "SHARK_FIN",
	"Wise Dragon Fragment":           "WISE_FRAGMENT",
	"Cobblestone":                    "COBBLESTONE",
	"Refined Mithril":                "REFINED_MITHRIL",
	"Perfect Ruby Gemstone":          "PERFECT_RUBY_GEM",
	"Enchanted Pufferfish":           "ENCHANTED_PUFFERFISH",
	"End Stone Shulker":              "END_STONE_SHULKER",
	"Potato":                         "POTATO_ITEM",
	"Enchanted Hard Stone":           "ENCHANTED_HARD_STONE",
	"Enchanted Brown Mushroom Block": "ENCHANTED_HUGE_MUSHROOM_1",
	"Ender Monocle":                  "ENDER_MONOCLE",
	"Refined Diamond":                "REFINED_DIAMOND",
	"Enchanted Cobblestone":          "ENCHANTED_COBBLESTONE",
	"Hazmat Enderman":                "HAZMAT_ENDERMAN",
	"Luxurious Spool":                "LUXURIOUS_SPOOL",
	"Flawed Opal Gemstone":           "FLAWED_OPAL_GEM",
	"Enchanted Red Mushroom Block":   "ENCHANTED_HUGE_MUSHROOM_2",
	"Raw Porkchop":                   "PORK",
	"Ice Bait":                       "ICE_BAIT",
	"Dark Oak Wood":                  "LOG_2:1",
	"Enchanted Snow Block":           "ENCHANTED_SNOW_BLOCK",
	"Tasty Cheese":                   "CHEESE_FUEL",
	"Kuudra Mandible":                "KUUDRA_MANDIBLE",
	"Golden Jerry Box":               "JERRY_BOX_GOLDEN",
	"Pumpkin Guts":                   "PUMPKIN_GUTS",
	"Blaze Rod Distillate":           "BLAZE_ROD_DISTILLATE",
	"Acacia Birdhouse":               "ACACIA_BIRDHOUSE",
	"Enchanted Carrot on a Stick":    "ENCHANTED_CARROT_STICK",
	"Enchanted End Stone":            "ENCHANTED_ENDSTONE",
	"Enchanted Lapis Block":          "ENCHANTED_LAPIS_LAZULI_BLOCK",
	"Enchanted Sand":                 "ENCHANTED_SAND",
	"Colossal Experience Bottle":     "COLOSSAL_EXP_BOTTLE",
	"Enchanted String":               "ENCHANTED_STRING",
	"Derelict Ashe":                  "DERELICT_ASHE",
	"Moil Log":                       "MOIL_LOG",
	"Enchanted Acacia Wood":          "ENCHANTED_ACACIA_LOG",
	"Sand":                           "SAND",
	"Soul Fragment":                  "SOUL_FRAGMENT",
	"Plasma Bucket":                  "PLASMA_BUCKET",
	"Ancient Claw":                   "ANCIENT_CLAW",
	"Enchanted Lapis Lazuli":         "ENCHANTED_LAPIS_LAZULI",
	"Enchanted Ghast Tear":           "ENCHANTED_GHAST_TEAR",
	"Enchanted Cocoa Bean":           "ENCHANTED_COCOA",
	"Melon":                          "MELON",
	"Amber Material":                 "AMBER_MATERIAL",
	"Carrot Bait":                    "CARROT_BAIT",
	"Fine Topaz Gemstone":            "FINE_TOPAZ_GEM",
	"Gabagool Distillate":            "CRUDE_GABAGOOL_DISTILLATE",
	"Seeds":                          "SEEDS",
	"Fine Ruby Gemstone":             "FINE_RUBY_GEM",
	"Enchanted Leather":              "ENCHANTED_LEATHER",
	"Enchanted Sponge":               "ENCHANTED_SPONGE",
	"Hay Bale":                       "HAY_BLOCK",
	"Wolf Tooth":                     "WOLF_TOOTH",
	"Enchanted Rotten Flesh":         "ENCHANTED_ROTTEN_FLESH",
	"Enchanted Grilled Pork":         "ENCHANTED_GRILLED_PORK",
	"Enchanted Nether Wart":          "ENCHANTED_NETHER_STALK",
	"Enchanted Ancient Claw":         "ENCHANTED_ANCIENT_CLAW",
	"Enchanted Redstone":             "ENCHANTED_REDSTONE",
	"Meteor Shard":                   "METEOR_SHARD",
	"Dwarven Super Compactor":        "DWARVEN_COMPACTOR",
	"Great White Shark Tooth":        "GREAT_WHITE_SHARK_TOOTH",
	"Enchanted Lava Bucket":          "ENCHANTED_LAVA_BUCKET",
	"Pufferfish":                     "RAW_FISH:3",
	"Recombobulator 3000":            "RECOMBOBULATOR_3000",
	"Diamond Atom":                   "DIAMOND_ATOM",
	"Clownfish":                      "RAW_FISH:2",
	"Raw Salmon":                     "RAW_FISH:1",
	"Precursor Gear":                 "PRECURSOR_GEAR",
	"Enchanted Pork":                 "ENCHANTED_PORK",
	"Necromancer's Brooch":           "NECROMANCER_BROOCH",
	"Red Goblin Egg":                 "GOBLIN_EGG_RED",
	"Deep Sea Orb":                   "DEEP_SEA_ORB",
	"Rough Jasper Gemstone":          "ROUGH_JASPER_GEM",
	"Feather":                        "FEATHER",
	"Whale Bait":                     "WHALE_BAIT",
	"Petrified Starfall":             "PETRIFIED_STARFALL",
	"Sponge":                         "SPONGE",
	"Enchanted Dark Oak Wood":        "ENCHANTED_DARK_OAK_LOG",
	"Flawless Topaz Gemstone":        "FLAWLESS_TOPAZ_GEM",
	"Enchanted Red Sand":             "ENCHANTED_RED_SAND",
	"Enchanted Clownfish":            "ENCHANTED_CLOWNFISH",
	"Enchanted Gold":                 "ENCHANTED_GOLD",
	"Enchanted Raw Chicken":          "ENCHANTED_RAW_CHICKEN",
	"Enchanted Lily Pad":             "ENCHANTED_WATER_LILY",
	"Red Scarf":                      "RED_SCARF",
	"Spruce Wood":                    "LOG:1",
	"Titanium":                       "TITANIUM_ORE",
	"Blue Shark Tooth":               "BLUE_SHARK_TOOTH",
	"Jungle Wood":                    "LOG:3",
	"Birch Wood":                     "LOG:2",
	"Sadan's Brooch":                 "SADAN_BROOCH",
	"Enchanted Glowstone Dust":       "ENCHANTED_GLOWSTONE_DUST",
	"Enchanted Cactus":               "ENCHANTED_CACTUS",
	"Enchanted Sugar Cane":           "ENCHANTED_SUGAR_CANE",
	"Onyx":                           "ONYX",
	"Enchanted Cooked Salmon":        "ENCHANTED_COOKED_SALMON",
	"Enchanted Sulphur":              "ENCHANTED_SULPHUR",
	"Enchanted Seeds":                "ENCHANTED_SEEDS",
	"Concentrated Stone":             "CONCENTRATED_STONE",
	"Enchanted Bone Block":           "ENCHANTED_BONE_BLOCK",
	"Ghast Tear":                     "GHAST_TEAR",
	"Omega Enchanted Egg":            "OMEGA_EGG",
	"Unstable Dragon Fragment":       "UNSTABLE_FRAGMENT",
	"Purple Candy":                   "PURPLE_CANDY",
	"Polished Pumpkin":               "POLISHED_PUMPKIN",
	"Magma Cream Distillate":         "MAGMA_CREAM_DISTILLATE",
	"Enchanted Gold Block":           "ENCHANTED_GOLD_BLOCK",
	"Fine Opal Gemstone":             "FINE_OPAL_GEM",
	"Enchanted Flint":                "ENCHANTED_FLINT",
	"Enchanted Emerald Block":        "ENCHANTED_EMERALD_BLOCK",
	"Furball":                        "FURBALL",
	"Enchanted Clay":                 "ENCHANTED_CLAY_BALL",
	"Giant Tooth":                    "GIANT_TOOTH",
	"Glowstone Dust":                 "GLOWSTONE_DUST",
	"Perfect Amethyst Gemstone":      "PERFECT_AMETHYST_GEM",
	"Enchanted Mutton":               "ENCHANTED_MUTTON",
	"End Stone Geode":                "ENDSTONE_GEODE",
	"Null Sphere":                    "NULL_SPHERE",
	"Stock of Stonks":                "STOCK_OF_STONKS",
	"Enchanted Bone":                 "ENCHANTED_BONE",
	"Optical Lens":                   "OPTICAL_LENS",
	"Enchanted Diamond Block":        "ENCHANTED_DIAMOND_BLOCK",
	"Superior Dragon Fragment":       "SUPERIOR_FRAGMENT",
	"Magma Bucket":                   "MAGMA_BUCKET",
	"Green Goblin Egg":               "GOBLIN_EGG_GREEN",
	"Light Bait":                     "LIGHT_BAIT",
	"Hot Potato Book":                "HOT_POTATO_BOOK",
	"Clay":                           "CLAY_BALL",
	"Old Dragon Fragment":            "OLD_FRAGMENT",
	"Worm Membrane":                  "WORM_MEMBRANE",
	"Beating Heart":                  "BEATING_HEART",
	"Rough Amber Gemstone":           "ROUGH_AMBER_GEM",
	"Lily Pad":                       "WATER_LILY",
	"Acacia Wood":                    "LOG_2",
	"Fine Amber Gemstone":            "FINE_AMBER_GEM",
	"Enchanted Quartz":               "ENCHANTED_QUARTZ",
	"Coal":                           "COAL",
	"Ender Pearl":                    "ENDER_PEARL",
	"Enchanted Block of Coal":        "ENCHANTED_COAL_BLOCK",
	"Perfect Topaz Gemstone":         "PERFECT_TOPAZ_GEM",
	"Goblin Egg":                     "GOBLIN_EGG",
	"Enchanted Prismarine Crystals":  "ENCHANTED_PRISMARINE_CRYSTALS",
	"Enchanted Wet Sponge":           "ENCHANTED_WET_SPONGE",
	"Flawed Jade Gemstone":           "FLAWED_JADE_GEM",
	"End Stone":                      "ENDER_STONE",
	"Nether Quartz":                  "QUARTZ",
	"Purple Jerry Box":               "JERRY_BOX_PURPLE",
	"Toil Log":                       "TOIL_LOG",
	"Salmon Opal":                    "SALMON_OPAL",
	"Sludge Juice":                   "SLUDGE_JUICE",
	"Magma Cream":                    "MAGMA_CREAM",
	"Sugar Cane":                     "SUGAR_CANE",
	"Enchanted Mithril":              "ENCHANTED_MITHRIL",
	"Red Gift":                       "RED_GIFT",
	"Enchanted Raw Beef":             "ENCHANTED_RAW_BEEF",
	"Warped Stone":                   "AOTE_STONE",
	"Enchanted Feather":              "ENCHANTED_FEATHER",
	"Enchanted Oak Wood":             "ENCHANTED_OAK_LOG",
	"Rabbit Hide":                    "RABBIT_HIDE",
	"White Gift":                     "WHITE_GIFT",
	"Nether Wart":                    "NETHER_STALK",
	"Dark Bait":                      "DARK_BAIT",
	"Enchanted Pumpkin":              "ENCHANTED_PUMPKIN",
	"Enchanted Cooked Fish":          "ENCHANTED_COOKED_FISH",
	"Obsidian":                       "OBSIDIAN",
	"Golden Ball":                    "GOLDEN_BALL",
	"Starfall":                       "STARFALL",
	"Leather":                        "LEATHER",
	"Soul String":                    "SOUL_STRING",
	"Enchanted Bread":                "ENCHANTED_BREAD",
	"Fuming Potato Book":             "FUMING_POTATO_BOOK",
	"Flawed Sapphire Gemstone":       "FLAWED_SAPPHIRE_GEM",
	"Sulphur":                        "SULPHUR_ORE",
	"Glowstone Distillate":           "GLOWSTONE_DUST_DISTILLATE",
	"Perfect Sapphire Gemstone":      "PERFECT_SAPPHIRE_GEM",
	"Snow Block":                     "SNOW_BLOCK",
	"Enchanted Baked Potato":         "ENCHANTED_BAKED_POTATO",
	"Compactor":                      "COMPACTOR",
	"Mandraa":                        "MANDRAA",
	"Brown Mushroom":                 "BROWN_MUSHROOM",
	"Spooky Shard":                   "SPOOKY_SHARD",
	"Carrot":                         "CARROT_ITEM",
	"Enchanted Potato":               "ENCHANTED_POTATO",
	"Experience Bottle":              "EXP_BOTTLE",
	"Molten Cube":                    "MOLTEN_CUBE",
	"Enchanted Slimeball":            "ENCHANTED_SLIME_BALL",
	"Bulky Stone":                    "BULKY_STONE",
	"Enchanted Red Mushroom":         "ENCHANTED_RED_MUSHROOM",
	"Enchanted Rabbit Hide":          "ENCHANTED_RABBIT_HIDE",
	"Inferno Vertex":                 "INFERNO_VERTEX",
	"Flawed Amethyst Gemstone":       "FLAWED_AMETHYST_GEM",
	"Perfect Jade Gemstone":          "PERFECT_JADE_GEM",
	"Enchanted Birch Wood":           "ENCHANTED_BIRCH_LOG",
	"Enchanted Gunpowder":            "ENCHANTED_GUNPOWDER",
	"Refined Amber":                  "REFINED_AMBER",
	"Enchanted Sugar":                "ENCHANTED_SUGAR",
	"Cactus":                         "CACTUS",
	"Enchanted Cake":                 "ENCHANTED_CAKE",
	"Pumpkin":                        "PUMPKIN",
	"Wheat":                          "WHEAT",
	"Nurse Shark Tooth":              "NURSE_SHARK_TOOTH",
	"Enchanted Raw Salmon":           "ENCHANTED_RAW_SALMON",
	"Enchanted Emerald":              "ENCHANTED_EMERALD",
	"Enchanted Spider Eye":           "ENCHANTED_SPIDER_EYE",
	"Red Mushroom":                   "RED_MUSHROOM",
	"Grand Experience Bottle":        "GRAND_EXP_BOTTLE",
	"Mutton":                         "MUTTON",
	"Raw Soulflow":                   "RAW_SOULFLOW",
	"Raw Fish":                       "RAW_FISH",
	"Spider Eye":                     "SPIDER_EYE",
	"Yoggie":                         "YOGGIE",
	"Perfect Jasper Gemstone":        "PERFECT_JASPER_GEM",
	"Enchanted Netherrack":           "ENCHANTED_NETHERRACK",
	"Jaderald":                       "JADERALD",
	"Tightly-Tied Hay Bale":          "TIGHTLY_TIED_HAY_BALE",
	"Nether Wart Distillate":         "NETHER_STALK_DISTILLATE",
	"Prismarine Crystals":            "PRISMARINE_CRYSTALS",
	"Rare Diamond":                   "RARE_DIAMOND",
	"Ice":                            "ICE",
	"Brown Mushroom Block":           "HUGE_MUSHROOM_1",
	"Tiger Shark Tooth":              "TIGER_SHARK_TOOTH",
	"Red Mushroom Block":             "HUGE_MUSHROOM_2",
	"String":                         "STRING",
	"Golden Tooth":                   "GOLDEN_TOOTH",
	"Crude Gabagool":                 "CRUDE_GABAGOOL",
	"Dragon Scale":                   "DRAGON_SCALE",
	"Hyper Catalyst":                 "HYPER_CATALYST",
	"Rabbit's Foot":                  "RABBIT_FOOT",
	"Redstone":                       "REDSTONE",
	"Enchanted Cactus Green":         "ENCHANTED_CACTUS_GREEN",
	"Booster Cookie":                 "BOOSTER_COOKIE",
	"Enchanted Cookie":               "ENCHANTED_COOKIE",
	"Strong Dragon Fragment":         "STRONG_FRAGMENT",
	"Slimeball":                      "SLIME_BALL",
	"Mycelium":                       "MYCEL",
	"Snowball":                       "SNOW_BALL",
	"Holy Dragon Fragment":           "HOLY_FRAGMENT",
	"Enchanted Egg":                  "ENCHANTED_EGG",
	"Flawed Ruby Gemstone":           "FLAWED_RUBY_GEM",
	"Fine Jade Gemstone":             "FINE_JADE_GEM",
	"Raw Chicken":                    "RAW_CHICKEN",
	"Flawless Jasper Gemstone":       "FLAWLESS_JASPER_GEM",
	"Enchanted Shark Fin":            "ENCHANTED_SHARK_FIN",
	"Perfect Amber Gemstone":         "PERFECT_AMBER_GEM",
	"Dragon Horn":                    "DRAGON_HORN",
	"Ink Sack":                       "INK_SACK",
	"Flint":                          "FLINT",
	"Enchanted Spruce Wood":          "ENCHANTED_SPRUCE_LOG",
	"Hardened Wood":                  "HARDENED_WOOD",
	"Enchanted Redstone Block":       "ENCHANTED_REDSTONE_BLOCK",
	"Enchanted Quartz Block":         "ENCHANTED_QUARTZ_BLOCK",
	"Heavy Gabagool":                 "HEAVY_GABAGOOL",
	"Inferno Apex":                   "INFERNO_APEX",
	"Green Candy":                    "GREEN_CANDY",
	"Eccentric Painting":             "ECCENTRIC_PAINTING",
	"Enchanted Redstone Lamp":        "ENCHANTED_REDSTONE_LAMP",
	"Reaper Pepper":                  "REAPER_PEPPER",
	"Treasurite":                     "TREASURITE",
	"Perfect Opal Gemstone":          "PERFECT_OPAL_GEM",
	"Gravel":                         "GRAVEL",
	"Enchanted Packed Ice":           "ENCHANTED_PACKED_ICE",
	"Diamonite":                      "DIAMONITE",
	"Enchanted Prismarine Shard":     "ENCHANTED_PRISMARINE_SHARD",
	"Spirit Stone":                   "SPIRIT_DECOY",
	"Hypergolic Gabagool":            "HYPERGOLIC_GABAGOOL",
	"Enchanted Iron Block":           "ENCHANTED_IRON_BLOCK",
	"Bone":                           "BONE",
	"Revenant Flesh":                 "REVENANT_FLESH",
	"Enchanted Glowstone":            "ENCHANTED_GLOWSTONE",
	"Obsidian Tablet":                "OBSIDIAN_TABLET",
	"Vitamin Death":                  "VITAMIN_DEATH",
	"Netherrack":                     "NETHERRACK",
	"Blaze Rod":                      "BLAZE_ROD",
	"Young Dragon Fragment":          "YOUNG_FRAGMENT",
	"Candy Corn":                     "CANDY_CORN",
	"Midas Jewel":                    "MIDAS_JEWEL",
	"Premium Flesh":                  "PREMIUM_FLESH",
	"Refined Mineral":                "REFINED_MINERAL",
	"Rough Amethyst Gemstone":        "ROUGH_AMETHYST_GEM",
	"Rough Ruby Gemstone":            "ROUGH_RUBY_GEM",
	"Blue Goblin Egg":                "GOBLIN_EGG_BLUE",
	"Flawless Ruby Gemstone":         "FLAWLESS_RUBY_GEM",
	"Null Atom":                      "NULL_ATOM",
	"Catalyst":                       "CATALYST",
	"Blessed Bait":                   "BLESSED_BAIT",
	"Pure Mithril":                   "PURE_MITHRIL",
	"Enchanted Ink Sack":             "ENCHANTED_INK_SACK",
	"Jerry Stone":                    "JERRY_STONE",
	"Molten Powder":                  "MOLTEN_POWDER",
	"Rock Candy":                     "ROCK_CANDY",
	"Flawless Sapphire Gemstone":     "FLAWLESS_SAPPHIRE_GEM",
	"Magma Urchin":                   "MAGMA_URCHIN",
	"Oak Wood":                       "LOG",
	"Jacob's Ticket":                 "JACOBS_TICKET",
	"Searing Stone":                  "SEARING_STONE",
	"Absolute Ender Pearl":           "ABSOLUTE_ENDER_PEARL",
	"Enchanted Ender Pearl":          "ENCHANTED_ENDER_PEARL",
	"Spiked Bait":                    "SPIKED_BAIT",
	"Enchanted Fermented Spider Eye": "ENCHANTED_FERMENTED_SPIDER_EYE",
	"Rough Jade Gemstone":            "ROUGH_JADE_GEM",
	"Enchanted Jungle Wood":          "ENCHANTED_JUNGLE_LOG",
	"Blessed Fruit":                  "BLESSED_FRUIT",
	"Iron Ingot":                     "IRON_INGOT",
	"Null Ovoid":                     "NULL_OVOID",
	"Rough Sapphire Gemstone":        "ROUGH_SAPPHIRE_GEM",
	"Gold Ingot":                     "GOLD_INGOT",
	"Revenant Viscera":               "REVENANT_VISCERA",
	"Tarantula Silk":                 "TARANTULA_SILK",
	"Titanic Experience Bottle":      "TITANIC_EXP_BOTTLE",
	"Enchanted Iron":                 "ENCHANTED_IRON",
	"Super Compactor 3000":           "SUPER_COMPACTOR_3000",
	"Super Enchanted Egg":            "SUPER_EGG",
	"Mithril":                        "MITHRIL_ORE",
	"Enchanted Hay Bale":             "ENCHANTED_HAY_BLOCK",
	"Enchanted Paper":                "ENCHANTED_PAPER",
	"Enchanted Titanium":             "ENCHANTED_TITANIUM",
	"Spooky Bait":                    "SPOOKY_BAIT",
	"Hot Stuff":                      "HOT_STUFF",
	"Emerald":                        "EMERALD",
	"Enchanted Rabbit Foot":          "ENCHANTED_RABBIT_FOOT",
	"Enchanted Ice":                  "ENCHANTED_ICE",
	"Suspicious Vial":                "SUSPICIOUS_VIAL",
	"Arachne's Keeper Fragment":      "ARACHNE_KEEPER_FRAGMENT",
	"Dirt Bottle":                    "DIRT_BOTTLE",
	"Green Gift":                     "GREEN_GIFT",
	"Flawless Amethyst Gemstone":     "FLAWLESS_AMETHYST_GEM",
	"Red Nose":                       "RED_NOSE",
	"Rough Topaz Gemstone":           "ROUGH_TOPAZ_GEM",
	"Packed Ice":                     "PACKED_ICE",
	"Hamster Wheel":                  "HAMSTER_WHEEL",
	"Enchanted Obsidian":             "ENCHANTED_OBSIDIAN",
	"Enchanted Coal":                 "ENCHANTED_COAL",
	"Rough Opal Gemstone":            "ROUGH_OPAL_GEM",
	"Flawless Opal Gemstone":         "FLAWLESS_OPAL_GEM",
	"Werewolf Skin":                  "WEREWOLF_SKIN",
	"Daedalus Stick":                 "DAEDALUS_STICK",
	"Dragon Claw":                    "DRAGON_CLAW",
	"Enchanted Red Sand Cube":        "ENCHANTED_RED_SAND_CUBE",
	"Enchanted Raw Fish":             "ENCHANTED_RAW_FISH",
	"Lucky Dice":                     "LUCKY_DICE",
	"Foul Flesh":                     "FOUL_FLESH",
	"Raw Beef":                       "RAW_BEEF",
	"Enchanted Eye of Ender":         "ENCHANTED_EYE_OF_ENDER",
	"Ectoplasm":                      "ECTOPLASM",
	"Fuel Gabagool":                  "FUEL_GABAGOOL",
	"Shark Bait":                     "SHARK_BAIT",
	"Precious Pearl":                 "PRECIOUS_PEARL",
	"Blue Jerry Box":                 "JERRY_BOX_BLUE",
	"Enchanted Slime Block":          "ENCHANTED_SLIME_BLOCK",
	"Scorched Books":                 "SCORCHED_BOOKS",
	"Enchanted Mycelium":             "ENCHANTED_MYCELIUM",
	"Fine Sapphire Gemstone":         "FINE_SAPPHIRE_GEM",
	"Raw Rabbit":                     "RABBIT",
	"Gunpowder":                      "SULPHUR",
	"Enchanted Carrot":               "ENCHANTED_CARROT",
	"Griffin Feather":                "GRIFFIN_FEATHER",
	"Rotten Flesh":                   "ROTTEN_FLESH",
	"Soulflow":                       "SOULFLOW",
	"Minnow Bait":                    "MINNOW_BAIT",
	"Enchanted Magma Cream":          "ENCHANTED_MAGMA_CREAM",
	"Enchanted Firework Rocket":      "ENCHANTED_FIREWORK_ROCKET",
	"Wither Blood":                   "WITHER_BLOOD",
	"Blaze Wax":                      "BLAZE_WAX",
	"Flawless Jade Gemstone":         "FLAWLESS_JADE_GEM",
	"Hard Stone":                     "HARD_STONE",
	"Flawed Topaz Gemstone":          "FLAWED_TOPAZ_GEM",
	"Lapis Crystal":                  "LAPIS_CRYSTAL",
	"Enchanted Cooked Mutton":        "ENCHANTED_COOKED_MUTTON",
	"Fine Amethyst Gemstone":         "FINE_AMETHYST_GEM",
	"Refined Titanium":               "REFINED_TITANIUM",
	"Shiny Prism":                    "SHINY_PRISM",
	"Red Sand":                       "SAND:1",
	"Enchanted Raw Rabbit":           "ENCHANTED_RABBIT",
	"Mutant Nether Wart":             "MUTANT_NETHER_STALK",
	"Fine Jasper Gemstone":           "FINE_JASPER_GEM",
	"Flawless Amber Gemstone":        "FLAWLESS_AMBER_GEM",
	"Enchanted Charcoal":             "ENCHANTED_CHARCOAL",
	"Enchanted Blaze Powder":         "ENCHANTED_BLAZE_POWDER",
	"Summoning Eye":                  "SUMMONING_EYE",
	"Enchanted Sulphur Cube":         "ENCHANTED_SULPHUR_CUBE",
	"Fish Bait":                      "FISH_BAIT",
	"Enchanted Diamond":              "ENCHANTED_DIAMOND",
}
