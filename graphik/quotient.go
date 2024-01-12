package graphik

import (
	"image/color"

	"github.com/Lama06/Ableiter/funktion"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func quotientenZeichnen(f funktion.Quotient) *ebiten.Image {
	dividendBild := FunktionZeichnen(f.Dividend)
	divisorBild := FunktionZeichnen(f.Divisor)
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
