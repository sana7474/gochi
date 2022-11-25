package core

func Prepare1000(num int) []int {
	var list []int

	for num != 0 {
		part := num % 1000
		num /= 1000
		list = append(list, part)
	}

	return list
}

func Prepare10(num int) []int {
	var list []int

	for num != 0 {
		part := num % 10
		num /= 10
		list = append(list, part)
	}

	return list
}
