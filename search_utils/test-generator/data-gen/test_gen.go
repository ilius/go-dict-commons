package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/ilius/go-dict-commons/search_utils"
)

var maxStringLength = 13

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

var stringCountByLen = map[int]int{
	1:  20,
	2:  30,
	3:  40,
	4:  50,
	5:  60,
	6:  60,
	7:  50,
	8:  35,
	9:  20,
	10: 10,
	11: 8,
	12: 6,
	13: 4,
}

var testCountByScoreRange = map[[2]int]int{
	{1, 0}:  10,
	{2, 0}:  20,
	{3, 0}:  30,
	{4, 0}:  40,
	{5, 0}:  50,
	{6, 0}:  50,
	{7, 0}:  40,
	{8, 0}:  30,
	{9, 0}:  20,
	{10, 0}: 20,
	{11, 0}: 20,
	{12, 0}: 20,
	{13, 0}: 20,

	{2, 100}: 50,

	{3, 60}:  100,
	{3, 130}: 100,

	{4, 50}:  200,
	{4, 100}: 200,
	{4, 150}: 200,

	{5, 40}:  300,
	{5, 80}:  300,
	{5, 120}: 300,
	{5, 160}: 300,

	{6, 30}:  300,
	{6, 60}:  300,
	{6, 100}: 300,
	{6, 130}: 300,
	{6, 160}: 300,

	{7, 20}:  200,
	{7, 50}:  200,
	{7, 80}:  200,
	{7, 110}: 200,
	{7, 140}: 200,
	{7, 170}: 200,

	{8, 20}:  150,
	{8, 50}:  150,
	{8, 70}:  150,
	{8, 100}: 150,
	{8, 120}: 150,
	{8, 150}: 150,
	{8, 170}: 150,

	{9, 20}:  100,
	{9, 40}:  100,
	{9, 60}:  100,
	{9, 80}:  100,
	{9, 110}: 100,
	{9, 130}: 100,
	{9, 150}: 100,
	{9, 170}: 100,

	{10, 20}:  50,
	{10, 40}:  50,
	{10, 60}:  50,
	{10, 80}:  50,
	{10, 100}: 50,
	{10, 120}: 50,
	{10, 140}: 50,
	{10, 160}: 50,
	{10, 180}: 50,

	{11, 10}:  40,
	{11, 30}:  40,
	{11, 50}:  40,
	{11, 70}:  40,
	{11, 90}:  40,
	{11, 100}: 40,
	{11, 120}: 40,
	{11, 140}: 40,
	{11, 160}: 40,
	{11, 180}: 40,

	{12, 10}:  40,
	{12, 30}:  40,
	{12, 50}:  40,
	{12, 60}:  40,
	{12, 80}:  40,
	{12, 100}: 40,
	{12, 110}: 40,
	{12, 130}: 40,
	{12, 150}: 40,
	{12, 160}: 40,
	{12, 180}: 40,

	{13, 10}:  40,
	{13, 30}:  40,
	{13, 40}:  40,
	{13, 60}:  40,
	{13, 70}:  40,
	{13, 90}:  40,
	{13, 100}: 40,
	{13, 120}: 40,
	{13, 130}: 40,
	{13, 150}: 40,
	{13, 160}: 40,
	{13, 180}: 40,
}

func init() {
	for length := 1; length <= maxStringLength; length++ {
		for dist := length; dist > 0; dist-- {
			score := 200 * (length - dist) / length
			scoreRange := int(score / 10 * 10)
			_, ok := testCountByScoreRange[[2]int{length, scoreRange}]
			if !ok {
				fmt.Printf("{%d, %d}: 40,\n", length, scoreRange)
			}
		}
	}
	// os.Exit(0)
	// fmt.Println("--------------------------------------------")
}

func genRandomRunes(length int, letterCount int) []rune {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(letterCount)]
	}
	return b
}

func genRandomRunesList(length int, letterCount int) [][]rune {
	if letterCount > len(letterRunes) {
		letterCount = len(letterRunes)
	}
	count := stringCountByLen[length] * 20
	strs := make([][]rune, count)
	for index := 0; index < count; index++ {
		strs[index] = genRandomRunes(length, letterCount)
	}
	return strs
}

// func randomTwoVarRange(n1 int, n2 int) chan [2]int {
// 	ch := make(chan [2]int)
// 	items := [][2]int{}
// 	for i := 0; i < n1; i++ {
// 		for j := 0; j < n2; j++ {
// 			items = append(items, [2]int{i, j})
// 		}
// 	}
// 	rand.Shuffle(len(items), func(i, j int) {
// 		items[i], items[j] = items[j], items[i]
// 	})

// 	go func() {
// 		defer close(ch)
// 		for _, item := range items {
// 			ch <- item
// 		}
// 	}()
// 	return ch
// }

func main() {
	buff := make([]uint16, maxStringLength+1)
	tests := [][2][]rune{}
	for a_length := 1; a_length <= maxStringLength; a_length++ {
		a_list := genRandomRunesList(a_length, 15+a_length*3)
		for b_length := a_length; b_length <= maxStringLength; b_length++ {
			b_list := genRandomRunesList(b_length, 10+b_length*3)
			for _, a := range a_list {
				for _, b := range b_list {
					if len(a) == len(b) && string(a) == string(b) {
						continue
					}
					tests = append(tests, [2][]rune{a, b})
				}
			}
		}
	}
	rand.Shuffle(len(tests), func(i, j int) {
		tests[i], tests[j] = tests[j], tests[i]
	})
	testCounts := map[int]int{}
	output := []map[string]any{}
	for _, item := range tests {
		a := item[0]
		b := item[1]
		score := search_utils.SimilaritySlow(a, b, buff, 0)
		if score == 200 {
			continue
		}
		scoreRange := int(score / 10 * 10)
		scoreRangeMax, ok := testCountByScoreRange[[2]int{len(b), scoreRange}]
		if !ok {
			log.Fatalf("scoreRange=%v, b_length=%v", scoreRange, len(b))
		}
		if testCounts[scoreRange] >= scoreRangeMax {
			continue
		}
		testCounts[scoreRange]++
		output = append(output, map[string]any{
			"a": string(a),
			"b": string(b),
			"s": score,
		})
	}
	outputJson, err := json.MarshalIndent(output, "", "\t")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("tests.json", outputJson, 0o644)
	if err != nil {
		panic(err)
	}
}
