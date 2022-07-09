package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

var Stats map[uint32]Stat
var OriginalSkills map[uint32]PassiveSkill
var AltTreeVers map[uint32]AlternateTreeVersion
var AltSkills map[uint32]AlternatePassiveSkill
var AltAdditions map[uint32]AlternatePassiveAddition
var SkillGroups map[string]SkillGroup

var fStats []Stat
var fOriginalSkills []PassiveSkill
var fAltTreeVers []AlternateTreeVersion
var fAltSkills []AlternatePassiveSkill
var fAltAdditions []AlternatePassiveAddition
var fSkillGroups []SkillGroup

func LoadFiles() {
	Stats = make(map[uint32]Stat)
	f, err := os.Open(statsFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.NewDecoder(f).Decode(&fStats)
	if err != nil {
		log.Fatalln(err)
	}
	for _, s := range fStats {
		Stats[s.Rid] = s
	}

	OriginalSkills = make(map[uint32]PassiveSkill)
	f, err = os.Open(passivesFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.NewDecoder(f).Decode(&fOriginalSkills)
	if err != nil {
		log.Fatalln(err)
	}
	for _, s := range fOriginalSkills {
		OriginalSkills[s.Rid] = s
	}

	AltTreeVers = make(map[uint32]AlternateTreeVersion)
	f, err = os.Open(altTreeVerFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.NewDecoder(f).Decode(&fAltTreeVers)
	if err != nil {
		log.Fatalln(err)
	}
	for _, s := range fAltTreeVers {
		AltTreeVers[s.Rid] = s
	}

	AltSkills = make(map[uint32]AlternatePassiveSkill)
	f, err = os.Open(altPassivesFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.NewDecoder(f).Decode(&fAltSkills)
	if err != nil {
		log.Fatalln(err)
	}
	for _, s := range fAltSkills {
		AltSkills[s.Rid] = s
	}

	AltAdditions = make(map[uint32]AlternatePassiveAddition)
	f, _ = os.Open(altAdditionsFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.NewDecoder(f).Decode(&fAltAdditions)
	if err != nil {
		log.Fatalln(err)
	}
	for _, s := range fAltAdditions {
		AltAdditions[s.Rid] = s
	}

	PrepSkillGroups()
}

func FindSkill(name string) uint32 {
	for _, s := range OriginalSkills {
		if s.Name == name {
			return s.Rid
		}
	}
	return 0
}

func PrepSkillGroups() {
	SkillGroups = make(map[string]SkillGroup)
	f, _ := os.Open(skillgroups)
	scanner := bufio.NewScanner(f)
	currentGroupName := ""
	currentGroup := SkillGroup{}
	currentGroupSkills := map[string]SkillGroupSkill{}
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			continue
		}
		if s[0] == []byte("_")[0] {
			if len(currentGroupName) != 0 && len(currentGroupSkills) != 0 {
				currentGroup.Name = currentGroupName
				currentGroup.Skills = currentGroupSkills
				SkillGroups[currentGroupName] = currentGroup
			}
			currentGroupName = s[1:]
			currentGroup = SkillGroup{}
			currentGroupSkills = make(map[string]SkillGroupSkill)
			continue
		}
		skill := SkillGroupSkill{OldSkillName: s, Rid: FindSkill(s), Skill: OriginalSkills[FindSkill(s)]}
		currentGroupSkills[skill.OldSkillName] = skill
	}
	if len(currentGroupName) != 0 && len(currentGroupSkills) != 0 {
		currentGroup.Name = currentGroupName
		currentGroup.Skills = currentGroupSkills
		SkillGroups[currentGroupName] = currentGroup
	}
	/*fmt.Println("Skills to be looked through:")
	for _, v := range SkillGroups {
		for k, v := range v.Skills {
			fmt.Println(k, ":", v.Skill.Name)
		}
	}*/
}
