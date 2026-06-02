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

	appTheme := &ui.AppTheme{ID: ui.ThemeParchment}
	a.Settings().SetTheme(appTheme)

	w := a.NewWindow("The Tale of the Fool")
	w.Resize(fyne.NewSize(820, 640))
	w.CenterOnScreen()

	story := data.LoadStory()

	// rebuild is called whenever theme switches — recreates all widgets
	// so colors reflect the new theme
	var rebuild func()
	rebuild = func() {
		reader := ui.NewStoryReader(story, w, appTheme)

		sidebar := ui.NewChapterSidebar(story, appTheme, func(idx int) {
			reader.GoToChapter(idx)
		})

		split := container.NewHSplit(sidebar, reader.Container())
		split.SetOffset(0.27)

		titleBar := ui.NewTitleBar(w, appTheme, func() {
			// toggle theme
			if appTheme.ID == ui.ThemeParchment {
				appTheme.ID = ui.ThemeWarm
			} else {
				appTheme.ID = ui.ThemeParchment
			}
			a.Settings().SetTheme(appTheme)
			rebuild()
		})

		w.SetContent(container.NewBorder(titleBar, nil, nil, nil, split))
		w.Content().Refresh()
	}

	rebuild()

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		// key nav works through the reader captured in rebuild closure
		// re-grab from content — simpler to just refresh rebuild on key
	})

	w.ShowAndRun()
}

func init() { _ = theme.DefaultTheme() }
