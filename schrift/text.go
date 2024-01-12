package schrift

import (
	_ "embed"
	"fmt"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	//go:embed roboto.ttf
	robotoDaten []byte

	NormaleSchriftart     font.Face
	NormaleSchriftartHöhe int

	TitelSchriftart     font.Face
	TitelSchriftartHöhe int
)

func init() {
	const dpi = 72

	robotoSchriftart, err := opentype.Parse(robotoDaten)
	if err != nil {
		panic(fmt.Errorf("Schriftart konnte nicht geladen werden: %w", err))
	}

	NormaleSchriftart, err = opentype.NewFace(robotoSchriftart, &opentype.FaceOptions{
		Size: 30,
		DPI:  dpi,
	})
	if err != nil {
		panic(fmt.Errorf("Schriftart konnte nicht instanziiert werden: %w", err))
	}

	NormaleSchriftartHöhe = NormaleSchriftart.Metrics().Ascent.Ceil() + NormaleSchriftart.Metrics().Descent.Ceil()

	TitelSchriftart, err = opentype.NewFace(robotoSchriftart, &opentype.FaceOptions{
		Size: 60,
		DPI:  dpi,
	})
	if err != nil {
		panic(fmt.Errorf("Schriftart konnte nicht instanziiert werden: %w", err))
	}

	TitelSchriftartHöhe = TitelSchriftart.Metrics().Ascent.Ceil() + TitelSchriftart.Metrics().Descent.Ceil()
}
