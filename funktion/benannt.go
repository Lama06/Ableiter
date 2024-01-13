package funktion

import (
	"image/color"
	"strings"

	"github.com/Lama06/Ableiter/schrift"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func benannteFunktionZeichnen(name string, argument Funktion) *ebiten.Image {
	nameBreite := font.MeasureString(schrift.NormaleSchriftart, name).Ceil()
	argumentBild := umklammern(argument.Zeichnen())
	breite := nameBreite + argumentBild.Bounds().Dx()
	höhe := max(schrift.NormaleSchriftartHöhe, argumentBild.Bounds().Dy())
	bild := ebiten.NewImage(breite, höhe)
	bild.Fill(color.White)
	text.Draw(
		bild,
		name,
		schrift.NormaleSchriftart,
		0,
		(höhe-schrift.NormaleSchriftartHöhe)/2+schrift.NormaleSchriftart.Metrics().Ascent.Ceil(),
		color.Black,
	)
	var argumentOptionen ebiten.DrawImageOptions
	argumentOptionen.GeoM.Translate(float64(nameBreite), float64((höhe-argumentBild.Bounds().Dy())/2))
	bild.DrawImage(argumentBild, &argumentOptionen)
	return bild
}

func liesBenannt(text string, name string, konstruktor func(argument Funktion) Funktion) Funktion {
	argument := strings.TrimPrefix(text, name)
	if argument == text {
		return nil
	}
	argumentFunktion := liesUmklammert(argument)
	if argumentFunktion == nil {
		return nil
	}
	return konstruktor(argumentFunktion)
}
