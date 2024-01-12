package funktion

type Funktion interface {
	Ableiten() Funktion

	Vereinfachen() Funktion
}

func Parse(text string) Funktion {
	alleLeser := [...]func(string) Funktion{
		ParseKonstante, ParseIdentität, ParseUmklammert
	}
	for _, leser := range alleLeser {
		if f := leser(text); f != nil {
			return f
		}
	}
	return nil
}

func ParseUmklammert(text string) Funktion {
	if text[0] != '(' || text[len(text)-1] != ')' {
		return nil
	}
	return Parse(text[1:len(text)-1])
}