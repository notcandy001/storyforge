package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewTitleBar(_ fyne.Window, t *AppTheme, onToggle func()) fyne.CanvasObject {
	var titleCol, subCol, sepCol color.Color
	var toggleLabel string

	if t.IsParchment() {
		titleCol = color.NRGBA{R: 0x3a, G: 0x2e, B: 0x22, A: 0xff}
		subCol = color.NRGBA{R: 0x8a, G: 0x7a, B: 0x66, A: 0xff}
		sepCol = color.NRGBA{R: 0xc0, G: 0xb4, B: 0xa0, A: 0xff}
		toggleLabel = "◑  warm"
	} else {
		titleCol = color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0xff}
		subCol = color.NRGBA{R: 0x5c, G: 0x4a, B: 0x38, A: 0xff}
		sepCol = color.NRGBA{R: 0x2e, G: 0x26, B: 0x1e, A: 0xff}
		toggleLabel = "◑  paper"
	}

	appTitle := canvas.NewText("✦  The Tale of the Fool", titleCol)
	appTitle.TextStyle = fyne.TextStyle{Bold: true}
	appTitle.TextSize = 13

	sub := canvas.NewText("praizeTheFool · 31 May 2026", subCol)
	sub.TextSize = 11

	left := container.NewVBox(appTitle, sub)

	themeBtn := widget.NewButton(toggleLabel, onToggle)

	sep := canvas.NewLine(sepCol)
	sep.StrokeWidth = 1

	return container.NewVBox(
		container.NewPadded(container.NewBorder(nil, nil, left, themeBtn)),
		sep,
	)
}
