package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewTitleBar(_ fyne.Window) fyne.CanvasObject {
	appTitle := canvas.NewText("✦  The Tale of the Fool", color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0xff})
	appTitle.TextStyle = fyne.TextStyle{Bold: true}
	appTitle.TextSize = 13

	sub := canvas.NewText("praizeTheFool · 31 May 2026", color.NRGBA{R: 0x5c, G: 0x4a, B: 0x38, A: 0xff})
	sub.TextSize = 11

	left := container.NewVBox(appTitle, sub)
	tag := widget.NewRichTextFromMarkdown("**personal story**")

	sep := canvas.NewLine(color.NRGBA{R: 0x2e, G: 0x26, B: 0x1e, A: 0xff})
	sep.StrokeWidth = 1

	return container.NewVBox(
		container.NewPadded(container.NewBorder(nil, nil, left, tag)),
		sep,
	)
}
