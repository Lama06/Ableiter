package funktion

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type Potenz struct {
	Basis    Funktion
	Exponent int
}

func liesPotenz(text string) Funktion {
	schnipsel := schneidenAußerhalbvonKlammern(text, '^')
	if len(schnipsel) != 2 {
		return nil
	}
	basis := Lesen(schnipsel[0].inhalt)
	if basis == nil {
		return nil
	}
	exponent, err := strconv.Atoi(schnipsel[1].inhalt)
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

func basisZeichnen(f Funktion) *ebiten.Image {
	bild := f.Zeichnen()
	switch fTypisiert := f.(type) {
	case *Konstante:
		if benötigtKonstanteKlammern(fTypisiert) {
			return umklammern(bild)
		}
	case Quotient, Summe, Produkt, Potenz:
		return umklammern(bild)
	}
	return bild
}

func (p Potenz) Zeichnen() *ebiten.Image {
	exponentBild := NeueKonstanteGanzzahl(p.Exponent).Zeichnen()
	basisBild := basisZeichnen(p.Basis)
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
