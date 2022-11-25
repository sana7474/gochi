package numbers

type Num8 struct {
	sMany   int
	sNum    int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num8) Result() string {
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

func (n *Num8) nominative() string {
	text := "восемь"

	return text
}

func (n *Num8) genitive() string {
	text := "восьми"

	return text
}

func (n *Num8) dative() string {
	text := "восьми"

	return text
}

func (n *Num8) accusative() string {
	text := "восемь"

	return text
}

func (n *Num8) instrumental() string {
	text := "восьмью"

	return text
}

func (n *Num8) prepositional() string {
	text := "восьми"

	return text
}
