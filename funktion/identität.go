package funktion

type Identität struct{}

func ParseIdentität(text string) Funktion {
	if text == "x" {
		return Identität{}
	}
	return nil
}

var _ Funktion = Identität{}

func (i Identität) Ableiten() Funktion {
	return NeueKonstanteGanzzahl(1)
}

func (i Identität) Vereinfachen() Funktion {
	return i
}
