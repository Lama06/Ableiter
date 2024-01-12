package funktion

import (
	"strconv"
	"strings"
)

type Potenz struct {
	Basis    Funktion
	Exponent int
}

func ParsePotenz(text string) Funktion {
	if !strings.Contains(text, "**") {
		return nil
	}
	potenz := strings.Split(text, "**")
	if len(potenz) < 2 {
		return nil
	}
	var basisText strings.Builder
	for i := 0; i < len(potenz)-1; i++ {
		basisText.WriteString(potenz[i])
	}
	basis := Parse(basisText.String())
	if basis == nil {
		return nil
	}
	exponent, err := strconv.Atoi(potenz[len(potenz)-1])
	if err != nil {
		return nil
	}
	return Potenz{
		Basis:    basis,
		Exponent: exponent,
	}
}

func (p Potenz) Ableiten() Funktion {
	return Produkt{
		NeueKonstanteGanzzahl(p.Exponent),
		Potenz{
			Basis:    p.Basis,
			Exponent: p.Exponent - 1,
		},
		p.Basis.Ableiten(),
	}
}

func (p Potenz) Vereinfachen() Funktion {
	if p.Exponent == 0 {
		return NeueKonstanteGanzzahl(1)
	}
	if p.Exponent == 1 {
		return p.Basis
	}
	return Potenz{
		Basis:    p.Basis.Vereinfachen(),
		Exponent: p.Exponent,
	}
}
