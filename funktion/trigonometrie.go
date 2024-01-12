package funktion

type Sinus struct {
	Argument Funktion
}

var _ Funktion = Sinus{}

func (s Sinus) Ableiten() Funktion {
	return Produkt{
		Kosinus{Argument: s.Argument},
		s.Argument.Ableiten(),
	}
}

func (s Sinus) Vereinfachen() Funktion {
	return Sinus{Argument: s.Argument.Vereinfachen()}
}

type Kosinus struct {
	Argument Funktion
}

var _ Funktion = Kosinus{}

func (k Kosinus) Ableiten() Funktion {
	return Produkt{
		Negieren(Sinus{Argument: k.Argument}),
		k.Argument.Ableiten(),
	}
}

func (k Kosinus) Vereinfachen() Funktion {
	return Kosinus{
		Argument: k.Argument.Vereinfachen(),
	}
}
