package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"

	"foolstory/internal/data"
	"foolstory/internal/ui"
)

func main() {
	a := app.NewWithID("com.praizethefool.foolstory")
	a.Settings().SetTheme(&ui.DarkTheme{})

	w := a.NewWindow("The Tale of the Fool")
	w.Resize(fyne.NewSize(780, 620))
	w.CenterOnScreen()

	story := data.LoadStory()
	reader := ui.NewStoryReader(story, w)

	sidebar := ui.NewChapterSidebar(story, func(idx int) {
		reader.GoToChapter(idx)
	})

	split := container.NewHSplit(sidebar, reader.Container())
	split.SetOffset(0.27)

	w.SetContent(container.NewBorder(
		ui.NewTitleBar(w),
		nil, nil, nil,
		split,
	))

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyRight, fyne.KeyDown:
			reader.NextChapter()
		case fyne.KeyLeft, fyne.KeyUp:
			reader.PrevChapter()
		}
	})

	w.ShowAndRun()
}

func init() { _ = theme.DefaultTheme() }
