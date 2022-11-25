package numbers

type Num5080 struct {
	sRoot   string
	sNum    int
	sGender string
	sCase   string
	sRate   string
}

func (n *Num5080) Result() string {
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

func (n *Num5080) nominative() string {
	return n.sRoot + "ьдесят"
}

func (n *Num5080) genitive() string {
	return n.sRoot + "идесяти"
}

func (n *Num5080) dative() string {
	return n.sRoot + "идесяти"
}

func (n *Num5080) accusative() string {
	return n.sRoot + "ьдесят"
}

func (n *Num5080) instrumental() string {
	return n.sRoot + "ьюдесятью"
}

func (n *Num5080) prepositional() string {
	return n.sRoot + "идесяти"
}
