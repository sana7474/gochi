package numbers

type Num1 struct {
	sNum    int
	sMany   int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num1) Result() string {
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

func (n *Num1) nominative() string {
	text := ""

	switch n.sGender {
	case "М":
		text = "один"
	case "С":
		text = "одно"
	case "Ж":
		text = "одна"
	}

	return text
}

func (n *Num1) genitive() string {
	text := ""

	switch n.sGender {
	case "М", "С":
		text = "одного"
	case "Ж":
		text = "одной"
	}

	return text
}

func (n *Num1) dative() string {
	text := ""

	switch n.sGender {
	case "М", "С":
		text = "одному"
	case "Ж":
		text = "одной"
	}

	return text
}

func (n *Num1) accusative() string {
	text := ""

	switch n.sGender {
	case "М":
		text = "один"
	case "С":
		text = "одно"
	case "Ж":
		text = "одну"
	}

	return text
}

func (n *Num1) instrumental() string {
	text := ""

	switch n.sGender {
	case "М", "С":
		text = "одним"
	case "Ж":
		text = "одной"
	}

	return text
}

func (n *Num1) prepositional() string {
	text := ""
	switch n.sGender {
	case "М", "С":
		text = "одном"
	case "Ж":
		text = "одной"
	}
	return text
}
