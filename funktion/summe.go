package funktion

import (
	"image/color"
	"math/big"

	"github.com/Lama06/Ableiter/schrift"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Summand struct {
	Vorzeichen bool
	Funktion   Funktion
}

type Summe []Summand

var _ Funktion = Summe{}

func Negieren(negiert Funktion) Funktion {
	return Summe{Summand{
		Funktion:   negiert,
		Vorzeichen: false,
	}}
}

func liesSumme(text string) Funktion {
	schnipsel := schneidenAußerhalbvonKlammern(text, '+', '-')
	if schnipsel == nil {
		return nil
	}
	if len(schnipsel) >= 2 && schnipsel[0].inhalt == "" && schnipsel[1].trennerDavor == '-' {
		schnipsel = schnipsel[1:]
	}
	if len(schnipsel) == 1 {
		return nil
	}
	summe := make(Summe, len(schnipsel))
	for i, einSchnipsel := range schnipsel {
		summand := Lesen(einSchnipsel.inhalt)
		if summand == nil {
			return nil
		}
		summe[i] = Summand{
			Vorzeichen: einSchnipsel.trennerDavor == '+' || einSchnipsel.trennerDavor == 0,
			Funktion:   summand,
		}
	}
	return summe
}

func (s Summe) Ableiten() Funktion {
	ableitung := make(Summe, len(s))
	for i, summand := range s {
		ableitung[i] = Summand{
			Funktion:   summand.Funktion.Ableiten(),
			Vorzeichen: summand.Vorzeichen,
		}
	}
	return ableitung
}

func (s Summe) summandenVereinfachen() Summe {
	ergebnis := make(Summe, len(s))
	for i, summand := range s {
		ergebnis[i] = Summand{
			Funktion:   summand.Funktion.Vereinfachen(),
			Vorzeichen: summand.Vorzeichen,
		}
	}
	return ergebnis
}

func (s Summe) konstantenSummieren() Summe {
	ergebnis := make(Summe, 0, len(s))
	konstantenSumme := &big.Rat{}
	for _, summand := range s {
		if konstante, ok := summand.Funktion.(*Konstante); ok {
			if summand.Vorzeichen {
				konstantenSumme.Add(konstantenSumme, (*big.Rat)(konstante))
			} else {
				konstantenSumme.Sub(konstantenSumme, (*big.Rat)(konstante))
			}
			continue
		}
		ergebnis = append(ergebnis, summand)
	}
	if konstantenSumme.Cmp(&big.Rat{}) == 0 {
		return ergebnis
	}
	vorzeichen := konstantenSumme.Sign()
	return append(ergebnis, Summand{
		Funktion:   (*Konstante)(konstantenSumme.Abs(konstantenSumme)),
		Vorzeichen: vorzeichen == 1,
	})
}

func (s Summe) untersummenEingliedern() Summe {
	ergebnis := make(Summe, 0, len(s))
	for _, summand := range s {
		if untersumme, ok := summand.Funktion.(Summe); ok {
			if summand.Vorzeichen {
				ergebnis = append(ergebnis, untersumme...)
				continue
			}
			for _, untersummand := range untersumme {
				ergebnis = append(ergebnis, Summand{
					Vorzeichen: !untersummand.Vorzeichen,
					Funktion:   untersummand.Funktion,
				})
			}
			continue
		}
		ergebnis = append(ergebnis, summand)
	}
	return ergebnis
}

func (s Summe) ggfAuflösen() Funktion {
	if len(s) == 0 {
		return NeueKonstanteGanzzahl(0)
	}
	if len(s) == 1 {
		if s[0].Vorzeichen {
			return s[0].Funktion
		}
		if konstante, ok := s[0].Funktion.(*Konstante); ok {
			negiert := (&big.Rat{}).Neg((*big.Rat)(konstante))
			return (*Konstante)(negiert)
		}
	}
	return s
}

func (s Summe) Vereinfachen() Funktion {
	return s.summandenVereinfachen().
		untersummenEingliedern().
		konstantenSummieren().
		ggfAuflösen()
}

func summandenZeichnen(f Funktion) *ebiten.Image {
	bild := f.Zeichnen()
	switch fTypisiert := f.(type) {
	case *Konstante:
		if benötigtKonstanteKlammern(fTypisiert) {
			return umklammern(bild)
		}
	case Summe:
		return umklammern(bild)
	}
	return bild
}

func (s Summe) Zeichnen() *ebiten.Image {
	minusBreite := font.MeasureString(schrift.NormaleSchriftart, "-").Ceil()
	plusBreite := font.MeasureString(schrift.NormaleSchriftart, "+").Ceil()

	summandenBilder := make([]*ebiten.Image, len(s))
	var höhe, breite int
	for i, summand := range s {
		if i != 0 || !summand.Vorzeichen {
			if summand.Vorzeichen {
				breite += plusBreite
			} else {
				breite += minusBreite
			}
		}
		summandenBilder[i] = summandenZeichnen(summand.Funktion)
		höhe = max(höhe, summandenBilder[i].Bounds().Dy())
		breite += summandenBilder[i].Bounds().Dx()
	}

	img := ebiten.NewImage(breite, höhe)
	img.Fill(color.White)
	var x int
	for i, summand := range s {
		if i != 0 || !summand.Vorzeichen {
			vorzeichenText := "-"
			if summand.Vorzeichen {
				vorzeichenText = "+"
			}
			text.Draw(
				img,
				vorzeichenText,
				schrift.NormaleSchriftart,
				x,
				(höhe-schrift.NormaleSchriftartHöhe)/2+schrift.NormaleSchriftart.Metrics().Ascent.Ceil(),
				color.Black,
			)
			x += font.MeasureString(schrift.NormaleSchriftart, vorzeichenText).Ceil()
		}
		bildHöhe := summandenBilder[i].Bounds().Dy()
		bildBreite := summandenBilder[i].Bounds().Dx()
		var einstellungen ebiten.DrawImageOptions
		einstellungen.GeoM.Translate(float64(x), float64((höhe-bildHöhe)/2))
		img.DrawImage(summandenBilder[i], &einstellungen)
		x += bildBreite
	}
	return img
}
