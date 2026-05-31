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

func NewChapterSidebar(story *data.Story, onSelect func(int)) fyne.CanvasObject {
	header := canvas.NewText("CHAPTERS", color.NRGBA{R: 0x5c, G: 0x4a, B: 0x38, A: 0xff})
	header.TextSize = 10
	header.TextStyle = fyne.TextStyle{Bold: true}

	items := make([]fyne.CanvasObject, 0, len(story.Chapters))
	for i, ch := range story.Chapters {
		idx := i
		chapter := ch
		label := fmt.Sprintf("%02d  %s", chapter.Number, chapter.Title)
		btn := widget.NewButton(label, func() { onSelect(idx) })
		btn.Alignment = widget.ButtonAlignLeading
		items = append(items, btn)
	}

	divLine := canvas.NewLine(color.NRGBA{R: 0x2e, G: 0x26, B: 0x1e, A: 0xff})
	divLine.StrokeWidth = 1

	return container.NewBorder(
		container.NewVBox(container.NewPadded(header), divLine),
		nil, nil, nil,
		container.NewVScroll(container.NewVBox(items...)),
	)
}
