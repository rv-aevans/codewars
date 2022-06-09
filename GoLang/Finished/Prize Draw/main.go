package main

import (
	"sort"
	"strings"
)

// Each letter has a rank A/a are Rank 1, B/b are 2, etc

// Length of the first name is added to the sum of these ranks - som

// An array of random weight is linked to the firstnames and each sum is multipied by it's corresponding weight
// to get what they call a winning number

// Now one can sort the firstnames in decreasing order of the winning numbers. When two people have the same winning
// number sort them alphabetically by their firstnames.

// st a string of names
// we an array of weights
// n rank

// Return the first name of the participant who rank is n (ranks are numbered from 1)

// if st empty return "No participants"

// If n is greater than the number of participants return "Not enough participants"

var (
	st        = "Artie,Henry,Arthur,Mary"
	we        = []int{2, 2, 2, 2}
	n         = 1
	letterMap = make(map[string]int)
	alphabet  = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

type Pair struct {
	Name  string
	Value int
}

func main() {
	solution(st, we, n)
}

func solution(st string, we []int, n int) string {

	if st == "" {
		return "No participants."
	}

	s := strings.Split(st, ",")
	if len(s) < n {
		return "Not enough participants."
	}

	PairList := assignNameScores(st, we)

	return PairList[n-1].Name
}

func assignNameScores(names string, multiplier []int) []Pair {
	letterMap = createLetterValueMap()

	splitNames := strings.Split(names, ",")

	namesToMultiplier := createNameToMultiplierMap(splitNames, multiplier)

	var PairList []Pair

	for _, name := range splitNames {
		nameValue := len(name)
		lowerName := strings.ToLower(name)
		for m := 0; m < len(name); m++ {
			nameValue = nameValue + letterMap[string(lowerName[m])]
		}
		nameValueWithMultiplier := nameValue * namesToMultiplier[name]
		pair := Pair{
			Name:  name,
			Value: nameValueWithMultiplier,
		}
		PairList = append(PairList, pair)
	}

	sort.Slice(PairList, func(i, j int) bool {
		return PairList[i].Value > PairList[j].Value
	})

	sort.Slice(PairList, func(i, j int) bool {
		if PairList[i].Value == PairList[j].Value {
			return PairList[i].Name < PairList[j].Name
		}
		return false
	})

	return PairList
}

func createLetterValueMap() map[string]int {
	for i := 0; i <= 25; i++ {
		letterMap[alphabet[i]] = i + 1
	}
	return letterMap
}

func createNameToMultiplierMap(names []string, multiplier []int) map[string]int {
	namesToMultiplier := make(map[string]int)

	for i := 0; i < len(names); i++ {
		namesToMultiplier[names[i]] = multiplier[i]
	}

	return namesToMultiplier
}
