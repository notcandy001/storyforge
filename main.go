package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"foolstory/internal/data"
	"foolstory/internal/ui"
)

func main() {
	a := app.NewWithID("com.praizethefool.foolstory")
	a.Settings().SetTheme(&ui.DarkTheme{})

	w := a.NewWindow("The Tale of the Fool")
	w.Resize(fyne.NewSize(780, 620))
	w.SetFixedSize(false)
	w.CenterOnScreen()

	story := data.LoadStory()
	reader := ui.NewStoryReader(story, w)

	// Progress bar at the bottom
	progress := widget.NewProgressBar()
	progress.Min = 0
	progress.Max = float64(len(story.Chapters) - 1)
	progress.Value = 0

	reader.OnChapterChange = func(idx int) {
		progress.SetValue(float64(idx))
	}

	statusBar := container.NewVBox(
		widget.NewSeparator(),
		container.NewPadded(progress),
	)

	// Sidebar: chapter list
	sidebar := ui.NewChapterSidebar(story, func(idx int) {
		reader.GoToChapter(idx)
	})

	split := container.NewHSplit(sidebar, reader.Container())
	split.SetOffset(0.27)

	main := container.NewBorder(
		ui.NewTitleBar(w),
		statusBar,
		nil,
		nil,
		split,
	)

	w.SetContent(main)

	icon := ui.AppIcon()
	if icon != nil {
		w.SetIcon(icon)
		a.SetIcon(icon)
	}

	// Keyboard nav
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

func init() {
	_ = theme.DefaultTheme()
}
