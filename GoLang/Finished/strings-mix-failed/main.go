package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s1, s2 = "Are they here", "yes, they are here"

	fmt.Println(Mix(s1, s2))
}

func Mix(s1, s2 string) string {
	alphabase := "abcdefghijklmnopqrstuvwxyz"
	result := []string{}
	for _, c := range alphabase {
		nb_s1 := strings.Count(s1, string(c))
		nb_s2 := strings.Count(s2, string(c))
		if nb_s1 > 1 || nb_s2 > 1 {
			if nb_s1 == nb_s2 {
				result = append(result, "=:"+strings.Repeat(string(c), nb_s1))
			}
			if nb_s1 > nb_s2 {
				result = append(result, "1:"+strings.Repeat(string(c), nb_s1))
			}
			if nb_s1 < nb_s2 {
				result = append(result, "2:"+strings.Repeat(string(c), nb_s2))
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) == len(result[j]) {
			return result[i] < result[j]
		}
		return len(result[i]) > len(result[j])
	})
	return strings.Join(result, "/")
}

// func Mix(s1, s2 string) string {
// 	splitOne, splitTwo := strings.Split(s1, ""), strings.Split(s2, "")

// 	mOne, mTwo := make(map[string]int), make(map[string]int)

// 	for _, char := range splitOne {
// 		lowerChar := strings.ToLower(char)
// 		if checkForLetter([]rune(lowerChar)) {
// 			mOne[lowerChar] = handleCharacter(mOne, lowerChar)
// 		}
// 	}
// 	for _, char := range splitTwo {
// 		lowerChar := strings.ToLower(char)
// 		if checkForLetter([]rune(lowerChar)) {
// 			mTwo[lowerChar] = handleCharacter(mTwo, lowerChar)
// 		}
// 	}

// 	var keys []string

// 	for k := range mOne {
// 		keys = append(keys, k)
// 	}

// 	for k := range mTwo {
// 		keys = append(keys, k)
// 	}

// 	allKeys := make(map[string]bool)
// 	dedupedKeys := []string{}
// 	for _, item := range keys {
// 		if _, value := allKeys[item]; !value {
// 			allKeys[item] = true
// 			dedupedKeys = append(dedupedKeys, item)
// 		}
// 	}

// 	parentMap := make(map[string]string)

// 	for _, char := range dedupedKeys {
// 		if _, okOne := mOne[char]; okOne {
// 			if _, okTwo := mTwo[char]; okTwo {
// 				parentMap[char+"-both"] = char
// 			} else {
// 				parentMap[char+"-one"] = char
// 			}
// 		} else {
// 			parentMap[char+"-two"] = char
// 		}
// 	}

// 	finalMap := make(map[string]int)

// 	for key := range parentMap {
// 		if strings.Contains(key, "both") {
// 			strippedKey := strings.TrimSuffix(key, "-both")
// 			if mOne[parentMap[key]] == mTwo[parentMap[key]] {
// 				if mOne[parentMap[key]] > 1 {
// 					finalMap[key] = mOne[parentMap[key]]
// 				}
// 			} else {
// 				if mOne[parentMap[key]] > mTwo[parentMap[key]] {
// 					if mOne[parentMap[key]] > 1 {
// 						finalMap[strippedKey+"-one"] = mOne[parentMap[key]]
// 					}
// 				} else {
// 					if mTwo[parentMap[key]] > 1 {
// 						finalMap[strippedKey+"-two"] = mTwo[parentMap[key]]
// 					}
// 				}
// 			}
// 		} else {
// 			if strings.Contains(key, "-one") {
// 				if mOne[parentMap[key]] > 1 {
// 					finalMap[key] = mOne[parentMap[key]]
// 				}
// 			} else {
// 				if mTwo[parentMap[key]] > 1 {
// 					finalMap[key] = mTwo[parentMap[key]]
// 				}
// 			}
// 		}
// 	}

// 	indexes := make([]string, 0, len(finalMap))

// 	for k := range finalMap {
// 		indexes = append(indexes, k)
// 	}

// 	sort.Slice(indexes, func(i, j int) bool {
// 		return indexes[i] > indexes[j]
// 	})

// 	fmt.Println(indexes)

// 	sort.SliceStable(indexes, func(i, j int) bool {
// 		return finalMap[indexes[i]] > finalMap[indexes[j]]
// 	})

// 	fmt.Println(indexes)

// 	var resString string

// 	var curLength int
// 	var curSlice int

// 	for _, index := range indexes {
// 		fmt.Println(curLength, index)
// 		if strings.Contains(index, "-both") {
// 			char := strings.TrimSuffix(index, "-both")
// 			var addString string
// 			for i := 0; i < mOne[char]; i++ {
// 				addString = addString + char
// 			}

// 			if curSlice == 1 || curSlice == 2 {
// 				resString = resString + "=:" + addString + "/"
// 			}
// 			curLength = mOne[char]
// 		}
// 		if strings.Contains(index, "-one") {
// 			char := strings.TrimSuffix(index, "-one")
// 			var addString string
// 			for i := 0; i < mOne[char]; i++ {
// 				addString = addString + char
// 			}
// 			fmt.Println("Length of One: ", mOne[char])
// 			if mOne[char] == curLength {
// 				resString = resString + "=:" + addString

// 			} else {
// 				resString = resString + "1:" + addString + "/"
// 			}
// 			curLength = mOne[char]
// 			curSlice = 1
// 		}
// 		if strings.Contains(index, "-two") {
// 			char := strings.TrimSuffix(index, "-two")
// 			var addString string
// 			for i := 0; i < mTwo[char]; i++ {
// 				addString = addString + char
// 			}
// 			fmt.Println("Length of Two: ", mTwo[char])
// 			if mTwo[char] == curLength {
// 				resString = resString + "=:" + addString
// 			} else {
// 				resString = resString + "2:" + addString + "/"
// 			}
// 			curLength = mTwo[char]
// 			curSlice = 2
// 		}
// 	}

// 	return strings.TrimSuffix(resString, "/")
// }

// func checkForLetter(charVar []rune) bool {
// 	if charVar[0] >= 'a' && charVar[0] <= 'z' {
// 		return true
// 	}
// 	return false
// }

// func handleCharacter(curMap map[string]int, char string) int {
// 	if _, ok := curMap[char]; ok {
// 		return curMap[char] + 1
// 	} else {
// 		return 1
// 	}
// }
