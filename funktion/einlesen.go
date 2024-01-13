package funktion

import (
	"strings"
)

func Lesen(text string) Funktion {
	text = strings.ReplaceAll(text, " ", "")
	alleLeser := [...]func(string) Funktion{
		liesUmklammert,
		liesIdentität,
		liesKonstante,
		liesSumme,
		liesProdukt,
		liesQuotienten,
		liesPotenz,
		liesSinus,
		liesKosinus,
	}
	for _, leser := range alleLeser {
		if f := leser(text); f != nil {
			return f
		}
	}
	return nil
}

type textSchipsel struct {
	inhalt       string
	trennerDavor rune
}

func schneidenAußerhalbvonKlammern(text string, trenner ...rune) []textSchipsel {
	var (
		schnipsel          []textSchipsel
		aktuellerSchnipsel strings.Builder
		letzterTrenner     rune
		buchstaben         = []rune(text)
		verschachtelung    int
	)
buchstaben:
	for i := 0; i < len(buchstaben); i++ {
		buchstabe := buchstaben[i]
		if buchstabe == '(' {
			verschachtelung++
			aktuellerSchnipsel.WriteRune(buchstabe)
			continue
		}
		if buchstabe == ')' {
			verschachtelung--
			aktuellerSchnipsel.WriteRune(buchstabe)
			if verschachtelung < 0 {
				return nil
			}
			continue
		}
		if verschachtelung != 0 {
			aktuellerSchnipsel.WriteRune(buchstabe)
			continue
		}
		for _, einTrenner := range trenner {
			if einTrenner == buchstabe {
				schnipsel = append(schnipsel, textSchipsel{
					inhalt:       aktuellerSchnipsel.String(),
					trennerDavor: letzterTrenner,
				})
				letzterTrenner = einTrenner
				aktuellerSchnipsel.Reset()
				continue buchstaben
			}
		}
		aktuellerSchnipsel.WriteRune(buchstabe)
	}
	if aktuellerSchnipsel.Len() != 0 {
		schnipsel = append(schnipsel, textSchipsel{
			inhalt:       aktuellerSchnipsel.String(),
			trennerDavor: letzterTrenner,
		})
	}
	return schnipsel
}
