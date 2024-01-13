package main

import (
	"image/color"
	"strconv"
	"strings"

	"github.com/Lama06/Ableiter/funktion"
	"github.com/Lama06/Ableiter/schrift"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

type taste struct {
	tastatur     *tastatur
	x, y         int
	beschriftung string
	callback     func()
}

func (t *taste) größe() (int, int) {
	textBreite := font.MeasureString(schrift.NormaleSchriftart, t.beschriftung).Ceil()
	return textBreite + 20, schrift.NormaleSchriftartHöhe + 10
}

func (t *taste) draw(screen *ebiten.Image) {
	breite, höhe := t.größe()
	vector.DrawFilledRect(screen, float32(t.x), float32(t.y), float32(breite), float32(höhe), colornames.Black, true)
	text.Draw(screen, t.beschriftung, schrift.NormaleSchriftart, t.x+10, t.y+schrift.NormaleSchriftart.Metrics().Ascent.Ceil()+5, color.White)
}

func (t *taste) drücken() {
	t.callback()
	t.tastatur.updateFunktion()
}

func (t *taste) klickBehandeln() {
	var mausX, mausY int
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		mausX, mausY = ebiten.CursorPosition()
	} else if touchIds := inpututil.AppendJustReleasedTouchIDs(nil); len(touchIds) == 1 {
		mausX, mausY = inpututil.TouchPositionInPreviousTick(touchIds[0])
	} else {
		return
	}

	breite, höhe := t.größe()
	if mausX < t.x || mausX > t.x+breite || mausY < t.y || mausY > t.y+höhe {
		return
	}
	t.drücken()
}

func (t *taste) update() {
	t.klickBehandeln()
}

type textTasteOptionen struct {
	beschriftung, text string
}

func neueTextTaste(tastatur *tastatur, daten textTasteOptionen, x, y int) *taste {
	positionVerschiebung := len(daten.text)
	if strings.ContainsRune(daten.text, '|') {
		positionVerschiebung = strings.IndexRune(daten.text, '|')
	}
	return &taste{
		tastatur:     tastatur,
		x:            x,
		y:            y,
		beschriftung: daten.beschriftung,
		callback: func() {
			davor, danach := tastatur.textVorCursor(), tastatur.textNachCursor()
			tastatur.text = davor + strings.ReplaceAll(daten.text, "|", "") + danach
			tastatur.position += positionVerschiebung
		},
	}
}

func neueCursorTaste(tastatur *tastatur, richtung, x, y int) *taste {
	beschriftung := "<-"
	if richtung == 1 {
		beschriftung = "->"
	}
	return &taste{
		tastatur:     tastatur,
		x:            x,
		y:            y,
		beschriftung: beschriftung,
		callback: func() {
			if richtung == 1 && tastatur.position < len(tastatur.text) {
				tastatur.position++
			}
			if richtung == -1 && tastatur.position > 0 {
				tastatur.position--
			}
		},
	}
}

func neueLöschenTaste(tastatur *tastatur, x, y int) *taste {
	return &taste{
		tastatur:     tastatur,
		x:            x,
		y:            y,
		beschriftung: "Löschen",
		callback: func() {
			tastatur.text = tastatur.textVorCursor()[:len(tastatur.textVorCursor())-1] + tastatur.textNachCursor()
			tastatur.position--
		},
	}
}

func neueAllesLöschenTaste(tastatur *tastatur, x, y int) *taste {
	return &taste{
		tastatur:     tastatur,
		x:            x,
		y:            y,
		beschriftung: "Alles löschen",
		callback: func() {
			tastatur.text = ""
			tastatur.position = 0
		},
	}
}

var alleTasten []func(tastatur *tastatur, x, y int) *taste

func init() {
	var textTasten []textTasteOptionen
	for i := 0; i < 10; i++ {
		iText := strconv.Itoa(i)
		textTasten = append(textTasten, textTasteOptionen{beschriftung: iText, text: iText})
	}
	textTasten = append(
		textTasten,
		textTasteOptionen{"Sinus", "sin(|)"},
		textTasteOptionen{"Kosinus", "cos(|)"},
		textTasteOptionen{"Mal", "*"},
		textTasteOptionen{"Plus", "+"},
		textTasteOptionen{"Minus", "-"},
		textTasteOptionen{"Hoch", "^"},
		textTasteOptionen{"Durch", "/"},
		textTasteOptionen{"Klammer", "(|)"},
		textTasteOptionen{"x", "x"},
	)

	for _, textTaste := range textTasten {
		textTaste := textTaste
		alleTasten = append(alleTasten, func(tastatur *tastatur, x, y int) *taste {
			return neueTextTaste(tastatur, textTaste, x, y)
		})
	}
	alleTasten = append(
		alleTasten,
		func(tastatur *tastatur, x, y int) *taste {
			return neueCursorTaste(tastatur, -1, x, y)
		},
		func(tastatur *tastatur, x, y int) *taste {
			return neueCursorTaste(tastatur, 1, x, y)
		},
		neueLöschenTaste, neueAllesLöschenTaste,
	)
}

type tastatur struct {
	funktionX, funktionY int
	letzteFunktion       funktion.Funktion
	letzteFunktionBild   *ebiten.Image

	textX, textY int
	text         string
	position     int

	tasten []*taste
}

func neueTastatur(y, textX, textY, funktionX, funktionY int) *tastatur {
	tastatur := tastatur{
		funktionX: funktionX, funktionY: funktionY,
		textX: textX, textY: textY,
		text:     "",
		position: 0,
	}
	y += 25
	x := 50
	tastatur.tasten = make([]*taste, len(alleTasten))
	for i, konstruktor := range alleTasten {
		tastatur.tasten[i] = konstruktor(&tastatur, x, y)
		tasteBreite, _ := tastatur.tasten[i].größe()
		x += tasteBreite + 50
		if x > breite-100 {
			y += 70
			x = 50
		}
	}
	return &tastatur
}

func (t *tastatur) textVorCursor() string {
	return t.text[:t.position]
}

func (t *tastatur) textNachCursor() string {
	return t.text[t.position:]
}

func (t *tastatur) draw(screen *ebiten.Image) {
	textMitCursor := t.textVorCursor() + "|" + t.textNachCursor()
	textBreite := font.MeasureString(schrift.NormaleSchriftart, textMitCursor).Ceil()
	text.Draw(
		screen,
		textMitCursor,
		schrift.NormaleSchriftart,
		t.textX-textBreite/2,
		t.textY-schrift.NormaleSchriftartHöhe/2+schrift.NormaleSchriftart.Metrics().Ascent.Ceil(),
		color.Black,
	)
	if t.letzteFunktionBild != nil {
		var funktionOptionen ebiten.DrawImageOptions
		funktionOptionen.GeoM.Translate(float64(t.funktionX-t.letzteFunktionBild.Bounds().Dx()/2), float64(t.funktionY-t.letzteFunktionBild.Bounds().Dy()/2))
		screen.DrawImage(t.letzteFunktionBild, &funktionOptionen)
	}
	for _, taste := range t.tasten {
		taste.draw(screen)
	}
}

func (t *tastatur) updateFunktion() {
	f := funktion.Lesen(t.text)
	if f != nil {
		t.letzteFunktion = f
		t.letzteFunktionBild = f.Zeichnen()
	}
}

func (t *tastatur) update() {
	for _, taste := range t.tasten {
		taste.update()
	}
	eingabe := string(ebiten.AppendInputChars(nil))
	if len(eingabe) != 0 {
		t.text = t.textVorCursor() + eingabe + t.textNachCursor()
		t.position += len(eingabe)
		t.updateFunktion()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) && t.position > 0 {
		t.position--
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) && t.position < len(t.text) {
		t.position++
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyBackspace) && len(t.textVorCursor()) != 0 {
		t.text = t.textVorCursor()[:len(t.textVorCursor())-1] + t.textNachCursor()
		t.position--
		t.updateFunktion()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyDelete) {
		t.text = ""
		t.position = 0
	}
}
