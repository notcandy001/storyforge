package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// NewTitleBar creates the top bar showing app title.
func NewTitleBar(w fyne.Window) fyne.CanvasObject {
	appTitle := canvas.NewText("✦  The Tale of the Fool", color.NRGBA{R: 0xc4, G: 0x9a, B: 0x6c, A: 0xff})
	appTitle.TextStyle = fyne.TextStyle{Bold: true}
	appTitle.TextSize = 13

	sub := canvas.NewText("praizeTheFool · 30 Dec 2025", color.NRGBA{R: 0x66, G: 0x5e, B: 0x78, A: 0xff})
	sub.TextSize = 11

	left := container.NewVBox(appTitle, sub)

	tag := widget.NewRichTextFromMarkdown("**personal story**")

	bar := container.NewPadded(
		container.NewBorder(nil, nil, left, tag),
	)

	sep := canvas.NewLine(color.NRGBA{R: 0x2a, G: 0x26, B: 0x36, A: 0xff})
	sep.StrokeWidth = 1

	return container.NewVBox(
		container.NewPadded(bar),
		sep,
	)
}

// goldDivider creates a thin decorative gold line.
func goldDivider() fyne.CanvasObject {
	line := canvas.NewLine(color.NRGBA{R: 0xc4, G: 0x9a, B: 0x6c, A: 0x55})
	line.StrokeWidth = 1
	return container.New(layout.NewStackLayout(), line)
}

// dimText creates muted-color canvas text.
func dimText(s string, size float32) *canvas.Text {
	t := canvas.NewText(s, color.NRGBA{R: 0x66, G: 0x5e, B: 0x78, A: 0xff})
	t.TextSize = size
	return t
}

// accentText creates gold-colored canvas text.
func accentText(s string, size float32, bold bool) *canvas.Text {
	t := canvas.NewText(s, color.NRGBA{R: 0xc4, G: 0x9a, B: 0x6c, A: 0xff})
	t.TextSize = size
	t.TextStyle = fyne.TextStyle{Bold: bold}
	return t
}
