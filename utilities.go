package main

import (
	"fmt"
	"sort"
	"strings"
)

var AltCounter int
var AddCounter int

func IsSkillPresented(needed []StatParam, skillName string) string {
	for _, s := range needed {
		if s.id == skillName {
			return skillName
		}
	}
	return ""
}

func IsStatPresented(needed []StatParam, stat SkillStat) string {
	for _, s := range needed {
		if s.id == stat.StatName {
			return stat.StatName
		}
	}
	return ""
}

func PrintAltSkills() {
	fmt.Println("List of possible transformed skills:")
	AltCounter = 0
	for _, j := range fAltSkills {
		if strings.Contains(j.Id, NodeRaces[race]) && !strings.Contains(j.Id, "keystone") {
			AltCounter++
			fmt.Printf("ID: %v\tName:\t%s\n", j.Rid, j.Name)
		}
	}
	if AltCounter < 1 {
		fmt.Println("No replacing skills for this jewel found")
	}
}

func PrintAddSkills() {
	fmt.Println("List of possible additional skills:")
	AddCounter = 0
	for _, j := range fAltAdditions {
		if strings.Contains(j.Id, NodeRaces[race]) && !strings.Contains(j.Id, "keystone") {
			AddCounter++
			fmt.Printf("ID: %v\tName:\t%s\n", j.Rid, replacer.Replace(j.Id))
		}
	}
	if AddCounter < 1 {
		fmt.Println("No additional skills for this jewel found")
	}
}

func SortSkills(quants map[string]byte) []string {
	count := make(map[byte][]string)
	var keys []byte
	for k, v := range quants {
		if _, ok := count[v]; !ok {
			keys = append(keys, v)
		}
		count[v] = append(count[v], k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	var res []string
	for _, k := range keys {
		res = append(res, count[k]...)
	}
	return res
}
