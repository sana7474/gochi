package numbers

type Num3 struct {
	sNum    int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num3) Result() string {
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

func (n *Num3) nominative() string {
	return "три"
}

func (n *Num3) genitive() string {
	return "трёх"
}

func (n *Num3) dative() string {
	return "трём"
}

func (n *Num3) accusative() string {
	return "три"
}

func (n *Num3) instrumental() string {
	return "тремя"
}

func (n *Num3) prepositional() string {
	return "трёх"
}
