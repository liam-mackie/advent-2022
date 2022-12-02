package one

import (
	"github.com/liam-mackie/advent2022/pkg/Utilities"
	"github.com/samber/lo"
	"testing"
)

func TestPartOne(t *testing.T) {
	currentDirectory := Utilities.GetDayDirectory()
	data := Utilities.GetAllTestData(currentDirectory)
	elves := getElvesSortedByCalories(data)
	biggestElf := getElfWithMostCalories(elves)
	wantedPosition := 4
	wantedCalories := int64(24000)

	if biggestElf.Position != wantedPosition && biggestElf.TotalCalories != wantedCalories {
		t.Fatalf(`Elf with the most calories is in position %v, and carries %v calories, wanted position %v and %v calories`, biggestElf.Position, biggestElf.TotalCalories, wantedPosition, wantedCalories)
	}
}

func TestPartTwo(t *testing.T) {
	currentDirectory := Utilities.GetDayDirectory()
	data := Utilities.GetAllTestData(currentDirectory)
	elves := getElvesSortedByCalories(data)
	elvesWithMostCalories, totalCalories := getTotalCaloriesOfTopElves(elves, 3)
	wantedElves := []int{4, 3, 5}
	wantedTotalCalories := int64(45000)
	for _, currentElf := range elvesWithMostCalories {
		if !lo.Contains(wantedElves, currentElf.Position) {
			t.Fatalf(`Elf position %v not expected in top elves. Expected %v.`, currentElf.Position, wantedElves)
		}
	}
	if totalCalories != wantedTotalCalories {
		t.Fatalf(`Total calories %v not the expected answer. Expected %v.`, totalCalories, wantedTotalCalories)
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		currentDirectory := Utilities.GetDayDirectory()
		data := Utilities.GetAllRealData(currentDirectory)
		getElvesSortedByCalories(data)
	}
}
