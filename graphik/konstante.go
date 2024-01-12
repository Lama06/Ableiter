package graphik

import (
	"image/color"
	"math/big"

	"github.com/Lama06/Ableiter/funktion"
	"github.com/Lama06/Ableiter/schrift"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func benötigtKonstanteKlammern(f *funktion.Konstante) bool {
	bruch := (*big.Rat)(f)
	return !bruch.IsInt() || bruch.Sign() == -1
}

func konstanteZeichnen(f *funktion.Konstante) *ebiten.Image {
	bruch := (*big.Rat)(f)
	if !bruch.IsInt() {
		nenner, zähler := &big.Rat{}, &big.Rat{}
		nenner.SetInt(bruch.Num())
		zähler.SetInt(bruch.Denom())
		return quotientenZeichnen(funktion.Quotient{
			Dividend: (*funktion.Konstante)(nenner),
			Divisor:  (*funktion.Konstante)(zähler),
		})
	}
	konstanteText := bruch.Num().String()
	breite := font.MeasureString(schrift.NormaleSchriftart, konstanteText).Ceil()
	img := ebiten.NewImage(breite, schrift.NormaleSchriftartHöhe)
	img.Fill(color.White)
	text.Draw(img, konstanteText, schrift.NormaleSchriftart, 0, schrift.NormaleSchriftart.Metrics().Ascent.Ceil(), color.Black)
	return img
}
