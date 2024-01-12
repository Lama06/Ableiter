package funktion

import (
	"math/big"
	"strconv"
	"strings"
)

type Konstante big.Rat

func ParseKonstante(text string) Funktion {
	if strings.Contains(text, "/") {
		bruch := strings.Split(text, "/")
		if len(bruch) != 2 {
			return nil
		}
		zähler, err := strconv.Atoi(bruch[0])
		if err != nil {
			return nil
		}
		nenner, err := strconv.Atoi(bruch[1])
		if err != nil {
			return nil
		}
		return (*Konstante)(big.NewRat(int64(zähler), int64(nenner)))
	}

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
	return (*Konstante)(&big.Rat{})
}

func (k *Konstante) Vereinfachen() Funktion {
	return k
}
