package graphik

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func umklammern(umklammert *ebiten.Image) *ebiten.Image {
	img := ebiten.NewImage(umklammert.Bounds().Dx()+10, umklammert.Bounds().Dy()+5)
	var pfad vector.Path
	pfad.MoveTo(5, 0)
	pfad.QuadTo(0, float32(img.Bounds().Dy()/2), 5, float32(img.Bounds().Dy()))
	pfad.MoveTo(float32(img.Bounds().Dx()-5), 0)
	pfad.QuadTo(float32(img.Bounds().Dx()), float32(img.Bounds().Dy()/2), float32(img.Bounds().Dx()-5), float32(img.Bounds().Dy()))
	ecken, indizes := pfad.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
		Width: 3,
	})
	weiß := ebiten.NewImage(3, 3)
	weiß.Fill(color.Black)
	img.DrawTriangles(ecken, indizes, weiß, nil)
	var umklammertOptionen ebiten.DrawImageOptions
	umklammertOptionen.GeoM.Translate(5, 5)
	img.DrawImage(umklammert, &umklammertOptionen)
	return img
}
