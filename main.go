package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type screen interface {
	update()

	draw(screen *ebiten.Image)
}

type Spiel struct {
	aktuellerScreen screen
}

func (s *Spiel) Update() error {
	s.aktuellerScreen.update()
	return nil
}

func (s *Spiel) Draw(screen *ebiten.Image) {
	s.aktuellerScreen.draw(screen)
}

const breite, höhe = 1920, 1080

func (s *Spiel) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return breite, höhe
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.RunGame(&Spiel{
		aktuellerScreen: neuerAbleiterScreen(),
	})
}
