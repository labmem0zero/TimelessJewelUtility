package main

type Stat struct {
	Rid                    uint32        `json:"_rid"`
	Id                     string        `json:"Id"`
	Text                   string        `json:"Text"`
	BelongsActiveSkillsKey []interface{} `json:"BelongsActiveSkillsKey"`
	Category               uint32        `json:"Category"`
}

type PassiveSkill struct {
	Rid                 uint32      `json:"_rid"`
	Id                  string      `json:"Id"`
	Stats               []uint32    `json:"Stats"`
	PassiveSkillGraphId uint32      `json:"PassiveSkillGraphId"`
	Name                string      `json:"Name"`
	IsKeystone          bool        `json:"IsKeystone"`
	IsNotable           bool        `json:"IsNotable"`
	IsJewelSocket       bool        `json:"IsJewelSocket"`
	SkillType           uint32      `json:"SkillType"`
	MasteryGroup        interface{} `json:"MasteryGroup"`
}

type AlternateTreeVersion struct {
	Rid                      uint32 `json:"_rid"`
	Id                       string `json:"Id"`
	IsSmallAttributeReplaced bool   `json:"Unknown2"`
	IsSmallNormalReplaced    bool   `json:"Unknown3"`
	MinAddsCount             uint32 `json:"Unknown6"`
	MaxAddsCount             uint32 `json:"Unknown7"`
	ReplaceSpawnWeight       uint32 `json:"Unknown10"`
}

type AlternatePassiveSkill struct {
	Rid                      uint32   `json:"_rid"`
	Id                       string   `json:"Id"`
	AlternateTreeVersionsKey uint32   `json:"AlternateTreeVersionsKey"`
	Name                     string   `json:"Name"`
	PassiveType              []uint32 `json:"PassiveType"`
	StatsKeys                []uint32 `json:"StatsKeys"`
	Stat1MinRoll             uint32   `json:"Stat1Min"`
	Stat1MaxRoll             uint32   `json:"Stat1Max"`
	Stat2MinRoll             uint32   `json:"Stat2Min"`
	Stat2MaxRoll             uint32   `json:"Stat2Max"`
	Stat3MinRoll             uint32   `json:"Unknown10"`
	Stat3MaxRoll             uint32   `json:"Unknown11"`
	Stat4MinRoll             uint32   `json:"Unknown12"`
	Stat4MaxRoll             uint32   `json:"Unknown13"`
	SpawnWeight              uint32   `json:"SpawnWeight"`
	ConqID                   uint32   `json:"Unknown19"`
	RandomMin                uint32   `json:"RandomMin"`
	RandomMax                uint32   `json:"RandomMax"`
	ConqVersion              uint32   `json:"Unknown25"`
}

type AlternatePassiveAddition struct {
	Rid                      uint32   `json:"_rid"`
	Id                       string   `json:"Id"`
	AlternateTreeVersionsKey uint32   `json:"AlternateTreeVersionsKey"`
	SpawnWeight              uint32   `json:"SpawnWeight"`
	StatsKeys                []uint32 `json:"StatsKeys"`
	Stat1MinRoll             uint32   `json:"Stat1Min"`
	Stat1MaxRoll             uint32   `json:"Stat1Max"`
	Stat2MinRoll             uint32   `json:"Unknown7"`
	Stat2MaxRoll             uint32   `json:"Unknown8"`
	PassiveType              []uint32 `json:"PassiveType"`
}

type SkillStat struct {
	StatName         string
	StatRoll         uint32
	MinRoll, MaxRoll uint32
}

type SkillGroupSkill struct {
	Skill            PassiveSkill `json:",omitempty"`
	Rid              uint32       `json:",omitempty"`
	OldSkillName     string
	NewSkillName     string
	MinAdds, MaxAdds uint32 `json:",omitempty"`
	Stats            []SkillStat
}

type SkillGroup struct {
	Name   string
	Skills map[string]SkillGroupSkill
}

func (sgs *SkillGroupSkill) AddAlts(alt AlternatePassiveSkill) {
	sgs.NewSkillName = alt.Name
	sgs.MinAdds = alt.RandomMin
	sgs.MaxAdds = alt.RandomMax
	if len(alt.StatsKeys) > 0 {
		sgs.Stats = append(sgs.Stats, SkillStat{Stats[alt.StatsKeys[0]].Id, 0, alt.Stat1MinRoll, alt.Stat1MaxRoll})
	}
	if len(alt.StatsKeys) > 1 {
		sgs.Stats = append(sgs.Stats, SkillStat{Stats[alt.StatsKeys[1]].Id, 0, alt.Stat2MinRoll, alt.Stat2MaxRoll})
	}
	if len(alt.StatsKeys) > 2 {
		sgs.Stats = append(sgs.Stats, SkillStat{Stats[alt.StatsKeys[2]].Id, 0, alt.Stat3MinRoll, alt.Stat4MaxRoll})
	}
	if len(alt.StatsKeys) > 3 {
		sgs.Stats = append(sgs.Stats, SkillStat{Stats[alt.StatsKeys[3]].Id, 0, alt.Stat4MinRoll, alt.Stat4MaxRoll})
	}
}

func (sgs *SkillGroupSkill) AddAdd(alt AlternatePassiveAddition, roll uint32) {
	sgs.Stats = append(sgs.Stats, SkillStat{alt.Id, roll, alt.Stat1MaxRoll, alt.Stat1MaxRoll})
}

type ResSkill struct {
	NewSkillName string
	Stats        map[string]ResSkillStat
}

type ResSkillStat struct {
	Roll uint32
}

func (skill *SkillGroupSkill) SgsToRs() ResSkill {
	res := ResSkill{}
	res.NewSkillName = skill.NewSkillName
	resStats := make(map[string]ResSkillStat)
	for _, s := range skill.Stats {
		resStats[s.StatName] = ResSkillStat{s.StatRoll}
	}
	res.Stats = resStats
	return res
}
