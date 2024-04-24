package challenge1

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func FirstChallenge() {
	_hands, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	hands := strings.Split(string(_hands), "\n")

	hands = hands[:len(hands)-1]

	handsOrder := []string{
		"Five of a kind",
		"Four of a kind",
		"Full house",
		"Three of a kind",
		"Two pair",
		"One pair",
		"high card",
	}

	bidAmountMap := make(map[string]int)
	matchWinQty := make(map[string]int)

	for i, v := range hands {
		hb := strings.Split(v, " ")
		ba, err := strconv.Atoi(hb[1])
		if err != nil {
			panic(err)
		}
		bidAmountMap[hb[0]] = ba
		matchWinQty[hb[0]] = 0
		hands[i] = hb[0]
	}

	for i, h := range hands {
		hValue := getHandValue(h)
		hHandStrength := slices.Index(handsOrder, hValue)
		for j, h2 := range hands {
			if i == j {
				continue
			}
			h2Value := getHandValue(h2)
			h2HandStrength := slices.Index(handsOrder, h2Value)

			if hHandStrength == h2HandStrength {
				hWins := sameResultHand(h, h2)
				if hWins {
					matchWinQty[h] += 1
				}
				continue
			}
			if hHandStrength < h2HandStrength {
				matchWinQty[h] += 1
			}
		}
	}

	handsRank := calcRank(matchWinQty)

	totalOfWinnings := calcTotalOfWinnings(handsRank, bidAmountMap)

	fmt.Println(totalOfWinnings)
}

func getHandValue(h string) string {
	lableCountMap := make(map[string]int)

	for _, l := range h {
		lableCountMap[string(l)] += 1
	}

	handValuesMap := map[string]string{
		"5":  "Five of a kind",
		"4":  "Four of a kind",
		"32": "Full house",
		"3":  "Three of a kind",
		"22": "Two pair",
		"2":  "One pair",
		"0":  "high card",
	}

	var handValue string
	var handValuesArr []int

	for _, v := range lableCountMap {
		if v > 1 {
			handValuesArr = append(handValuesArr, v)
		}
	}

	if handValuesArr == nil {
		return handValuesMap["0"]
	}

	sort.Slice(handValuesArr, func(i, j int) bool {
		return handValuesArr[i] > handValuesArr[j]
	})

	for _, v := range handValuesArr {
		handValue += fmt.Sprintf("%d", v)
	}

	return handValuesMap[handValue]
}

func sameResultHand(h, h2 string) bool {
	strenghtOrder := []string{
		"A",
		"K",
		"Q",
		"J",
		"T",
		"9",
		"8",
		"7",
		"6",
		"5",
		"4",
		"3",
		"2",
	}
	for i, v := range h {
		hStrength := slices.Index(strenghtOrder, string(v))
		h2Strength := slices.Index(strenghtOrder, string(h2[i]))

		if hStrength == h2Strength {
			continue
		}

		if hStrength < h2Strength {
			return true
		}

		return false
	}
	return false
}

func calcRank(matchWinQty map[string]int) []map[string]int {
	var handsRank []map[string]int

	for k, v := range matchWinQty {
		handsRank = append(handsRank, map[string]int{k: v})
	}

	sort.Slice(handsRank, func(i, j int) bool {
		var iValue, jValue int
		for _, v := range handsRank[i] {
			iValue = v
		}
		for _, v := range handsRank[j] {
			jValue = v
		}
		return iValue < jValue
	})

	return handsRank
}

func calcTotalOfWinnings(handsRank []map[string]int, bidAmountMap map[string]int) int {
	var totalOfWinnings int

	for i, v := range handsRank {
		for k := range v {
			totalOfWinnings += bidAmountMap[k] * (i + 1)
		}
	}

	return totalOfWinnings
}
