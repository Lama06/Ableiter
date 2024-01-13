package funktion

import "math/rand"

func PolynomGenerieren(tiefe int) Funktion {
	polynom := make(Summe, tiefe)
	for i := 0; i < tiefe; i++ {
		vorzeichen := rand.Float64() > 0.5
		koeffizient := NeueKonstanteGanzzahl(rand.Intn(11) + 1)
		if rand.Float64() < 0.3 {
			koeffizient = NeueKonstanteBruch(rand.Intn(rand.Intn(11)+1), rand.Intn(11)+1)
		}
		if i == tiefe-1 {
			polynom[i] = Summand{Vorzeichen: vorzeichen, Funktion: koeffizient}
			continue
		}
		if i == tiefe-2 {
			polynom[i] = Summand{Vorzeichen: vorzeichen, Funktion: Produkt{
				koeffizient,
				Identität{},
			}}
			continue
		}
		polynom[i] = Summand{Vorzeichen: vorzeichen, Funktion: Produkt{
			koeffizient,
			Potenz{
				Basis:    Identität{},
				Exponent: tiefe - 1 - i,
			},
		}}
	}
	return polynom
}
