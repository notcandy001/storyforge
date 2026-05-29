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

// StoryReader is the main reading view.
type StoryReader struct {
	story          *data.Story
	window         fyne.Window
	currentIdx     int
	content        *fyne.Container
	scroll         *container.Scroll
	OnChapterChange func(int)
}

// NewStoryReader constructs the reader.
func NewStoryReader(story *data.Story, w fyne.Window) *StoryReader {
	r := &StoryReader{
		story:  story,
		window: w,
	}

	r.content = container.NewVBox()
	r.scroll = container.NewVScroll(r.content)

	r.renderChapter(0)
	return r
}

// Container returns the scrollable reading area.
func (r *StoryReader) Container() fyne.CanvasObject {
	nav := r.buildNav()
	return container.NewBorder(nil, nav, nil, nil, r.scroll)
}

// GoToChapter navigates to a specific chapter index.
func (r *StoryReader) GoToChapter(idx int) {
	if idx < 0 || idx >= len(r.story.Chapters) {
		return
	}
	r.renderChapter(idx)
}

// NextChapter advances to the next chapter.
func (r *StoryReader) NextChapter() {
	r.GoToChapter(r.currentIdx + 1)
}

// PrevChapter goes back one chapter.
func (r *StoryReader) PrevChapter() {
	r.GoToChapter(r.currentIdx - 1)
}

func (r *StoryReader) renderChapter(idx int) {
	ch := r.story.Chapters[idx]
	r.currentIdx = idx

	r.content.Objects = nil

	// Chapter number badge
	numLabel := canvas.NewText(fmt.Sprintf("Chapter %d", ch.Number), color.NRGBA{R: 0xc4, G: 0x9a, B: 0x6c, A: 0x99})
	numLabel.TextSize = 11
	numLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Chapter title
	titleLabel := canvas.NewText(ch.Title, color.NRGBA{R: 0xe8, G: 0xe0, B: 0xd5, A: 0xff})
	titleLabel.TextSize = 22
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Gold divider
	div := canvas.NewLine(color.NRGBA{R: 0xc4, G: 0x9a, B: 0x6c, A: 0x55})
	div.StrokeWidth = 1
	divBox := container.New(layout.NewStackLayout(), div)

	// Chapter body text
	bodyLabel := widget.NewRichTextFromMarkdown(ch.Body)
	bodyLabel.Wrapping = fyne.TextWrapWord

	header := container.NewVBox(
		container.NewPadded(numLabel),
		container.NewPadded(titleLabel),
		container.NewPadded(divBox),
		widget.NewSeparator(),
	)

	r.content.Add(container.NewPadded(header))
	r.content.Add(container.NewPadded(bodyLabel))

	// Optional poem / quote
	if ch.Poem != "" {
		r.content.Add(widget.NewSeparator())
		poemCard := r.buildPoemCard(ch.Poem)
		r.content.Add(container.NewPadded(poemCard))
	}

	// Spacer before nav
	r.content.Add(layout.NewSpacer())

	r.content.Refresh()
	r.scroll.ScrollToTop()

	if r.OnChapterChange != nil {
		r.OnChapterChange(idx)
	}
}

func (r *StoryReader) buildPoemCard(poem string) fyne.CanvasObject {
	quoteMark := canvas.NewText("❝", color.NRGBA{R: 0xc4, G: 0x9a, B: 0x6c, A: 0xcc})
	quoteMark.TextSize = 28

	poemText := widget.NewRichTextFromMarkdown("*" + poem + "*")
	poemText.Wrapping = fyne.TextWrapWord

	return container.NewVBox(
		quoteMark,
		container.NewPadded(poemText),
	)
}

func (r *StoryReader) buildNav() fyne.CanvasObject {
	prevBtn := widget.NewButton("← Prev", r.PrevChapter)
	nextBtn := widget.NewButton("Next →", r.NextChapter)

	// Chapter counter label
	counter := widget.NewLabel(fmt.Sprintf("%d / %d", r.currentIdx+1, len(r.story.Chapters)))
	counter.Alignment = fyne.TextAlignCenter

	// Disable buttons at boundaries
	if r.currentIdx == 0 {
		prevBtn.Disable()
	}
	if r.currentIdx == len(r.story.Chapters)-1 {
		nextBtn.Disable()
	}

	navRow := container.NewBorder(nil, nil, prevBtn, nextBtn, counter)
	return container.NewPadded(navRow)
}

// AppIcon returns a simple generated icon (nil = use default).
func AppIcon() fyne.Resource {
	return nil
}
