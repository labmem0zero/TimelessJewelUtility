package main

import (
	"strings"
)

func (tj *TimelessJewel) GetSkillType(skill PassiveSkill) uint32 {
	if skill.IsNotable {
		return 3
	}
	if skill.IsKeystone {
		return 4
	}
	if skill.IsJewelSocket {
		return 5
	}
	if len(skill.Stats) == 1 {
		var bitPosition = (skill.Stats[0] + 1) - 574
		if (bitPosition < 7) && (0x49&(1<<int(bitPosition)) != 0) {
			return 1
		}
	}
	return 2
}

func (tjs *TimelessJewels) PrepReplaceableSkills() {
	for _, j := range AltSkills {
		if strings.Contains(j.Id, NodeRaces[race]) && !strings.Contains(j.Id, "keystone") {
			tjs.ReplaceableSkills = append(tjs.ReplaceableSkills, j)
		}
	}
}

func (tjs *TimelessJewels) PrepAdditionSkills() {
	for _, j := range AltAdditions {
		if strings.Contains(j.Id, NodeRaces[race]) && !strings.Contains(j.Id, "keystone") {
			tjs.AdditionableeSkills = append(tjs.AdditionableeSkills, j)
		}
	}
}

func (tj *TimelessJewel) FindSkillType(a AlternatePassiveSkill, skillType uint32) bool {
	for _, t := range a.PassiveType {
		if t == skillType {
			return true
		}
	}
	return false
}

func (tj *TimelessJewel) FindAdditionType(a AlternatePassiveAddition, skillType uint32) bool {
	for _, t := range a.PassiveType {
		if t == skillType {
			return true
		}
	}
	return false
}

func (tj *TimelessJewel) GetApplicableSkills(baseSkill PassiveSkill) []AlternatePassiveSkill {
	passType := tj.GetSkillType(baseSkill)
	var reps []AlternatePassiveSkill
	for _, skill := range fAltSkills {
		if skill.AlternateTreeVersionsKey != race || !tj.FindSkillType(skill, passType) {
			continue
		}
		reps = append(reps, skill)
	}
	return reps
}
func (tj *TimelessJewel) GetApplicableAdditions(baseSkill PassiveSkill) []AlternatePassiveAddition {
	passType := tj.GetSkillType(baseSkill)
	var adds []AlternatePassiveAddition
	for _, addition := range fAltAdditions {
		if addition.AlternateTreeVersionsKey != race || !tj.FindAdditionType(addition, passType) {
			continue
		}
		adds = append(adds, addition)
	}
	return adds
}

func (tj *TimelessJewel) IsPassiveReplaced(sgs SkillGroupSkill) bool {
	if sgs.Skill.IsKeystone {
		return true
	}
	if sgs.Skill.IsNotable {
		if AltTreeVers[race].ReplaceSpawnWeight >= 100 {
			return true
		}
		tj.Rng.InitRandomize([]uint32{sgs.Skill.PassiveSkillGraphId, tj.Seed})
		return tj.Rng.GenerateTwo(0, 100) < AltTreeVers[race].ReplaceSpawnWeight
	}
	if len(sgs.Skill.Stats) == 1 {
		bit := sgs.Skill.Stats[0] + 1 - 574
		if bit <= 6 && (0x49&(1<<bit) != 0) {
			return AltTreeVers[race].IsSmallAttributeReplaced
		}
	}
	return AltTreeVers[race].IsSmallNormalReplaced
}

func (tj *TimelessJewel) ReplaceKeystone(conqId, version uint32, baseskill PassiveSkill) SkillGroupSkill {
	for _, skill := range Jewels.ReplaceableSkills {
		if skill.ConqID == conqId && skill.ConqVersion == version {
			return SkillGroupSkill{OldSkillName: baseskill.Name, NewSkillName: skill.Name}
		}
	}
	return SkillGroupSkill{}
}

func (tj *TimelessJewel) AlternateSkill(skill PassiveSkill) SkillGroupSkill {
	if skill.IsKeystone {
		return tj.ReplaceKeystone(tj.ConqID, tj.ConqVer, skill)
	}

	tj.Rng.InitRandomize([]uint32{skill.PassiveSkillGraphId, tj.Seed})

	tj.ApplicableSkills = tj.GetApplicableSkills(skill)
	var rolledSkill AlternatePassiveSkill
	currentSpawnWeight := uint32(0)
	if skill.IsNotable {
		tj.Rng.GenerateTwo(0, 100)
	}
	for _, a := range tj.ApplicableSkills {
		currentSpawnWeight += a.SpawnWeight
		roll := tj.Rng.GenerateOne(currentSpawnWeight)
		if roll < a.SpawnWeight {
			rolledSkill = a
		}
	}
	newSgs := SkillGroupSkill{}
	newSgs.OldSkillName = skill.Name
	newSgs.NewSkillName = rolledSkill.Name
	newSgs.AddAlts(rolledSkill)

	for i, stat := range newSgs.Stats {
		roll := stat.MinRoll
		if stat.MaxRoll > stat.MinRoll {
			roll = tj.Rng.GenerateTwo(stat.MinRoll, stat.MaxRoll)
		}
		newSgs.Stats[i].StatRoll = roll
	}

	if rolledSkill.RandomMin == 0 && rolledSkill.RandomMax == 0 {
		return newSgs
	}

	minAdds := AltTreeVers[tj.RaceID].MinAddsCount + rolledSkill.RandomMin
	maxAdds := AltTreeVers[tj.RaceID].MaxAddsCount + rolledSkill.RandomMax
	countAdds := minAdds
	if maxAdds > minAdds {
		countAdds = tj.Rng.GenerateTwo(minAdds, maxAdds)
	}
	for i := 0; i < int(countAdds); i++ {
		var addition AlternatePassiveAddition
		for addition.Rid == 0 {
			addition = tj.RollAddition()
			if addition.Id == "zero" {
				break
			}
		}
		if addition.Id == "zero" {
			continue
		}
		roll := addition.Stat1MinRoll
		if addition.Stat1MaxRoll > addition.Stat1MinRoll {
			roll = tj.Rng.GenerateTwo(addition.Stat1MinRoll, addition.Stat1MaxRoll)
		}
		newSgs.AddAdd(addition, roll)
	}
	return newSgs
}

func (tj *TimelessJewel) MakeAdditions(skill PassiveSkill) SkillGroupSkill {
	tj.Rng.InitRandomize([]uint32{skill.PassiveSkillGraphId, tj.Seed})
	tj.ApplicableAdditions = tj.GetApplicableAdditions(skill)
	if skill.IsNotable {
		tj.Rng.GenerateTwo(0, 100)
	}
	minAdds := AltTreeVers[tj.RaceID].MinAddsCount
	maxAdds := AltTreeVers[tj.RaceID].MaxAddsCount
	countAdds := minAdds
	if maxAdds > minAdds {
		countAdds = tj.Rng.GenerateTwo(minAdds, maxAdds)
	}
	newSgs := SkillGroupSkill{}
	newSgs.Skill = skill
	newSgs.OldSkillName = skill.Name
	newSgs.NewSkillName = skill.Name
	for i := 0; i < int(countAdds); i++ {
		var addition AlternatePassiveAddition
		for addition.Rid == 0 {
			addition = tj.RollAddition()
			if addition.Id == "zero" {
				break
			}
		}
		if addition.Id == "zero" {
			continue
		}
		roll := addition.Stat1MinRoll
		if addition.Stat1MaxRoll > addition.Stat1MinRoll {
			roll = tj.Rng.GenerateTwo(addition.Stat1MinRoll, addition.Stat1MaxRoll)
		}
		newSgs.AddAdd(addition, roll)
	}
	return newSgs
}

func (tj *TimelessJewel) RollAddition() AlternatePassiveAddition {
	var sum uint32
	for _, v := range tj.ApplicableAdditions {
		sum += v.SpawnWeight
	}
	if sum == 0 {
		return AlternatePassiveAddition{Id: "zero"}
	}
	roll := tj.Rng.GenerateOne(sum)
	for _, v := range tj.ApplicableAdditions {
		if v.SpawnWeight > roll {
			return v
		}
		roll -= v.SpawnWeight
	}
	return AlternatePassiveAddition{}
}

func (tj *TimelessJewel) ProcessSkill(sgs SkillGroupSkill) SkillGroupSkill {
	replaced := tj.IsPassiveReplaced(sgs)
	if replaced {
		return tj.AlternateSkill(sgs.Skill)
	}
	return tj.MakeAdditions(sgs.Skill)
}
