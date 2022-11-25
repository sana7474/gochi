package numbers

type Num200 struct {
	sNum    int
	sMany   int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num200) Result() string {
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

func (n *Num200) nominative() string {
	text := "двести"

	return text
}

func (n *Num200) genitive() string {
	text := "двухсот"

	return text
}

func (n *Num200) dative() string {
	text := "двумстам"

	return text
}

func (n *Num200) accusative() string {
	text := "двести"

	return text
}

func (n *Num200) instrumental() string {
	text := "двумястами"

	return text
}

func (n *Num200) prepositional() string {
	text := "двухстах"

	return text
}
