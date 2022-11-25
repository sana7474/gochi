package numbers

import (
	"strconv"
)

type Num interface {
	nominative() string
	genitive() string
	dative() string
	accusative() string
	instrumental() string
	prepositional() string
	Result() string
}

func Load(sNum int, sRate, sGender, sCase string) string {
	result := ""
	sText := strconv.Itoa(sNum)

	switch sRate {
	case "d":
		sText += "0"
	case "s":
		sText += "00"
	}

	if sRate == "tt" {
		sGender = "М"
	}

	if sRate == "t" {
		sGender = "Ж"
	}

	sMany := 3
	if sNum == 1 { // 1 тясяча, 2 тысячи, 15 тысяч
		sMany = 1
	}
	if sNum >= 2 && sNum <= 4 {
		sMany = 2
	}
	if sNum >= 5 && sNum <= 19 {
		sMany = 3
	}

	root := GetRoot(sNum, sRate)

	// 5-20	30
	// надо сделать на общем шаболне
	// fmt.Sprintf("%sи", "пят")
	// fmt.Sprintf("%sи", "тридцат")
	// В структуру нужных чисел надо добавить root значение пят

	sNum, _ = strconv.Atoi(sText)
	if sNum == 1 {
		num1 := &Num1{sMany: sMany, sRate: sRate, sGender: sGender, sCase: sCase}
		result += num1.Result()
	}
	if sNum == 2 {
		num2 := &Num2{sRate: sRate, sGender: sGender, sCase: sCase}
		result += num2.Result()
	}
	if sNum == 3 {
		num3 := &Num3{sRate: sRate, sGender: sGender, sCase: sCase}
		result += num3.Result()
	}
	if sNum == 4 {
		num4 := &Num4{sRate: sRate, sGender: sGender, sCase: sCase}
		result += num4.Result()
	}
	if (sNum >= 5 && sNum <= 20) || sNum == 30 {
		num52030 := &Num52030{sRoot: root, sRate: sRate, sGender: sGender, sCase: sCase}
		result += num52030.Result()
	}
	if sNum == 40 {
		num40 := &Num40{sRate: sRate, sGender: sGender, sCase: sCase}
		result += num40.Result()
	}
	if sNum == 90 {
		num90 := &Num90{sRate: sRate, sGender: sGender, sCase: sCase}
		result += num90.Result()
	}
	if sNum == 100 {
		num100 := &Num100{sRate: sRate, sGender: sGender, sCase: sCase}
		result += num100.Result()
	}
	if sNum >= 50 && sNum <= 80 {
		num5080 := &Num5080{sRoot: root, sRate: sRate, sGender: sGender, sCase: sCase}
		result += num5080.Result()
	}

	if sNum == 200 {
		num200 := &Num200{sRate: sRate, sGender: sGender, sCase: sCase}
		result += num200.Result()
	}
	if sNum == 300 {
		num300 := &Num300{sRate: sRate, sGender: sGender, sCase: sCase}
		result += num300.Result()
	}
	if sNum == 400 {
		num400 := &Num400{sRate: sRate, sGender: sGender, sCase: sCase}
		result += num400.Result()
	}

	if sNum == 500 || sNum == 600 || sNum == 700 || sNum == 900 {
		num500900 := &Num500900{sRoot: root, sRate: sRate, sGender: sGender, sCase: sCase}
		result += num500900.Result()
	}

	if sNum == 800 {
		num800 := &Num800{sRate: sRate, sGender: sGender, sCase: sCase}
		result += num800.Result()
	}

	if sRate == "ttt" {
		Num1000000000 := &Num1000000000{sMany: sMany, sRate: sRate, sGender: sGender, sCase: sCase}
		result += " " + Num1000000000.Result()
	}

	if sRate == "tt" {
		Num1000000 := &Num1000000{sMany: sMany, sRate: sRate, sGender: sGender, sCase: sCase}
		result += " " + Num1000000.Result()
	}

	if sRate == "t" {
		num1000 := &Num1000{sMany: sMany, sRate: sRate, sGender: sGender, sCase: sCase}
		result += " " + num1000.Result()
	}

	return result
}
func GetRoot(sNum int, rate string) string {
	root := ""
	switch strconv.Itoa(sNum) {
	case "3":
		if rate == "d" {
			root = "тридцат"
		}
	case "2":
		if rate == "d" {
			root = "двадцат"
		}
	case "5":
		root = "пят"
	case "6":
		root = "шест"
	case "7":
		root = "сем"
	case "8":
		root = "восьм"
	case "9":
		root = "девят"
	case "10":
		root = "десят"
	case "11":
		root = "одинацат"
	case "12":
		root = "двенадцат"
	case "13":
		root = "тринадцат"
	case "14":
		root = "четырнадцат"
	case "15":
		root = "пятнадцат"
	case "16":
		root = "шестнадцат"
	case "17":
		root = "семнадцат"
	case "18":
		root = "восемьнадцат"
	case "19":
		root = "девятнадцат"
	case "2d":
		root = "двадцат"
	case "3d":
		root = "тридцат"
	}
	return root
}
