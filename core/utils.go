package core

import (
	"sort"
	"strconv"
)

var thousands = map[int]string{
	1: "t",
	2: "tt",
	3: "ttt",
	4: "tttt",
}

var tens = map[int]string{
	1: "d",
	2: "s",
}

type NumRank struct {
	Rate string
	Num  int
}

func PartNum(num int) []NumRank {
	arrText := make([]int, 0)
	arrNumRate := make([]NumRank, 0)
	arrText = append(arrText, Prepare1000(num)...)

	for th, part := range arrText {
		for ts, n := range Prepare10(part) {
			if ts == 0 {
				arrNumRate = append(arrNumRate, NumRank{Num: n, Rate: thousands[th]})
			} else {
				arrNumRate = append(arrNumRate, NumRank{Num: n, Rate: tens[ts]})
			}
		}
	}
	sort.Slice(arrNumRate, func(i, j int) bool {
		return true
	})

	d := ""
	arrNumRatePrepare := make([]NumRank, 0)

	// Числа от 11 - 19 приводим к норм виду
	for _, n := range arrNumRate {
		if n.Num == 1 && n.Rate == "d" {
			d += strconv.Itoa(n.Num)
			continue
		}

		i, _ := strconv.Atoi(d + strconv.Itoa(n.Num))
		arrNumRatePrepare = append(arrNumRatePrepare, NumRank{
			Num:  i,
			Rate: n.Rate,
		})
		d = ""
	}

	return arrNumRatePrepare
}
