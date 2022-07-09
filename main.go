package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var race, conq, altrid, addrid, Seed uint32
var NeededStats []StatParam
var NeededSkills []StatParam
var JewelNames []string
var ConqNames []string

type StatParam struct {
	id    string
	count byte
}

func ReadInputs() {
	fmt.Println("List of jewels:")
	for i, n := range JewelNames {
		fmt.Printf("%v.\t%s\n", i+1, n)
	}
	fmt.Println("Input jewel number")
	fmt.Scanln(&race)
	fmt.Println("List of conquerors:")
	for i, j := range JewelsAndConquerors[race] {
		fmt.Printf("%v.\t%s\n", i, j.ConqName)
	}
	fmt.Println("Input Conqueror number")
	fmt.Scanln(&conq)
	Jewels.PrepReplaceableSkills()
	Jewels.PrepAdditionSkills()
	PrintAltSkills()
	var transCount int
	if AltCounter > 0 {
		fmt.Println("Input needed alternative skill count(0 if no need)")
		fmt.Scanln(&transCount)
		if transCount > 0 {
			for i := 0; i < transCount; i++ {
				fmt.Println("Input needed alternative skill id")
				var transId uint32
				fmt.Scanln(&transId)
				fmt.Println("Input needed alternative skill amount inside skill group(around jewel socket)")
				var transAmount byte
				fmt.Scanln(&transAmount)
				neededSkill := StatParam{AltSkills[transId].Name, transAmount}
				NeededSkills = append(NeededSkills, neededSkill)
			}
		}
		fmt.Println("Needed skills")
		fmt.Println(NeededSkills)
	}

	PrintAddSkills()
	if AddCounter > 0 {
		var addsCount int
		fmt.Println("Input needed additional skill count(0 if no need)")
		fmt.Scanln(&addsCount)
		if addsCount > 0 {
			for i := 0; i < addsCount; i++ {
				fmt.Println("Input needed additional skill id")
				var addId uint32
				fmt.Scanln(&addId)
				fmt.Println("Input needed additional skill amount inside skill group(around jewel socket)")
				var addAmount byte
				fmt.Scanln(&addAmount)
				neededStat := StatParam{AltAdditions[addId].Id, addAmount}
				NeededStats = append(NeededStats, neededStat)
			}
		}
	}
}

func StartRollingOverAllSeeds() {
	ReadInputs()
	fmt.Println("Started rolling over all seeds. Wait for message that it is ended")
	Jewels.AddJewels()
	res := Jewels.RollOverSeeds()
	f, _ := os.Create("result.json")
	defer f.Close()
	d, _ := json.MarshalIndent(res, "", "\t")
	f.Write(d)

	f2, _ := os.Create("result_readable.txt")
	defer f2.Close()
	for k, v := range res {
		f2.Write([]byte(fmt.Sprintf("Seed:%v\n", k)))
		for k2, v2 := range v {
			f2.Write([]byte(fmt.Sprintf("\tJewel socket:%v\n", k2)))
			for s, _ := range v2 {
				f2.Write([]byte(fmt.Sprintf("\t\tSkill name:%v\n", s)))
			}
		}
		f2.Write([]byte(fmt.Sprintf("\n")))
	}
}

func StartRollingOverExactSeed() {
	fmt.Println("List of jewels:")
	for i, n := range JewelNames {
		fmt.Printf("%v.\t%s\n", i+1, n)
	}
	fmt.Println("Input jewel number")
	fmt.Scanln(&race)
	fmt.Println("Input seed number")
	fmt.Scanln(&Seed)
	if race == 5 {
		Seed /= 5
	}
	Jewels.FindAllPassives(race, Seed)
}

func main() {
	LoadFiles()
	InitConqs()
	fmt.Println("Do you want:")
	fmt.Println("1.\tRoll over all seeds looking for matched parameters")
	fmt.Println("2.\tRoll over exact jewel seed looking for all presented stats")
	var choice byte
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		StartRollingOverAllSeeds()
	case 2:
		StartRollingOverExactSeed()
	default:
		fmt.Println("Wrong input. Goodbye!")
	}
	fmt.Println("Ended. Results in files: result.json, result_readable.txt")
	fmt.Scanln(&choice)

}
