package numbers

type Num500900 struct {
	sRoot   string
	sNum    int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num500900) Result() string {
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

func (n *Num500900) nominative() string {
	return n.sRoot + "ьсот"
}

func (n *Num500900) genitive() string {
	return n.sRoot + "исот"
}

func (n *Num500900) dative() string {
	return n.sRoot + "истам"
}

func (n *Num500900) accusative() string {
	return n.sRoot + "ьсот"
}

func (n *Num500900) instrumental() string {
	return n.sRoot + "ьюстами"
}

func (n *Num500900) prepositional() string {
	return n.sRoot + "истах"
}
