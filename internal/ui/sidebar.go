package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"foolstory/internal/data"
)

func NewChapterSidebar(story *data.Story, t *AppTheme, onSelect func(int)) fyne.CanvasObject {
	var headerCol, divCol color.Color
	if t.IsParchment() {
		headerCol = color.NRGBA{R: 0x8a, G: 0x7a, B: 0x66, A: 0xff}
		divCol    = color.NRGBA{R: 0xc0, G: 0xb4, B: 0xa0, A: 0xff}
	} else {
		headerCol = color.NRGBA{R: 0x5c, G: 0x4a, B: 0x38, A: 0xff}
		divCol    = color.NRGBA{R: 0x2e, G: 0x26, B: 0x1e, A: 0xff}
	}

	header := canvas.NewText("CHAPTERS", headerCol)
	header.TextSize = 10
	header.TextStyle = fyne.TextStyle{Bold: true}

	items := make([]fyne.CanvasObject, 0, len(story.Chapters))
	for i, ch := range story.Chapters {
		idx := i
		chapter := ch
		btn := widget.NewButton(fmt.Sprintf("%02d  %s", chapter.Number, chapter.Title), func() { onSelect(idx) })
		btn.Alignment = widget.ButtonAlignLeading
		items = append(items, btn)
	}

	divLine := canvas.NewLine(divCol)
	divLine.StrokeWidth = 1

	return container.NewBorder(
		container.NewVBox(container.NewPadded(header), divLine),
		nil, nil, nil,
		container.NewVScroll(container.NewVBox(items...)),
	)
}
