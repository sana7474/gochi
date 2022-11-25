package numbers

type Num90 struct {
	sNum    int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num90) Result() string {
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

func (n *Num90) nominative() string {
	return ""
}

func (n *Num90) genitive() string {
	return ""
}

func (n *Num90) dative() string {
	return ""
}

func (n *Num90) accusative() string {
	return ""
}

func (n *Num90) instrumental() string {
	return ""
}

func (n *Num90) prepositional() string {
	return ""
}
