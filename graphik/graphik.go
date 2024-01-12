package graphik

import (
	"github.com/Lama06/Ableiter/funktion"
	"github.com/hajimehoshi/ebiten/v2"
)

func FunktionZeichnen(f funktion.Funktion) *ebiten.Image {
	switch fTypisiert := f.(type) {
	case *funktion.Konstante:
		return konstanteZeichnen(fTypisiert)
	case funktion.Identität:
		return identitätZeichnen()
	case funktion.Summe:
		return summeZeichnen(fTypisiert)
	case funktion.Produkt:
		return produktZeichnen(fTypisiert)
	case funktion.Quotient:
		return quotientenZeichnen(fTypisiert)
	case funktion.Potenz:
		return potenzZeichnen(fTypisiert)
	case funktion.Sinus:
		return benannteFunktionZeichnen("sin", fTypisiert.Argument)
	case funktion.Kosinus:
		return benannteFunktionZeichnen("cos", fTypisiert.Argument)
	default:
		panic("fehlende Fallentscheidung")
	}
}
