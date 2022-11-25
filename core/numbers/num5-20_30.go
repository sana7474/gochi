package numbers

type Num52030 struct {
	sRoot   string
	sNum    int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num52030) Result() string {
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

func (n *Num52030) nominative() string {
	return n.sRoot + "ь"
}

func (n *Num52030) genitive() string {
	return n.sRoot + "и"
}

func (n *Num52030) dative() string {
	return n.sRoot + "и"
}

func (n *Num52030) accusative() string {
	return n.sRoot + "ь"
}

func (n *Num52030) instrumental() string {
	return n.sRoot + "ью"
}

func (n *Num52030) prepositional() string {
	return n.sRoot + "и"
}
