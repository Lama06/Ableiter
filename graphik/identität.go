package graphik

import (
	"image/color"

	"github.com/Lama06/Ableiter/schrift"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func identitätZeichnen() *ebiten.Image {
	const idText = "x"
	breite := font.MeasureString(schrift.NormaleSchriftart, idText).Ceil()
	img := ebiten.NewImage(breite, schrift.NormaleSchriftartHöhe)
	img.Fill(color.White)
	text.Draw(img, idText, schrift.NormaleSchriftart, 0, schrift.NormaleSchriftart.Metrics().Ascent.Ceil(), color.Black)
	return img
}
