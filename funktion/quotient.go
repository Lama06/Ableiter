package funktion

import (
	"image/color"
	"math/big"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

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

func liesQuotienten(text string) Funktion {
	schnipsel := schneidenAußerhalbvonKlammern(text, '/')
	if len(schnipsel) != 2 {
		return nil
	}
	dividend := Lesen(schnipsel[0].inhalt)
	if dividend == nil {
		return nil
	}
	divisor := Lesen(schnipsel[1].inhalt)
	if divisor == nil {
		return nil
	}
	return Quotient{
		Dividend: dividend,
		Divisor:  divisor,
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

func (q Quotient) Zeichnen() *ebiten.Image {
	dividendBild := q.Dividend.Zeichnen()
	divisorBild := q.Divisor.Zeichnen()
	breite := max(dividendBild.Bounds().Dx(), divisorBild.Bounds().Dx())
	höhe := dividendBild.Bounds().Dy() + 10 + divisorBild.Bounds().Dy()
	img := ebiten.NewImage(breite, höhe)
	bruchstrichY := float32(dividendBild.Bounds().Dy() + 5)
	vector.StrokeLine(
		img,
		0, bruchstrichY,
		float32(breite), bruchstrichY,
		2,
		color.Black,
		true,
	)
	var dividendOptionen, divisorOptionen ebiten.DrawImageOptions
	dividendOptionen.GeoM.Translate(float64((breite-dividendBild.Bounds().Dx())/2), 0)
	divisorOptionen.GeoM.Translate(
		float64((breite-divisorBild.Bounds().Dx())/2), float64(dividendBild.Bounds().Dy()+10),
	)
	img.DrawImage(dividendBild, &dividendOptionen)
	img.DrawImage(divisorBild, &divisorOptionen)
	return img
}
