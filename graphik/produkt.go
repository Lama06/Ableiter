package graphik

import (
	"image/color"

	"github.com/Lama06/Ableiter/funktion"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func faktorenZeichnen(f funktion.Funktion) *ebiten.Image {
	bild := FunktionZeichnen(f)
	switch fTypisiert := f.(type) {
	case *funktion.Konstante:
		if benötigtKonstanteKlammern(fTypisiert) {
			return umklammern(bild)
		}
	case funktion.Summe:
		return umklammern(bild)
	}
	return bild
}

func produktZeichnen(f funktion.Produkt) *ebiten.Image {
	const (
		malPunktBreite = 15
		malPunktRadius = 3
	)

	faktorenBilder := make([]*ebiten.Image, len(f))
	var höhe, breite int
	for i, faktor := range f {
		if i != 0 {
			breite += malPunktBreite
		}
		faktorenBilder[i] = faktorenZeichnen(faktor)
		höhe = max(höhe, faktorenBilder[i].Bounds().Dy())
		breite += faktorenBilder[i].Bounds().Dx()
	}

	img := ebiten.NewImage(breite, höhe)
	img.Fill(color.White)
	var x int
	for i := range f {
		if i != 0 {
			vector.DrawFilledCircle(img, float32(x+malPunktBreite/2), float32(höhe/2), malPunktRadius, color.Black, true)
			x += malPunktBreite
		}
		bildHöhe := faktorenBilder[i].Bounds().Dy()
		bildBreite := faktorenBilder[i].Bounds().Dx()
		var einstellungen ebiten.DrawImageOptions
		einstellungen.GeoM.Translate(float64(x), float64((höhe-bildHöhe)/2))
		img.DrawImage(faktorenBilder[i], &einstellungen)
		x += bildBreite
	}
	return img
}
