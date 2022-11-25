package numbers

type Num1000000 struct {
	sNum    int
	sMany   int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num1000000) Result() string {
	text := ""
	switch n.sCase {
	case "И":
		text = n.nominative()
	case "Р":
		text = n.genitive()
	case "Д":
		text = n.dative()
	case "В":
		text = n.accusative()
	case "Т":
		text = n.instrumental()
	case "П":
		text = n.prepositional()
	}
	return text
}

func (n *Num1000000) nominative() string {
	text := "миллион"

	switch n.sMany {
	case 1:
		text = "миллион"
	case 2:
		text = "миллиона"
	case 3:
		text = "миллионов"
	}

	return text
}

func (n *Num1000000) genitive() string {
	text := "миллиона"

	switch n.sMany {
	case 1:
		text = "миллиону"
	case 2:
		text = "миллионов"
	case 3:
		text = "миллионов"
	}

	return text
}

func (n *Num1000000) dative() string {
	text := "миллиону"

	switch n.sMany {
	case 1:
		text = "миллиону"
	case 2:
		text = "миллионам"
	case 3:
		text = "миллионам"
	}

	return text
}

func (n *Num1000000) accusative() string {
	text := "миллионов"

	return text
}

func (n *Num1000000) instrumental() string {
	text := "миллионом"

	switch n.sMany {
	case 1:
		text = "миллионом"
	case 2:
		text = "миллионами"
	case 3:
		text = "миллионами"
	}

	return text
}

func (n *Num1000000) prepositional() string {
	text := "миллионах"

	switch n.sMany {
	case 1:
		text = "миллионе"
	case 2:
		text = "миллионах"
	case 3:
		text = "миллионах"
	}

	return text
}
