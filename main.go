package main

import (
	"github.com/Lama06/Ableiter/funktion"
	"github.com/Lama06/Ableiter/graphik"
	"github.com/hajimehoshi/ebiten/v2"
)

type Spiel struct {
}

func (s *Spiel) Update() error {
	return nil

}

func (s *Spiel) Draw(screen *ebiten.Image) {
	summe := funktion.Kosinus{
		Argument: funktion.Summe{
			funktion.Summand{false, funktion.Identität{}},
			funktion.Summand{true, funktion.Potenz{
				Basis:    funktion.Identität{},
				Exponent: 3,
			}},
		},
	}.Ableiten().Vereinfachen()
	screen.DrawImage(graphik.FunktionZeichnen(summe), nil)
}

func (s *Spiel) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.RunGame(&Spiel{})
}
