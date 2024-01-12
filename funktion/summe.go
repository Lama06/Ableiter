package funktion

import (
	"math/big"
)

type Summand struct {
	Vorzeichen bool
	Funktion   Funktion
}

type Summe []Summand

func Negieren(negiert Funktion) Funktion {
	return Summe{Summand{
		Funktion:   negiert,
		Vorzeichen: false,
	}}
}

var _ Funktion = Summe{}

func (s Summe) Ableiten() Funktion {
	ableitung := make(Summe, len(s))
	for i, summand := range s {
		ableitung[i] = Summand{
			Funktion:   summand.Funktion.Ableiten(),
			Vorzeichen: summand.Vorzeichen,
		}
	}
	return ableitung
}

func (s Summe) summandenVereinfachen() Summe {
	ergebnis := make(Summe, len(s))
	for i, summand := range s {
		ergebnis[i] = Summand{
			Funktion:   summand.Funktion.Vereinfachen(),
			Vorzeichen: summand.Vorzeichen,
		}
	}
	return ergebnis
}

func (s Summe) konstantenSummieren() Summe {
	ergebnis := make(Summe, 0, len(s))
	konstantenSumme := &big.Rat{}
	for _, summand := range s {
		if konstante, ok := summand.Funktion.(*Konstante); ok {
			if summand.Vorzeichen {
				konstantenSumme.Add(konstantenSumme, (*big.Rat)(konstante))
			} else {
				konstantenSumme.Sub(konstantenSumme, (*big.Rat)(konstante))
			}
			continue
		}
		ergebnis = append(ergebnis, summand)
	}
	if konstantenSumme.Cmp(&big.Rat{}) == 0 {
		return ergebnis
	}
	vorzeichen := konstantenSumme.Sign()
	return append(ergebnis, Summand{
		Funktion:   (*Konstante)(konstantenSumme.Abs(konstantenSumme)),
		Vorzeichen: vorzeichen == 1,
	})
}

func (s Summe) untersummenEingliedern() Summe {
	ergebnis := make(Summe, 0, len(s))
	for _, summand := range s {
		if untersumme, ok := summand.Funktion.(Summe); ok {
			if summand.Vorzeichen {
				ergebnis = append(ergebnis, untersumme...)
				continue
			}
			for _, untersummand := range untersumme {
				ergebnis = append(ergebnis, Summand{
					Vorzeichen: !untersummand.Vorzeichen,
					Funktion:   untersummand.Funktion,
				})
			}
			continue
		}
		ergebnis = append(ergebnis, summand)
	}
	return ergebnis
}

func (s Summe) summandenSortieren() {

}

func (s Summe) ggfAuflösen() Funktion {
	if len(s) == 0 {
		return NeueKonstanteGanzzahl(0)
	}
	if len(s) == 1 && s[0].Vorzeichen {
		return s[0].Funktion
	}
	return s
}

func (s Summe) Vereinfachen() Funktion {
	return s.summandenVereinfachen().
		untersummenEingliedern().
		konstantenSummieren().
		ggfAuflösen()
}
