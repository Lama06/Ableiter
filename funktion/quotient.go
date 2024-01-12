package funktion

import "math/big"

type Quotient struct {
	Dividend Funktion
	Divisor  Funktion
}

func Kehrwert(f Funktion) Funktion {
	return Quotient{
		Dividend: NeueKonstanteGanzzahl(1),
		Divisor:  f,
	}
}

func (q Quotient) Ableiten() Funktion {
	return Quotient{
		Dividend: Summe{
			Summand{
				Vorzeichen: true,
				Funktion:   Produkt{q.Dividend.Ableiten(), q.Divisor},
			},
			Summand{
				Vorzeichen: false,
				Funktion:   Produkt{q.Dividend, q.Divisor.Ableiten()},
			},
		},
		Divisor: Potenz{
			Basis:    q.Divisor,
			Exponent: 2,
		},
	}
}

func (q Quotient) Vereinfachen() Funktion {
	zähler, zählerOk := q.Dividend.(*Konstante)
	nenner, nennerOk := q.Divisor.(*Konstante)
	if zählerOk && nennerOk {
		bruch := &big.Rat{}
		bruch.Quo((*big.Rat)(zähler), (*big.Rat)(nenner))
		return (*Konstante)(bruch)
	}

	return Quotient{
		Dividend: q.Dividend.Vereinfachen(),
		Divisor:  q.Divisor.Vereinfachen(),
	}
}
