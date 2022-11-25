package numbers

type Num800 struct {
	sNum    int
	sMany   int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num800) Result() string {
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

func (n *Num800) nominative() string {
	text := "восемьсот"

	return text
}

func (n *Num800) genitive() string {
	text := "восьмисот"

	return text
}

func (n *Num800) dative() string {
	text := "восьмистам"

	return text
}

func (n *Num800) accusative() string {
	text := "восемьсот"

	return text
}

func (n *Num800) instrumental() string {
	text := "восемьюстами"

	return text
}

func (n *Num800) prepositional() string {
	text := "восьмистам"

	return text
}
