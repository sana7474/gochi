package numbers

type Num1000000000 struct {
	sNum    int
	sMany   int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num1000000000) Result() string {
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

func (n *Num1000000000) nominative() string {
	text := "миллиард"

	switch n.sMany {
	case 1:
		text = "миллиард"
	case 2:
		text = "миллиарда"
	case 3:
		text = "миллиардов"
	}

	return text
}

func (n *Num1000000000) genitive() string {
	text := "миллиарда"

	switch n.sMany {
	case 1:
		text = "миллиарда"
	case 2, 3:
		text = "миллиардов"
	}

	return text
}

func (n *Num1000000000) dative() string {
	text := "миллиарду"

	switch n.sMany {
	case 1:
		text = "миллиарду"
	case 2, 3:
		text = "миллиардам"
	}

	return text
}

func (n *Num1000000000) accusative() string {
	text := "миллиард"

	switch n.sMany {
	case 1:
		text = "миллиард"
	case 2:
		text = "миллиарда"
	case 3:
		text = "миллиардов"
	}

	return text
}

func (n *Num1000000000) instrumental() string {
	text := "миллиардом"

	switch n.sMany {
	case 1:
		text = "миллиардом"
	case 2, 3:
		text = "миллиардом"
	}

	return text
}

func (n *Num1000000000) prepositional() string {
	text := "миллиарде"

	switch n.sMany {
	case 1:
		text = "миллиарде"
	case 2, 3:
		text = "миллиардах"
	}

	return text
}
