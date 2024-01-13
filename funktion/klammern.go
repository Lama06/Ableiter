package funktion

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func umklammern(umklammert *ebiten.Image) *ebiten.Image {
	img := ebiten.NewImage(umklammert.Bounds().Dx()+20, umklammert.Bounds().Dy()+10)
	var pfad vector.Path
	pfad.MoveTo(9, 0)
	pfad.QuadTo(1, float32(img.Bounds().Dy()/2), 9, float32(img.Bounds().Dy()))
	pfad.MoveTo(float32(img.Bounds().Dx()-9), 0)
	pfad.QuadTo(float32(img.Bounds().Dx()-1), float32(img.Bounds().Dy()/2), float32(img.Bounds().Dx()-9), float32(img.Bounds().Dy()))
	ecken, indizes := pfad.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
		Width: 3,
	})
	weiß := ebiten.NewImage(3, 3)
	weiß.Fill(color.Black)
	img.DrawTriangles(ecken, indizes, weiß, nil)
	var umklammertOptionen ebiten.DrawImageOptions
	umklammertOptionen.GeoM.Translate(10, 10)
	img.DrawImage(umklammert, &umklammertOptionen)
	return img
}

func liesUmklammert(text string) Funktion {
	if len(text) <= 2 {
		return nil
	}
	if text[0] != '(' || text[len(text)-1] != ')' {
		return nil
	}
	return Lesen(text[1 : len(text)-1])
}
