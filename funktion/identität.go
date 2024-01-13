package funktion

import (
	"image/color"

	"github.com/Lama06/Ableiter/schrift"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Identität struct{}

func liesIdentität(text string) Funktion {
	if text == "x" {
		return Identität{}
	}
	return nil
}

var _ Funktion = Identität{}

func (i Identität) Ableiten() Funktion {
	return NeueKonstanteGanzzahl(1)
}

func (i Identität) Vereinfachen() Funktion {
	return i
}

func (i Identität) Zeichnen() *ebiten.Image {
	const idText = "x"
	breite := font.MeasureString(schrift.NormaleSchriftart, idText).Ceil()
	img := ebiten.NewImage(breite, schrift.NormaleSchriftartHöhe)
	img.Fill(color.White)
	text.Draw(img, idText, schrift.NormaleSchriftart, 0, schrift.NormaleSchriftart.Metrics().Ascent.Ceil(), color.Black)
	return img
}
