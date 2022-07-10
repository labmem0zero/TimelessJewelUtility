package main

import "fmt"

type TimelessJewel struct {
	JewelName           string
	ConqName            string
	RaceID              uint32
	ConqID              uint32
	ConqVer             uint32
	LowRoll             uint32
	HighRol             uint32
	Seed                uint32
	SkillGroups         map[string]SkillGroup
	Rng                 Rngesus
	SkillWorker         SkillWorker
	ApplicableSkills    []AlternatePassiveSkill
	ApplicableAdditions []AlternatePassiveAddition
}

type SkillWorker struct {
	tj *TimelessJewel
}

type TimelessJewels struct {
	Jewels              []TimelessJewel
	ReplaceableSkills   []AlternatePassiveSkill
	AdditionableeSkills []AlternatePassiveAddition
}

var Jewels TimelessJewels

func (tjs *TimelessJewels) AddNewJewel(raceId, ConqID, Seed uint32) {
	tj := TimelessJewel{}
	tj.SkillGroups = make(map[string]SkillGroup)

	tj.JewelName = JewelsAndConquerors[raceId][ConqID].JewelName
	tj.ConqName = JewelsAndConquerors[raceId][ConqID].ConqName
	tj.RaceID = raceId
	tj.ConqID = ConqID
	tj.ConqVer = JewelsAndConquerors[raceId][ConqID].Ver
	tj.LowRoll = JewelsAndConquerors[raceId][ConqID].Min
	tj.HighRol = JewelsAndConquerors[raceId][ConqID].Max
	tj.Seed = Seed
	if tj.RaceID == 5 {
		tj.Seed = tj.Seed / 20
	}
	tj.SkillGroups = SkillGroups
	tj.SkillWorker.tj = &tj

	tjs.Jewels = append(tjs.Jewels, tj)
}

func (tjs *TimelessJewels) AddJewels() {
	for i := JewelsAndConquerors[race][conq].Min; i < JewelsAndConquerors[race][conq].Max+1; i++ {
		tjs.AddNewJewel(race, conq, i)
	}
}

type FoundSkillGroup map[string]ResSkill

func (tj *TimelessJewel) RollOverSkills() map[string]FoundSkillGroup {
	skillsGroupsFound := make(map[string]FoundSkillGroup)
	for skillGroupName, skillGroup := range tj.SkillGroups {
		founds := []SkillGroupSkill{}
		statsFound := make(map[string]byte)
		for _, skill := range skillGroup.Skills {
			foundSkill := SkillGroupSkill{}
			newSkill := tj.ProcessSkill(skill)

			if statName := IsSkillPresented(NeededSkills, newSkill.NewSkillName); statName != "" {
				statsFound[statName]++
				foundSkill = newSkill
			}

			for _, stat := range newSkill.Stats {
				if statName := IsStatPresented(NeededStats, stat); statName != "" {
					statsFound[statName]++
					foundSkill = newSkill
				}
			}
			if foundSkill.NewSkillName != "" {
				founds = append(founds, foundSkill)
			}
		}
		check := 0
		for _, n := range NeededSkills {
			if statsFound[n.id] >= n.count {
				check++
			}
		}
		for _, n := range NeededStats {
			if statsFound[n.id] >= n.count {
				check++
			}
		}
		if check == len(NeededSkills)+len(NeededStats) {
			fs := make(map[string]ResSkill)
			for _, v := range founds {
				fs[v.OldSkillName] = v.SgsToRs()
			}
			skillsGroupsFound[skillGroupName] = fs
		}
	}
	return skillsGroupsFound
}

func (tjs *TimelessJewels) RollOverSeeds() map[uint32]map[string]FoundSkillGroup {
	results := make(map[uint32]map[string]FoundSkillGroup)
	for _, tj := range tjs.Jewels {
		res := tj.RollOverSkills()
		if len(res) > 0 {
			var seed uint32
			if tj.RaceID == 5 {
				seed = tj.Seed * 20
			} else {
				seed = tj.Seed
			}
			results[seed] = make(map[string]FoundSkillGroup)
			results[seed] = res
		}
	}
	return results
}

type FoundSkill struct {
	params []string
}

type FoundSkills struct {
	skills      map[string]FoundSkill
	skillsQuant map[string]byte
}

func (tjs *TimelessJewels) FindAllPassives(race uint32, seed uint32) {
	tjs.AddNewJewel(race, 1, seed)
	results := make(map[string]FoundSkills)
	for _, tj := range tjs.Jewels {
		if race == 5 {
			seed = seed / 20
		}
		fmt.Println("Seed:", tj.Seed)
		if tj.Seed == seed {
			for skillGroupName, skillGroup := range tj.SkillGroups {
				fs := FoundSkills{}
				fs.skills = make(map[string]FoundSkill)
				fs.skillsQuant = make(map[string]byte)
				for _, skill := range skillGroup.Skills {
					newskill := tj.ProcessSkill(skill)
					if skillGroupName == "Left Scion" {
						fmt.Println(newskill.OldSkillName, " -> ", newskill.Stats)
					}
					fso := FoundSkill{}
					for _, v := range newskill.Stats {
						fso.params = append(fso.params, v.StatName)
						fs.skillsQuant[v.StatName]++
					}
					fs.skills[newskill.OldSkillName] = fso
				}
				results[skillGroupName] = fs
			}
		}
	}
	fmt.Println("Ended")
	for i, j := range results {
		fmt.Println(i, ":")
		res := SortSkills(j.skillsQuant)
		for _, v := range res {
			fmt.Println("\t", replacer.Replace(v), " - ", j.skillsQuant[v])
		}
	}
}
