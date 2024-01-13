package funktion

import (
	"image/color"
	"math/big"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Produkt []Funktion

func liesProdukt(text string) Funktion {
	schnipsel := schneidenAußerhalbvonKlammern(text, '*')
	if len(schnipsel) < 2 {
		return nil
	}
	produkt := make(Produkt, len(schnipsel))
	for i, einSchnipsel := range schnipsel {
		produkt[i] = Lesen(einSchnipsel.inhalt)
		if produkt[i] == nil {
			return nil
		}
	}
	return produkt
}

func (p Produkt) Ableiten() Funktion {
	ergebnis := make(Summe, len(p))
	for i := 0; i < len(p); i++ {
		summand := make(Produkt, len(p))
		for j, faktor := range p {
			if i == j {
				summand[j] = faktor.Ableiten()
				continue
			}
			summand[j] = faktor
		}
		ergebnis[i] = Summand{
			Funktion:   summand,
			Vorzeichen: true,
		}
	}
	return ergebnis
}

func (p Produkt) faktorenVereinfachen() Produkt {
	ergebnis := make(Produkt, len(p))
	for i, faktor := range p {
		ergebnis[i] = faktor.Vereinfachen()
	}
	return ergebnis
}

func (p Produkt) konstantenMultiplizieren() Produkt {
	ergebnis := make(Produkt, 0, len(p))
	konstantenProdukt := big.NewRat(1, 1)
	for _, faktor := range p {
		if konstante, ok := faktor.(*Konstante); ok {
			konstantenProdukt.Mul(konstantenProdukt, (*big.Rat)(konstante))
			continue
		}
		ergebnis = append(ergebnis, faktor)
	}
	if konstantenProdukt.Cmp(big.NewRat(1, 1)) == 0 {
		return ergebnis
	}
	return append(ergebnis, (*Konstante)(konstantenProdukt))
}

func (p Produkt) unterprodukteEingliedern() Produkt {
	ergebnis := make(Produkt, 0, len(p))
	for _, faktor := range p {
		if unterprodukt, ok := faktor.(Produkt); ok {
			ergebnis = append(ergebnis, unterprodukt...)
			continue
		}
		ergebnis = append(ergebnis, faktor)
	}
	return ergebnis
}

func (p Produkt) ggfAuflösen() Funktion {
	if len(p) == 0 {
		return NeueKonstanteGanzzahl(1)
	}
	if len(p) == 1 {
		return p[0]
	}
	for _, faktor := range p {
		if konstante, ok := faktor.(*Konstante); ok {
			if (*big.Rat)(konstante).Cmp(&big.Rat{}) == 0 {
				return NeueKonstanteGanzzahl(0)
			}
		}
	}
	return p
}

func (p Produkt) Vereinfachen() Funktion {
	return p.faktorenVereinfachen().
		unterprodukteEingliedern().
		konstantenMultiplizieren().
		ggfAuflösen()
}

func faktorenZeichnen(f Funktion) *ebiten.Image {
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

func (p Produkt) Zeichnen() *ebiten.Image {
	const (
		malPunktBreite = 15
		malPunktRadius = 3
	)

	faktorenBilder := make([]*ebiten.Image, len(p))
	var höhe, breite int
	for i, faktor := range p {
		if i != 0 {
			breite += malPunktBreite
		}
		faktorenBilder[i] = faktorenZeichnen(faktor)
		höhe = max(höhe, faktorenBilder[i].Bounds().Dy())
		breite += faktorenBilder[i].Bounds().Dx()
	}

	img := ebiten.NewImage(breite, höhe)
	img.Fill(color.White)
	var x int
	for i := range p {
		if i != 0 {
			vector.DrawFilledCircle(img, float32(x+malPunktBreite/2), float32(höhe/2), malPunktRadius, color.Black, true)
			x += malPunktBreite
		}
		bildHöhe := faktorenBilder[i].Bounds().Dy()
		bildBreite := faktorenBilder[i].Bounds().Dx()
		var einstellungen ebiten.DrawImageOptions
		einstellungen.GeoM.Translate(float64(x), float64((höhe-bildHöhe)/2))
		img.DrawImage(faktorenBilder[i], &einstellungen)
		x += bildBreite
	}
	return img
}
