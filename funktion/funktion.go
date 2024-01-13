package funktion

import "github.com/hajimehoshi/ebiten/v2"

type Funktion interface {
	Ableiten() Funktion

	Vereinfachen() Funktion

	Zeichnen() *ebiten.Image
}
