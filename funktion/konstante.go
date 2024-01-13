package funktion

import (
	"image/color"
	"math/big"
	"strconv"

	"github.com/Lama06/Ableiter/schrift"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Konstante big.Rat

func liesKonstante(text string) Funktion {
	zahl, err := strconv.Atoi(text)
	if err != nil {
		return nil
	}
	return (*Konstante)(big.NewRat(int64(zahl), 1))
}

func NeueKonstanteGanzzahl(zahl int) *Konstante {
	return (*Konstante)(big.NewRat(int64(zahl), 1))
}

func NeueKonstanteBruch(zähler, nenner int) *Konstante {
	return (*Konstante)(big.NewRat(int64(zähler), int64(nenner)))
}

func (k *Konstante) Ableiten() Funktion {
	return NeueKonstanteGanzzahl(0)
}

func (k *Konstante) Vereinfachen() Funktion {
	return k
}

func benötigtKonstanteKlammern(f *Konstante) bool {
	bruch := (*big.Rat)(f)
	return !bruch.IsInt() || bruch.Sign() == -1
}

func (k *Konstante) Zeichnen() *ebiten.Image {
	bruch := (*big.Rat)(k)
	if !bruch.IsInt() {
		nenner, zähler := &big.Rat{}, &big.Rat{}
		nenner.SetInt(bruch.Num())
		zähler.SetInt(bruch.Denom())
		return Quotient{
			Dividend: (*Konstante)(nenner),
			Divisor:  (*Konstante)(zähler),
		}.Zeichnen()
	}
	konstanteText := bruch.Num().String()
	breite := font.MeasureString(schrift.NormaleSchriftart, konstanteText).Ceil()
	img := ebiten.NewImage(breite, schrift.NormaleSchriftartHöhe)
	img.Fill(color.White)
	text.Draw(
		img,
		konstanteText,
		schrift.NormaleSchriftart,
		0,
		schrift.NormaleSchriftart.Metrics().Ascent.Ceil(),
		color.Black,
	)
	return img
}
