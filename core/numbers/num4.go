package numbers

type Num4 struct {
	sNum    int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num4) Result() string {
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

func (n *Num4) nominative() string {
	return "четыре"
}

func (n *Num4) genitive() string {
	return "четырёх"
}

func (n *Num4) dative() string {
	return "четырём"
}

func (n *Num4) accusative() string {
	return "четыре"
}

func (n *Num4) instrumental() string {
	return "четырьмя"
}

func (n *Num4) prepositional() string {
	return "четырёх"
}
