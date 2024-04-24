package challenge1

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type BestCardToTurn struct {
	key   string
	value int
}

func SecondChallenge() {
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
	strenghtOrder := []string{
		"A",
		"K",
		"Q",
		"T",
		"9",
		"8",
		"7",
		"6",
		"5",
		"4",
		"3",
		"2",
		"J",
	}

	bidAmountMap := make(map[string]int)
	matchWinQty := make(map[string]int)

	for i, v := range hands {
		hb := strings.Split(v, " ")
		newH := hb[0]
		if strings.Contains(hands[i], "J") {
			newH = treatJockerCards(hb[0], strenghtOrder)
		}
		ba, err := strconv.Atoi(hb[1])
		if err != nil {
			panic(err)
		}
		bidAmountMap[newH] = ba
		matchWinQty[newH] = 0
		hands[i] = newH
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
				hWins := sameResultHand(h, h2, strenghtOrder)
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

func treatJockerCards(h string, strenghtOrder []string) string {
	newH := h
	lableCountMap := make(map[string]int)

	if strings.Count(newH, "J") == 5 {
		return newH
	}

	for _, l := range h {
		lableCountMap[string(l)] += 1
	}

	var bestCardToTurn BestCardToTurn
	var sameValuesInArr []int

	for k, v := range lableCountMap {
		if k == "J" {
			continue
		}
		sameValuesInArr = append(sameValuesInArr, v)
		if v > bestCardToTurn.value {
			bestCardToTurn = BestCardToTurn{
				key:   k,
				value: v,
			}
		}
	}

	condition := func(num int) bool {
		return num == bestCardToTurn.value
	}

	sameValuesInArr = filter(sameValuesInArr, condition)

	if len(sameValuesInArr) > 1 {
		var keys []string

		for k, v := range lableCountMap {
			if v == sameValuesInArr[0] {
				keys = append(keys, k)
			}
		}

		greatLabel := BestCardToTurn{
			value: 1<<63 - 1,
		}

		for _, v := range keys {
			index := slices.Index(strenghtOrder, string(v))
			if index < greatLabel.value {
				greatLabel = BestCardToTurn{
					key:   v,
					value: index,
				}
			}
		}
		// t := bestCardToTurn
		bestCardToTurn = greatLabel
		// fmt.Println("h: ", newH, "old: ", t, "new: ", bestCardToTurn)
	}

	pattern := fmt.Sprintf(`%s`, "J")
	reg, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	newH = reg.ReplaceAllString(newH, bestCardToTurn.key)
	fmt.Println("new: ", newH, "old: ", h)
	// FIX: NÃO SOBREPOR A MÃO ATUAL, SOMENTE MUDAR O VALOR NA ATRIBUIÇÃO DA MÃO
	// POR EXEMPLO, O J PODE FAZER ELA FORMAR UM FULL HOUSE, MAS DEVE CONTINUAR SENDO UM J
	// PORQUE QUANDO DUAS MÃOS OBTIVEREM O MESMO TIPO O J DEVE SER CONTADO COMO J E NÃO COMO
	// UMA CARTA MAIS FORTE

	return newH
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

func sameResultHand(h, h2 string, strenghtOrder []string) bool {
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

func filter(numbers []int, condition func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if condition(num) {
			result = append(result, num)
		}
	}
	return result
}
