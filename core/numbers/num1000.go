package numbers

type Num1000 struct {
	sNum    int
	sMany   int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num1000) Result() string {
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

func (n *Num1000) nominative() string {
	text := "тысяч "

	switch n.sMany {
	case 1:
		text = "тысяча"
	case 2:
		text = "тысячи"
	case 3:
		text = "тысяч"
	}

	return text
}

func (n *Num1000) genitive() string {
	text := "тысяч"

	switch n.sMany {
	case 1:
		text = "тысячи"
	case 2:
		text = "тысяч"
	case 3:
		text = "тысяч"
	}

	return text
}

func (n *Num1000) dative() string {
	text := "тысячам"

	switch n.sMany {
	case 1:
		text = "тысяче"
	case 2:
		text = "тысячам"
	case 3:
		text = "тысячам"
	}

	return text
}

func (n *Num1000) accusative() string {
	text := "тысяч"

	switch n.sMany {
	case 1:
		text = "тысячу"
	case 2:
		text = "тысячи"
	case 3:
		text = "тысяч"
	}

	return text
}

func (n *Num1000) instrumental() string {
	text := "тысячами"

	switch n.sMany {
	case 1:
		text = "тысячей"
	case 2, 3:
		text = "тысячами"
	}

	return text
}

func (n *Num1000) prepositional() string {
	text := "тысячах"

	switch n.sMany {
	case 1:
		text = "тысяче"
	case 2, 3:
		text = "тысячах"
	}

	return text
}
