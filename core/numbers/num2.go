package numbers

type Num2 struct {
	sNum    int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num2) Result() string {
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

func (n *Num2) nominative() string {
	text := ""

	switch n.sGender {
	case "М", "С":
		text = "два"
	case "Ж":
		text = "две"
	}

	return text
}

func (n *Num2) genitive() string {
	text := "двух"

	return text
}

func (n *Num2) dative() string {
	text := "двум"

	return text
}

func (n *Num2) accusative() string {
	text := ""

	switch n.sGender {
	case "М", "С":
		text = "два"
	case "Ж":
		text = "две"
	}

	return text
}

func (n *Num2) instrumental() string {
	text := "двумя"

	return text
}

func (n *Num2) prepositional() string {
	text := "двух"

	return text
}
