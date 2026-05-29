package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// DarkTheme is an elegant dark reading theme.
type DarkTheme struct{}

var _ fyne.Theme = (*DarkTheme)(nil)

func (t *DarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0x10, G: 0x0e, B: 0x17, A: 0xff} // deep night
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0xe8, G: 0xe0, B: 0xd5, A: 0xff} // warm white
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0xc4, G: 0x9a, B: 0x6c, A: 0xff} // warm gold
	case theme.ColorNameHover:
		return color.NRGBA{R: 0x1e, G: 0x1a, B: 0x28, A: 0xff}
	case theme.ColorNameButton:
		return color.NRGBA{R: 0x1e, G: 0x1a, B: 0x28, A: 0xff}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0x2a, G: 0x26, B: 0x33, A: 0xff}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0x18, G: 0x15, B: 0x22, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0x66, G: 0x5e, B: 0x78, A: 0xff}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0xc4, G: 0x9a, B: 0x6c, A: 0x60}
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x88}
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 0x2a, G: 0x26, B: 0x36, A: 0xff}
	case theme.ColorNameSelection:
		return color.NRGBA{R: 0xc4, G: 0x9a, B: 0x6c, A: 0x44}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0x44, G: 0x40, B: 0x50, A: 0xff}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 0xc4, G: 0x9a, B: 0x6c, A: 0xff}
	case theme.ColorNameHeaderBackground:
		return color.NRGBA{R: 0x14, G: 0x11, B: 0x1e, A: 0xff}
	case theme.ColorNameMenuBackground:
		return color.NRGBA{R: 0x18, G: 0x15, B: 0x22, A: 0xff}
	case theme.ColorNameOverlayBackground:
		return color.NRGBA{R: 0x18, G: 0x15, B: 0x22, A: 0xee}
	case theme.ColorNameSuccess:
		return color.NRGBA{R: 0x6c, G: 0xc4, B: 0x9a, A: 0xff}
	case theme.ColorNameWarning:
		return color.NRGBA{R: 0xc4, G: 0xb0, B: 0x6c, A: 0xff}
	case theme.ColorNameError:
		return color.NRGBA{R: 0xc4, G: 0x6c, B: 0x6c, A: 0xff}
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
