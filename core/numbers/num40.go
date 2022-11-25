package numbers

type Num40 struct {
	sNum    int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num40) Result() string {
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

func (n *Num40) nominative() string {
	return ""
}

func (n *Num40) genitive() string {
	return ""
}

func (n *Num40) dative() string {
	return ""
}

func (n *Num40) accusative() string {
	return ""
}

func (n *Num40) instrumental() string {
	return ""
}

func (n *Num40) prepositional() string {
	return ""
}
