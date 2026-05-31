package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"foolstory/internal/data"
)

type StoryReader struct {
	story      *data.Story
	window     fyne.Window
	currentIdx int
	content    *fyne.Container
	scroll     *container.Scroll

	// nav widgets kept as fields so they update on chapter change
	prevBtn *widget.Button
	nextBtn *widget.Button
	counter *widget.Label

	OnChapterChange func(int)
}

func NewStoryReader(story *data.Story, w fyne.Window) *StoryReader {
	r := &StoryReader{story: story, window: w}

	r.content = container.NewVBox()
	r.scroll = container.NewVScroll(r.content)

	r.prevBtn = widget.NewButton("← Prev", func() { r.PrevChapter() })
	r.nextBtn = widget.NewButton("Next →", func() { r.NextChapter() })
	r.counter = widget.NewLabel("")
	r.counter.Alignment = fyne.TextAlignCenter

	r.renderChapter(0)
	return r
}

func (r *StoryReader) Container() fyne.CanvasObject {
	nav := container.NewPadded(
		container.NewBorder(nil, nil, r.prevBtn, r.nextBtn, r.counter),
	)
	return container.NewBorder(nil, nav, nil, nil, r.scroll)
}

func (r *StoryReader) GoToChapter(idx int) {
	if idx < 0 || idx >= len(r.story.Chapters) {
		return
	}
	r.renderChapter(idx)
}

func (r *StoryReader) NextChapter() { r.GoToChapter(r.currentIdx + 1) }
func (r *StoryReader) PrevChapter() { r.GoToChapter(r.currentIdx - 1) }

func (r *StoryReader) renderChapter(idx int) {
	ch := r.story.Chapters[idx]
	r.currentIdx = idx
	r.content.Objects = nil

	numLabel := canvas.NewText(
		fmt.Sprintf("Chapter %d", ch.Number),
		color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0x99},
	)
	numLabel.TextSize = 11
	numLabel.TextStyle = fyne.TextStyle{Bold: true}

	titleLabel := canvas.NewText(ch.Title, color.NRGBA{R: 0xe8, G: 0xdd, B: 0xd0, A: 0xff})
	titleLabel.TextSize = 22
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	div := canvas.NewLine(color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0x55})
	div.StrokeWidth = 1
	divBox := container.New(layout.NewStackLayout(), div)

	bodyLabel := widget.NewRichTextFromMarkdown(ch.Body)
	bodyLabel.Wrapping = fyne.TextWrapWord

	r.content.Add(container.NewPadded(container.NewVBox(
		container.NewPadded(numLabel),
		container.NewPadded(titleLabel),
		container.NewPadded(divBox),
		widget.NewSeparator(),
	)))
	r.content.Add(container.NewPadded(bodyLabel))

	if ch.Poem != "" {
		r.content.Add(widget.NewSeparator())
		r.content.Add(container.NewPadded(r.buildPoemCard(ch.Poem)))
	}

	r.content.Add(layout.NewSpacer())
	r.content.Refresh()
	r.scroll.ScrollToTop()

	// update nav state
	r.counter.SetText(fmt.Sprintf("%d / %d", idx+1, len(r.story.Chapters)))
	if idx == 0 {
		r.prevBtn.Disable()
	} else {
		r.prevBtn.Enable()
	}
	if idx == len(r.story.Chapters)-1 {
		r.nextBtn.Disable()
	} else {
		r.nextBtn.Enable()
	}

	if r.OnChapterChange != nil {
		r.OnChapterChange(idx)
	}
}

func (r *StoryReader) buildPoemCard(poem string) fyne.CanvasObject {
	quoteMark := canvas.NewText("❝", color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0xcc})
	quoteMark.TextSize = 28
	poemText := widget.NewRichTextFromMarkdown("*" + poem + "*")
	poemText.Wrapping = fyne.TextWrapWord
	return container.NewVBox(quoteMark, container.NewPadded(poemText))
}

func AppIcon() fyne.Resource { return nil }
