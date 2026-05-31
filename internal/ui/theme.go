package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// palette:
// bg      #1a1510
// fg      #e8ddd0
// dim     #5c4a38
// border  #b8a898

type DarkTheme struct{}

var _ fyne.Theme = (*DarkTheme)(nil)

var (
	bg     = color.NRGBA{R: 0x1a, G: 0x15, B: 0x10, A: 0xff}
	fg     = color.NRGBA{R: 0xe8, G: 0xdd, B: 0xd0, A: 0xff}
	dim    = color.NRGBA{R: 0x5c, G: 0x4a, B: 0x38, A: 0xff}
	border = color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0xff}
	bgL1   = color.NRGBA{R: 0x22, G: 0x1c, B: 0x16, A: 0xff} // slightly lighter bg
	bgL2   = color.NRGBA{R: 0x2a, G: 0x22, B: 0x1a, A: 0xff} // hover/button
	divCol = color.NRGBA{R: 0x2e, G: 0x26, B: 0x1e, A: 0xff}
)

func (t *DarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return bg
	case theme.ColorNameForeground:
		return fg
	case theme.ColorNamePrimary:
		return border
	case theme.ColorNameHover:
		return bgL2
	case theme.ColorNameButton:
		return bgL1
	case theme.ColorNameDisabledButton:
		return bg
	case theme.ColorNameInputBackground:
		return bgL1
	case theme.ColorNamePlaceHolder:
		return dim
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0x60}
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x66}
	case theme.ColorNameSeparator:
		return divCol
	case theme.ColorNameSelection:
		return color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0x44}
	case theme.ColorNameDisabled:
		return dim
	case theme.ColorNameFocus:
		return border
	case theme.ColorNameHeaderBackground:
		return bg
	case theme.ColorNameMenuBackground:
		return bgL1
	case theme.ColorNameOverlayBackground:
		return bgL1
	case theme.ColorNameSuccess:
		return color.NRGBA{R: 0x98, G: 0xb8, B: 0xa0, A: 0xff}
	case theme.ColorNameWarning:
		return color.NRGBA{R: 0xb8, G: 0xaa, B: 0x80, A: 0xff}
	case theme.ColorNameError:
		return color.NRGBA{R: 0xb8, G: 0x80, B: 0x80, A: 0xff}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t *DarkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *DarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *DarkTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 14
	case theme.SizeNameHeadingText:
		return 22
	case theme.SizeNameSubHeadingText:
		return 17
	case theme.SizeNamePadding:
		return 6
	case theme.SizeNameInnerPadding:
		return 8
	case theme.SizeNameLineSpacing:
		return 5
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameScrollBar:
		return 10
	}
	return theme.DefaultTheme().Size(name)
}
