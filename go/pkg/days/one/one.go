package one

import (
	"fmt"
	"github.com/liam-mackie/advent2022/pkg/Utilities"
	"sort"
	"strconv"
)

type elf struct {
	Position      int
	TotalCalories int64
	Items         []heldItem
}

type heldItem struct {
	Calories int64
}

func Run() {
	currentDirectory := Utilities.GetDayDirectory()
	data := Utilities.GetAllRealData(currentDirectory)
	elves := getElvesSortedByCalories(data)
	biggestElf := getElfWithMostCalories(elves)
	fmt.Printf("The elf with the most calories is in position %v and has %v calories.\n", biggestElf.Position, biggestElf.TotalCalories)
	_, totalCalories := getTotalCaloriesOfTopElves(elves, 3)
	fmt.Printf("The total calories of the 3 top elves is %v.\n", totalCalories)
}

func getElvesSortedByCalories(data string) []elf {
	rawElves := Utilities.SplitByEmptyNewline(data)
	elves := make([]elf, 0)
	for i, rawElf := range rawElves {
		currentElf := new(elf)
		heldItems := make([]heldItem, 0)
		items := Utilities.SplitByNewline(rawElf)
		for _, item := range items {
			itemCalories, err := strconv.ParseInt(item, 10, 64)
			heldItems = append(heldItems, heldItem{itemCalories})
			currentElf.TotalCalories = currentElf.TotalCalories + itemCalories
			Utilities.CheckError(err)
		}
		currentElf.Items = heldItems
		currentElf.Position = i + 1
		elves = append(elves, *currentElf)
	}
	sortElvesByCalories(elves)
	return elves
}

func getTotalCaloriesOfTopElves(elves []elf, count int) ([]elf, int64) {
	topElves := elves[:count]
	var totalCalories int64
	for _, currentElf := range topElves {
		totalCalories = totalCalories + currentElf.TotalCalories
	}
	return topElves, totalCalories
}
func getElfWithMostCalories(elves []elf) elf {
	return elves[0]
}

func sortElvesByCalories(elves []elf) []elf {
	sort.SliceStable(elves, func(p, q int) bool {
		return elves[p].TotalCalories > elves[q].TotalCalories
	})
	return elves
}
