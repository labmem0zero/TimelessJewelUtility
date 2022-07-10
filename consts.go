package main

import "strings"

const (
	statsFile        = "data/stats.json"
	passivesFile     = "data/passive_skills.json"
	altTreeVerFile   = "data/alternate_tree_versions.json"
	altPassivesFile  = "data/alternate_passive_skills.json"
	altAdditionsFile = "data/alternate_passive_additions.json"
	skillgroups      = "skillGroups.txt"
)

var NodeRaces map[uint32]string

type ConqInfo struct {
	JewelName string
	ConqName  string
	Ver       uint32
	Min, Max  uint32
}

type Conquerors map[uint32]ConqInfo

var JewelsAndConquerors map[uint32]Conquerors

func InitConqs() {
	NodeRaces = make(map[uint32]string)
	NodeRaces[1] = "vaal"
	NodeRaces[2] = "karui"
	NodeRaces[3] = "maraketh"
	NodeRaces[4] = "templar"
	NodeRaces[5] = "eternal"

	JewelsAndConquerors = make(map[uint32]Conquerors)
	for i := 1; i < 6; i++ {
		JewelsAndConquerors[uint32(i)] = make(map[uint32]ConqInfo)
	}
	JewelsAndConquerors[1][1] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Glorious Vanity", ConqName: "Xibaqua", Ver: 0, Min: 100, Max: 8000}
	JewelsAndConquerors[1][2] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Glorious Vanity", ConqName: "Ahuana", Ver: 1, Min: 100, Max: 8000}
	JewelsAndConquerors[1][3] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Glorious Vanity", ConqName: "Doryani", Ver: 0, Min: 100, Max: 8000}

	JewelsAndConquerors[2][1] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Lethal Pride", ConqName: "Kaom", Ver: 0, Min: 10000, Max: 18000}
	JewelsAndConquerors[2][2] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Lethal Pride", ConqName: "Rakiata", Ver: 0, Min: 10000, Max: 18000}
	JewelsAndConquerors[2][3] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Lethal Pride", ConqName: "Akoya", Ver: 1, Min: 10000, Max: 18000}

	JewelsAndConquerors[3][1] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Brutal Restraint", ConqName: "Balbala", Ver: 1, Min: 500, Max: 8000}
	JewelsAndConquerors[3][2] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Brutal Restraint", ConqName: "Asenath", Ver: 0, Min: 500, Max: 8000}
	JewelsAndConquerors[3][3] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Brutal Restraint", ConqName: "Nasima", Ver: 0, Min: 500, Max: 8000}

	JewelsAndConquerors[4][1] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Militant Faith", ConqName: "Maxarius", Ver: 1, Min: 2000, Max: 10000}
	JewelsAndConquerors[4][2] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Militant Faith", ConqName: "Dominus", Ver: 0, Min: 2000, Max: 10000}
	JewelsAndConquerors[4][3] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Militant Faith", ConqName: "Avarius", Ver: 0, Min: 2000, Max: 10000}

	JewelsAndConquerors[5][1] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Elegant Hubris", ConqName: "Cadiro", Ver: 0, Min: 2000, Max: 10000}
	JewelsAndConquerors[5][2] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Elegant Hubris", ConqName: "Victario", Ver: 0, Min: 2000, Max: 10000}
	JewelsAndConquerors[5][3] = struct {
		JewelName string
		ConqName  string
		Ver       uint32
		Min, Max  uint32
	}{JewelName: "Elegant Hubris", ConqName: "Caspiro", Ver: 3, Min: 100, Max: 8000}
	JewelNames = []string{"Glorious Vanity", "Lethal Pride", "Brutal Restraint", "Militant Faith", "Elegant Hubris"}
	Jewels = TimelessJewels{}
}

var replacer = strings.NewReplacer(
	"maximum_energy_shield_+%", "Maximum ES",
	"energy_shield_recharge_rate_+%", "Energy shield recharge rate",
	"physical_damage_reduction_rating_+%", "Armour",
	"base_additional_physical_damage_reduction_%", "Physical damage reduction",
	"fire_damage_+%", "Increased fire damage",
	"base_life_leech_from_fire_damage_permyriad", "Fire damage leeched as life",
	"base_lightning_damage_resistance_%", "Lightning resistance",
	"base_maximum_lightning_damage_resistance_%", "Maximum lightning resistance",
	"base_chaos_damage_resistance_%", "Chaos resistance",
	"base_maximum_chaos_damage_resistance_%", "Maximum chaos resistance",
	"base_spell_suppression_chance_%", "Spell suppression chance",
	"base_resist_all_elements_%", "All elemental resistances",
	"lightning_damage_+%", "Increased lightning damage",
	"base_physical_damage_%_to_convert_to_lightning", "Phys converted to lightning",
	"curse_effect_+%", "Curse effect",
	"curse_skill_effect_duration_+%", "Curse duration",
	"non_curse_aura_effect_+%", "Non curse aura effect",
	"base_aura_area_of_effect_+%", "Aura AoE radius",
	"base_life_leech_from_lightning_damage_permyriad", "Lightning damage leeched as life",
	"maximum_life_+%", "Increased maximum life",
	"base_life_leech_from_attack_damage_permyriad", "Attak damage leeched as life",
	"physical_damage_+%", "Increased physical damage",
	"chance_to_deal_double_damage_%", "Chance to double damage",
	"life_regeneration_rate_per_minute_%", "Life regeneration rate",
	"base_life_leech_from_physical_damage_permyriad", "Physical damage leeched as life",
	"base_physical_damage_%_to_convert_to_fire", "Phys converted to fire",
	"chaos_damage_+%", "Increased chaos damage",
	"base_life_leech_from_chaos_damage_permyriad", "Chaos damage leeched as life",
	"damage_over_time_+%", "Increased DoT",
	"skill_effect_duration_+%", "Skill effect duration",
	"minion_damage_+%", "Increased minion damage",
	"minion_maximum_life_+%", "Increased minion maximum life",
	"withered_on_hit_for_2_seconds_%_chance", "Chance to inflict withered on hit",
	"base_reduce_enemy_lightning_resistance_%", "Lightning penetration",
	"base_energy_shield_leech_from_spell_damage_permyriad", "Spell damage leeched as ES",
	"additional_block_%", "Chance to block attacks",
	"life_gained_on_block", "Life gained on block",
	"cold_damage_+%", "Increased cold damage",
	"base_reduce_enemy_cold_resistance_%", "Cold penetration",
	"base_reduce_enemy_fire_resistance_%", "Fire penetration",
	"base_cold_damage_resistance_%", "Cold resistance",
	"base_maximum_cold_damage_resistance_%", "Maximum cold resistance",
	"maximum_mana_+%", "Increased maximummana",
	"mana_regeneration_rate_+%", "Mana regeneration rate",
	"base_spell_block_%", "Chance to block spells",
	"shield_armour_+%", "Armor from equipped shield",
	"avoid_all_elemental_status_%", "Chance to avoid elemental ailments",
	"base_avoid_stun_%", "Chance to avoid being stunned",
	"evasion_rating_+%", "Increased evasion",
	"global_chance_to_blind_on_hit_%", "Chance to lbind on hit",
	"base_fire_damage_resistance_%", "Fire resistance",
	"base_maximum_fire_damage_resistance_%", "Maximum fire resistance",
	"spell_damage_+%", "Increased spell damage",
	"spell_critical_strike_chance_+%", "Increased spell crit chance",
	"base_life_leech_from_cold_damage_permyriad", "Cold damage leeched as life",
	"faster_bleed_%", "Bleedings faster",
	"base_physical_damage_%_to_convert_to_cold", "Phys converted to cold",
	"vaal_small_fire_damage", "Increased fire damage(small)",
	"vaal_small_cold_damage", "Increased cold damage(small)",
	"vaal_small_lightning_damage", "Increased lightning damage(small)",
	"vaal_small_physical_damage", "Increased ogysical damage(small)",
	"vaal_small_chaos_damage", "Increased chaos damage(small)",
	"vaal_small_minion_damage", "Increased minion damage(small)",
	"vaal_small_attack_damage", "Increased attack damage(small)",
	"vaal_small_spell_damage", "Increased spell damage(small)",
	"vaal_small_area_damage", "Increased AoE damage(small)",
	"vaal_small_projectile_damage", "Increased fire damage(small)",
	"vaal_small_damage_over_time", "Increased DoT(small)",
	"vaal_small_area_of_effect", "Increased AoE(small)",
	"vaal_small_projectile_speed", "Increased projectile speed(small)",
	"vaal_small_critical_strike_chance", "Increased critc. chance(small)",
	"vaal_small_critical_strike_multiplier", "Increased crit. mult(small)",
	"vaal_small_attack_speed", "Increased attack speed (small)",
	"vaal_small_cast_speed", "Increased cast speed (small)",
	"vaal_small_movement_speed", "Increased movement speed(small)",
	"vaal_small_chance_to_ignite", "Chance to ignite(small)",
	"vaal_small_chance_to_freeze", "Chance to freeze(small)",
	"vaal_small_chance_to_shock", "Chance to shock(small)",
	"vaal_small_duration", "Increased skill effect duration(small)",
	"vaal_small_life", "Maximum life(small)",
	"vaal_small_mana", "Maximum mana(small)",
	"vaal_small_mana_regeneration", "Increased mana regen(small)",
	"vaal_small_armour", "Increased armor(small)",
	"vaal_small_evasion", "Increased evasion(small)",
	"vaal_small_energy_shield", "Increased ES(small)",
	"vaal_small_attack_block", "Chance to block attacks(small)",
	"vaal_small_spell_block", "Chance to block spells(small)",
	"vaal_small_attack_dodge", "Chance to dodge attacks(small)",
	"vaal_small_spell_dodge", "Chance to dodge spells(small)",
	"vaal_small_aura_effect", "Increased aura effect(small)",
	"vaal_small_curse_effect", "Increased curse effect(small)",
	"vaal_small_fire_resistance", "Fire res(small)",
	"vaal_small_cold_resistance", "Cold res(small)",
	"vaal_small_lightning_resistance", "Lightning res(small)",
	"vaal_small_chaos_resistance", "Chaos res(small)",
	"vaal_notable_fire_damage_1", "Increased fire damage. Fire pen.",
	"vaal_notable_fire_damage_2", "Increased fire damage. Fire leech",
	"vaal_notable_fire_damage_3", "Increased fire damage. Phys convert",
	"vaal_notable_cold_damage_1", "Increased cold damage. Cold pen.",
	"vaal_notable_cold_damage_2", "Increased cold damage. Cold leech",
	"vaal_notable_cold_damage_3", "Increased cold damage. Phys convert",
	"vaal_notable_lightning_damage_1", "Increased lightning damage. Lightning pen.",
	"vaal_notable_lightning_damage_2", "Increased lightning damage. Lightning leech",
	"vaal_notable_lightning_damage_3", "Increased lightning damage. Phys convert",
	"vaal_notable_physical_damage_1", "Increased physical damage. Chance to double damage",
	"vaal_notable_physical_damage_2", "Increased physical damage. Phys leech",
	"vaal_notable_physical_damage_3", "Increased physical damage. Faster bleeding",
	"vaal_notable_chaos_damage_1", "Increased chaos damage. Chaos leech",
	"vaal_notable_chaos_damage_2", "Increased chaos damage. Withered on hit",
	"vaal_notable_spell_damage_1", "Increased spell damage. Spell crit",
	"vaal_notable_minion_damage_1", "Increased minion damage. Minion life",
	"vaal_notable_damage_over_time_1", "Increased DoT. Skill effect duration",
	"vaal_notable_life_1", "Increased maximum life. Life regen",
	"vaal_notable_life_2", "Increased maximum life. Attack life leech",
	"vaal_notable_mana_1", "Increased maximum mana. Mana regen rate",
	"vaal_notable_armour_1", "Increased armor. Phys damage reduction",
	"vaal_notable_evasion_1", "Increased evasion. Chance to blind",
	"vaal_notable_energy_shield_1", "Increased ES. ES recharge rate",
	"vaal_notable_energy_shield_2", "Increased ES. Spell damage ES leech",
	"vaal_notable_block_1", "Chance to block. Life gain on block",
	"vaal_notable_block_2", "Chance to block. Defences from shield",
	"vaal_notable_dodge_1", "Chance to dodge attacks. Chance to avoid stuns",
	"vaal_notable_dodge_2", "Chance to dodge attacks. All res",
	"vaal_notable_aura_1", "Aura effect. Aura AoE",
	"vaal_notable_curse_1", "Curse effect. Curse duration",
	"vaal_notable_fire_resistance_1", "Fire res. Max fire res",
	"vaal_notable_cold_resistance_1", "Cold res. Max cold res",
	"vaal_notable_lightning_resistance_1", "Lightning res. Max lightning res",
	"vaal_notable_chaos_resistance_1", "Chaos res. Max chaos res",
	"vaal_notable_random_offense", "Random offence",
	"vaal_notable_random_defence", "Random deffence",
	"templar_devotion_node", "Add devotion",
	"templar_notable_fire_conversion", "Phys to fire convert",
	"templar_notable_cold_conversion", "Phys to cold convert",
	"templar_notable_lightning_conversion", "Phys to lightning convert",
	"templar_notable_mana_added_as_energy_shield", "Add mana as es",
	"templar_notable_arcane_surge", "Arcane surge",
	"templar_notable_minimum_endurance_charge", "Minimum endurance charges",
	"templar_notable_minimum_frenzy_charge", "Minimum frenzy charges",
	"templar_notable_minimum_power_charge", "Minimum pwoer charges",
	"templar_notable_consecrated_ground_ailments", "Consecrated ground ailments immunity",
	"templar_notable_additional_physical_reduction", "Phys reduction",
	"templar_notable_max_resistances", "All max res",
	"templar_notable_fire_exposure", "Chance to fire exposure",
	"templar_notable_cold_exposure", "Chance to cold exposure",
	"templar_notable_lightning_exposure", "Chance to lightning exposure",
	"eternal_notable_crit_1", "Crit chance",
	"eternal_notable_crit_2", "Crit mult",
	"eternal_notable_endurance_1", "Gain endurance charges when hit",
	"eternal_notable_endurance_2", "Phys reduction per endurance charge",
	"eternal_notable_endurance_3", "Damage per endurance charge",
	"eternal_notable_frenzy_1", "Frenzy charge on hit",
	"eternal_notable_frenzy_2", "Evasion per frenzy charge",
	"eternal_notable_frenzy_3", "Damage per frenzy charge",
	"eternal_notable_power_1", "Power charge on hit",
	"eternal_notable_power_2", "ES per frenzy charge",
	"eternal_notable_power_3", "Damage per power charge",
	"eternal_notable_chill_1", "Chill effect",
	"eternal_notable_chill_2", "Avoid chilled",
	"eternal_notable_shock_1", "Shock effect",
	"eternal_notable_shock_2", "Avoid shock",
	"eternal_notable_block_1", "Block attack chance",
	"eternal_notable_block_2", "Block spell chance",
	"eternal_notable_dodge_1", "Avoid elemental ailments",
	"eternal_notable_dodge_2", "Spell suppression chance",
	"eternal_notable_aura_1", "Aura effect",
	"eternal_notable_minion_1", "Minion damage",
	"eternal_notable_minion_2", "Minion life",
	"eternal_notable_spell_1", "Spell damage",
	"eternal_notable_spell_2", "Spell crit",
	"eternal_notable_fire_attack_1", "Fire attack damage",
	"eternal_notable_cold_attack_1", "Cold attack damage",
	"eternal_notable_lightning_attack_1", "Lightning attack damage",
	"eternal_notable_physical_damage_1", "Physical damage",
	"eternal_notable_physical_damage_2", "Melee physical damage",
	"eternal_notable_bleed_damage_1", "Bleeding damage",
	"eternal_notable_projectile_attack_damage_1", "Projectile damage",
	"eternal_notable_attack_speed_1", "Attack speed",
	"eternal_notable_cast_speed_1", "Cast speed",
	"eternal_notable_rarity_1", "Item rarity",
	"eternal_notable_armour_1", "Armour",
	"eternal_notable_evasion_1", "Evasion",
	"eternal_notable_fire_resistance_1", "Fire res",
	"eternal_notable_cold_resistance_1", "Cold res",
	"eternal_notable_lightning_resistance_1", "Lightning res",
	"eternal_notable_chaos_resistance_1", "Chaos res",
	"eternal_notable_life_1", "Increased life",
	"eternal_notable_mana_1", "Increased mana",
	"eternal_notable_mana_regen_1", "Mana regen",
	"eternal_notable_accuracy_1", "Accuracy",
	"eternal_notable_flask_duration_1", "Flask duration",
	"vaal_small_fire_damage", "Increased fire damage",
	"vaal_small_cold_damage", "Increased cold damage",
	"vaal_small_lightning_damage", "Increased lightning damage",
	"vaal_small_physical_damage", "Increased phys damage",
	"vaal_small_chaos_damage", "Increased chaos damage",
	"vaal_small_minion_damage", "Increased minion damage",
	"vaal_small_attack_damage", "Increased a damage",
	"vaal_small_spell_damage", "Increased spell damage",
	"vaal_small_area_damage", "Increased AoE damage",
	"vaal_small_projectile_damage", "Increased fire damage",
	"vaal_small_damage_over_time", "Increased DoT",
	"vaal_small_area_of_effect", "Increased AoE",
	"vaal_small_projectile_speed", "Increased projectile speed",
	"vaal_small_critical_strike_chance", "Global crit chance",
	"vaal_small_critical_strike_multiplier", "Global crit mult",
	"vaal_small_attack_speed", "Increased attack speed",
	"vaal_small_cast_speed", "Increased cast speed",
	"vaal_small_movement_speed", "Increased movement speed",
	"vaal_small_chance_to_ignite", "Chance to ignite",
	"vaal_small_chance_to_freeze", "Chance to freeze",
	"vaal_small_chance_to_shock", "Chance to shock",
	"vaal_small_duration", "Increased skill effect duration",
	"vaal_small_life", "Increased life",
	"vaal_small_mana", "Increased mana",
	"vaal_small_mana_regeneration", "Mana regen rate",
	"vaal_small_armour", "Increased armour",
	"vaal_small_evasion", "Increased evasion",
	"vaal_small_energy_shield", "Increased ES",
	"vaal_small_attack_block", "Chance to block attacks",
	"vaal_small_spell_block", "Chance to block spells",
	"vaal_small_attack_dodge", "Chance to dodge attacks",
	"vaal_small_spell_dodge", "Chance to suppress spell damage",
	"vaal_small_aura_effect", "Increased aura effect",
	"vaal_small_curse_effect", "Increased curse effect",
	"vaal_small_fire_resistance", "Fire res",
	"vaal_small_cold_resistance", "Cold res",
	"vaal_small_lightning_resistance", "Lightning res",
	"vaal_small_chaos_resistance", "Chaos res",
	"karui_attribute_strength", "Added strength",
	"karui_notable_add_strength", "Added strength",
	"karui_notable_add_percent_strength", "Increased strength",
	"karui_notable_add_armour", "Increased armour",
	"karui_notable_add_leech", "Attack leech",
	"karui_notable_add_double_damage", "Chance to deal double damage",
	"karui_notable_add_life", "Increased life",
	"karui_notable_add_fortify_effect", "Max fortify",
	"karui_notable_add_life_regen", "Life regen",
	"karui_notable_add_fire_resistance", "Fire res",
	"karui_notable_add_melee_damage", "Melee damage",
	"karui_notable_add_damage_from_crits", "Reduced damage from crits taken",
	"karui_notable_add_melee_crit_chance", "Melee crit chance",
	"karui_notable_add_burning_damage", "Burning damage",
	"karui_notable_add_totem_damage", "Totem damage",
	"karui_notable_add_melee_crit_multi", "Melee crit mult",
	"karui_notable_add_physical_damage", "Increased phys damage",
	"karui_notable_add_warcry_buff_effect", "Warcrt buff effect",
	"karui_notable_add_totem_placement_speed", "Totem placement speed",
	"karui_notable_add_stun_duration", "Stun duration",
	"karui_notable_add_faster_burn", "Ignites faster",
	"karui_notable_add_reduced_stun_threshold", "Enemy stun threshold",
	"karui_notable_add_physical_added_as_fire", "Gain phys extra fire",
	"karui_notable_add_physical_taken_as_fire", "Phys taken as fire",
	"karui_notable_add_endurance_charge_on_kill", "Endurance charge on kill",
	"karui_notable_add_intimidate", "Intimidiate",
	"maraketh_attribute_dex", "Dex",
	"maraketh_notable_add_dexterity", "Dex",
	"maraketh_notable_add_percent_dexterity", "Percent dex",
	"maraketh_notable_add_evasion", "Evasion",
	"maraketh_notable_add_flask_charges", "Flask charges",
	"maraketh_notable_add_speed", "Attack and cast speed",
	"maraketh_notable_add_life", "Increased life",
	"maraketh_notable_add_blind", "Blind on hit",
	"maraketh_notable_add_movement_speed", "Movement speed",
	"maraketh_notable_add_cold_resistance", "Cold res",
	"maraketh_notable_add_projectile_damage", "Projectile damage",
	"maraketh_notable_add_ailment_avoid", "Avoid stun",
	"maraketh_notable_add_global_crit_chance", "Global crit chance",
	"maraketh_notable_add_poison_damage", "Posion damage",
	"maraketh_notable_add_minion_damage", "Minion damage",
	"maraketh_notable_add_accuracy", "Accuracy",
	"maraketh_notable_add_elemental_damage", "Elemental damage",
	"maraketh_notable_add_aura_effect", "Aura effect",
	"maraketh_notable_add_minion_movement_speed", "Minion movement speed",
	"maraketh_notable_add_ailment_duration", "Elemental ailment duration",
	"maraketh_notable_add_faster_poison", "Faster posion",
	"maraketh_notable_add_ailment_effect", "Elemental ailment effect",
	"maraketh_notable_add_physical_added_as_cold", "Gain phys extra cold",
	"maraketh_notable_add_flask_effect", "Alchemist's genius",
	"maraketh_notable_add_frenzy_charge_on_kill", "Frenzy charge on kill",
	"maraketh_notable_add_onslaught", "Onslaught",
)
