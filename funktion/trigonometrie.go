package funktion

import "github.com/hajimehoshi/ebiten/v2"

type Sinus struct {
	Argument Funktion
}

var _ Funktion = Sinus{}

func liesSinus(text string) Funktion {
	return liesBenannt(text, "sin", func(argument Funktion) Funktion {
		return Sinus{Argument: argument}
	})
}

func (s Sinus) Ableiten() Funktion {
	return Produkt{
		Kosinus{Argument: s.Argument},
		s.Argument.Ableiten(),
	}
}

func (s Sinus) Vereinfachen() Funktion {
	return Sinus{Argument: s.Argument.Vereinfachen()}
}

func (s Sinus) Zeichnen() *ebiten.Image {
	return benannteFunktionZeichnen("sin", s.Argument)
}

type Kosinus struct {
	Argument Funktion
}

var _ Funktion = Kosinus{}

func liesKosinus(text string) Funktion {
	return liesBenannt(text, "cos", func(argument Funktion) Funktion {
		return Kosinus{Argument: argument}
	})
}

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

func (k Kosinus) Zeichnen() *ebiten.Image {
	return benannteFunktionZeichnen("cos", k.Argument)
}
