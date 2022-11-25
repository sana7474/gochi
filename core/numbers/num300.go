package numbers

type Num300 struct {
	sNum    int
	sMany   int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num300) Result() string {
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

func (n *Num300) nominative() string {
	text := "триста"

	return text
}

func (n *Num300) genitive() string {
	text := "трёхсот"

	return text
}

func (n *Num300) dative() string {
	text := "трёмстам"

	return text
}

func (n *Num300) accusative() string {
	text := "триста"

	return text
}

func (n *Num300) instrumental() string {
	text := "тремястами"

	return text
}

func (n *Num300) prepositional() string {
	text := "трёхстах"

	return text
}
