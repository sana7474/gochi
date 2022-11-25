package numbers

type Num100 struct {
	sNum    int
	sMany   int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num100) Result() string {
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

func (n *Num100) nominative() string {
	text := "сто"

	return text
}

func (n *Num100) genitive() string {
	text := "ста"

	return text
}

func (n *Num100) dative() string {
	text := "ста"

	return text
}

func (n *Num100) accusative() string {
	text := "сто"

	return text
}

func (n *Num100) instrumental() string {
	text := "ста"

	return text
}

func (n *Num100) prepositional() string {
	text := "ста"

	return text
}
