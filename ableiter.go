package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ableiter struct {
	eingabe   *tastatur
	ableitung *ebiten.Image
}

func neuerAbleiterScreen() *ableiter {
	return &ableiter{
		eingabe: neueTastatur(höhe-300, breite/2, 100, breite/2, 300),
	}
}

func (a *ableiter) draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	a.eingabe.draw(screen)
	if a.ableitung != nil {
		var optionen ebiten.DrawImageOptions
		optionen.GeoM.Translate(float64(breite/2-a.ableitung.Bounds().Dx()/2), höhe/2)
		screen.DrawImage(a.ableitung, &optionen)
	}
}

func (a *ableiter) updateAbleitung() {
	a.ableitung = a.eingabe.letzteFunktion.Vereinfachen().Ableiten().Vereinfachen().Zeichnen()
}

func (a *ableiter) update() {
	touchIds := inpututil.AppendJustReleasedTouchIDs(nil)
	if len(touchIds) != 0 {
		_, y := inpututil.TouchPositionInPreviousTick(touchIds[0])
		if y < höhe/2 {
			a.updateAbleitung()
		}
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) && a.eingabe.letzteFunktion != nil {
		a.updateAbleitung()
	}
	a.eingabe.update()
}
