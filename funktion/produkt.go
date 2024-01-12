package funktion

import "math/big"

type Produkt []Funktion

func (p Produkt) Ableiten() Funktion {
	ergebnis := make(Summe, len(p))
	for i := 0; i < len(p); i++ {
		summand := make(Produkt, len(p))
		for j, faktor := range p {
			if i == j {
				summand[j] = faktor.Ableiten()
				continue
			}
			summand[j] = faktor
		}
		ergebnis[i] = Summand{
			Funktion:   summand,
			Vorzeichen: true,
		}
	}
	return ergebnis
}

func (p Produkt) faktorenVereinfachen() Produkt {
	ergebnis := make(Produkt, len(p))
	for i, faktor := range p {
		ergebnis[i] = faktor.Vereinfachen()
	}
	return ergebnis
}

func (p Produkt) konstantenMultiplizieren() Produkt {
	ergebnis := make(Produkt, 0, len(p))
	konstantenProdukt := big.NewRat(1, 1)
	for _, faktor := range p {
		if konstante, ok := faktor.(*Konstante); ok {
			konstantenProdukt.Mul(konstantenProdukt, (*big.Rat)(konstante))
			continue
		}
		ergebnis = append(ergebnis, faktor)
	}
	if konstantenProdukt.Cmp(big.NewRat(1, 1)) == 0 {
		return ergebnis
	}
	return append(ergebnis, (*Konstante)(konstantenProdukt))
}

func (p Produkt) unterprodukteEingliedern() Produkt {
	ergebnis := make(Produkt, 0, len(p))
	for _, faktor := range p {
		if unterprodukt, ok := faktor.(Produkt); ok {
			ergebnis = append(ergebnis, unterprodukt...)
			continue
		}
		ergebnis = append(ergebnis, faktor)
	}
	return ergebnis
}

func (p Produkt) faktorenSortieren() {

}

func (p Produkt) ggfAuflösen() Funktion {
	if len(p) == 0 {
		return NeueKonstanteGanzzahl(1)
	}
	if len(p) == 1 {
		return p[0]
	}
	for _, faktor := range p {
		if konstante, ok := faktor.(*Konstante); ok {
			if (*big.Rat)(konstante).Cmp(&big.Rat{}) == 0 {
				return NeueKonstanteGanzzahl(1)
			}
		}
	}
	return p
}

func (p Produkt) Vereinfachen() Funktion {
	return p.faktorenVereinfachen().
		unterprodukteEingliedern().
		konstantenMultiplizieren().
		ggfAuflösen()
}
