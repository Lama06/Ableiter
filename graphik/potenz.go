package graphik

import (
	"github.com/Lama06/Ableiter/funktion"
	"github.com/hajimehoshi/ebiten/v2"
)

func basisZeichnen(f funktion.Funktion) *ebiten.Image {
	bild := FunktionZeichnen(f)
	switch fTypisiert := f.(type) {
	case *funktion.Konstante:
		if benötigtKonstanteKlammern(fTypisiert) {
			return umklammern(bild)
		}
	case funktion.Quotient, funktion.Summe, funktion.Produkt, funktion.Potenz:
		return umklammern(bild)
	}
	return bild
}

func potenzZeichnen(f funktion.Potenz) *ebiten.Image {
	exponentBild := konstanteZeichnen(funktion.NeueKonstanteGanzzahl(f.Exponent))
	basisBild := basisZeichnen(f.Basis)
	breite := exponentBild.Bounds().Dx() + basisBild.Bounds().Dx()
	höhe := max(exponentBild.Bounds().Dy(), basisBild.Bounds().Dy()) + 10
	img := ebiten.NewImage(breite, höhe)
	var basisOptionen, exponentOptionen ebiten.DrawImageOptions
	basisOptionen.GeoM.Translate(0, 10)
	exponentOptionen.GeoM.Translate(float64(basisBild.Bounds().Dx()), 0)
	img.DrawImage(basisBild, &basisOptionen)
	img.DrawImage(exponentBild, &exponentOptionen)
	return img
}
