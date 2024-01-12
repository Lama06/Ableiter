package graphik

import (
	"image/color"

	"github.com/Lama06/Ableiter/funktion"
	"github.com/Lama06/Ableiter/schrift"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func summandenZeichnen(f funktion.Funktion) *ebiten.Image {
	bild := FunktionZeichnen(f)
	switch fTypisiert := f.(type) {
	case *funktion.Konstante:
		if benötigtKonstanteKlammern(fTypisiert) {
			return umklammern(bild)
		}
	}
	return bild
}

func summeZeichnen(f funktion.Summe) *ebiten.Image {
	minusBreite := font.MeasureString(schrift.NormaleSchriftart, "-").Ceil()
	plusBreite := font.MeasureString(schrift.NormaleSchriftart, "+").Ceil()

	summandenBilder := make([]*ebiten.Image, len(f))
	var höhe, breite int
	for i, summand := range f {
		if i != 0 || !summand.Vorzeichen {
			if summand.Vorzeichen {
				breite += plusBreite
			} else {
				breite += minusBreite
			}
		}
		summandenBilder[i] = summandenZeichnen(summand.Funktion)
		höhe = max(höhe, summandenBilder[i].Bounds().Dy())
		breite += summandenBilder[i].Bounds().Dx()
	}

	img := ebiten.NewImage(breite, höhe)
	img.Fill(color.White)
	var x int
	for i, summand := range f {
		if i != 0 || !summand.Vorzeichen {
			vorzeichenText := "-"
			if summand.Vorzeichen {
				vorzeichenText = "+"
			}
			text.Draw(
				img,
				vorzeichenText,
				schrift.NormaleSchriftart,
				x,
				(höhe-schrift.NormaleSchriftartHöhe)/2+schrift.NormaleSchriftart.Metrics().Ascent.Ceil(),
				color.Black,
			)
			x += font.MeasureString(schrift.NormaleSchriftart, vorzeichenText).Ceil()
		}
		bildHöhe := summandenBilder[i].Bounds().Dy()
		bildBreite := summandenBilder[i].Bounds().Dx()
		var einstellungen ebiten.DrawImageOptions
		einstellungen.GeoM.Translate(float64(x), float64((höhe-bildHöhe)/2))
		img.DrawImage(summandenBilder[i], &einstellungen)
		x += bildBreite
	}
	return img
}
