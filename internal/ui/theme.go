package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// ThemeID identifies which theme is active
type ThemeID int

const (
	ThemeParchment ThemeID = iota // sepia ink (default, like the screenshot)
	ThemeWarm                     // dark warm brown (original)
)

// AppTheme is the switchable theme
type AppTheme struct {
	ID ThemeID
}

var _ fyne.Theme = (*AppTheme)(nil)

// parchment palette (#e8ddd0 bg, dark ink fg)
var parchment = map[fyne.ThemeColorName]color.Color{
	theme.ColorNameBackground:       color.NRGBA{R: 0xe8, G: 0xdf, B: 0xd0, A: 0xff}, // warm paper
	theme.ColorNameForeground:       color.NRGBA{R: 0x1a, G: 0x15, B: 0x10, A: 0xff}, // dark ink
	theme.ColorNamePrimary:          color.NRGBA{R: 0x3a, G: 0x2e, B: 0x22, A: 0xff}, // deep ink
	theme.ColorNameHover:            color.NRGBA{R: 0xd8, G: 0xcf, B: 0xbf, A: 0xff}, // slightly darker paper
	theme.ColorNameButton:           color.NRGBA{R: 0xd8, G: 0xcf, B: 0xbf, A: 0xff},
	theme.ColorNameDisabledButton:   color.NRGBA{R: 0xe0, G: 0xd8, B: 0xcb, A: 0xff},
	theme.ColorNameInputBackground:  color.NRGBA{R: 0xf0, G: 0xe8, B: 0xd8, A: 0xff},
	theme.ColorNamePlaceHolder:      color.NRGBA{R: 0x8a, G: 0x7a, B: 0x66, A: 0xff},
	theme.ColorNameScrollBar:        color.NRGBA{R: 0x5c, G: 0x4a, B: 0x38, A: 0x60},
	theme.ColorNameShadow:           color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x22},
	theme.ColorNameSeparator:        color.NRGBA{R: 0xc0, G: 0xb4, B: 0xa0, A: 0xff},
	theme.ColorNameSelection:        color.NRGBA{R: 0x5c, G: 0x4a, B: 0x38, A: 0x33},
	theme.ColorNameDisabled:         color.NRGBA{R: 0xa0, G: 0x90, B: 0x7a, A: 0xff},
	theme.ColorNameFocus:            color.NRGBA{R: 0x3a, G: 0x2e, B: 0x22, A: 0xff},
	theme.ColorNameHeaderBackground: color.NRGBA{R: 0xdc, G: 0xd3, B: 0xc3, A: 0xff},
	theme.ColorNameMenuBackground:   color.NRGBA{R: 0xf0, G: 0xe8, B: 0xd8, A: 0xff},
}

// warm dark palette (original brown)
var warm = map[fyne.ThemeColorName]color.Color{
	theme.ColorNameBackground:       color.NRGBA{R: 0x1a, G: 0x15, B: 0x10, A: 0xff},
	theme.ColorNameForeground:       color.NRGBA{R: 0xe8, G: 0xdd, B: 0xd0, A: 0xff},
	theme.ColorNamePrimary:          color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0xff},
	theme.ColorNameHover:            color.NRGBA{R: 0x2a, G: 0x22, B: 0x1a, A: 0xff},
	theme.ColorNameButton:           color.NRGBA{R: 0x22, G: 0x1c, B: 0x16, A: 0xff},
	theme.ColorNameDisabledButton:   color.NRGBA{R: 0x1a, G: 0x15, B: 0x10, A: 0xff},
	theme.ColorNameInputBackground:  color.NRGBA{R: 0x22, G: 0x1c, B: 0x16, A: 0xff},
	theme.ColorNamePlaceHolder:      color.NRGBA{R: 0x5c, G: 0x4a, B: 0x38, A: 0xff},
	theme.ColorNameScrollBar:        color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0x60},
	theme.ColorNameShadow:           color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x66},
	theme.ColorNameSeparator:        color.NRGBA{R: 0x2e, G: 0x26, B: 0x1e, A: 0xff},
	theme.ColorNameSelection:        color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0x44},
	theme.ColorNameDisabled:         color.NRGBA{R: 0x5c, G: 0x4a, B: 0x38, A: 0xff},
	theme.ColorNameFocus:            color.NRGBA{R: 0xb8, G: 0xa8, B: 0x98, A: 0xff},
	theme.ColorNameHeaderBackground: color.NRGBA{R: 0x17, G: 0x12, B: 0x0e, A: 0xff},
	theme.ColorNameMenuBackground:   color.NRGBA{R: 0x22, G: 0x1c, B: 0x16, A: 0xff},
}

func (t *AppTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	var palette map[fyne.ThemeColorName]color.Color
	if t.ID == ThemeParchment {
		palette = parchment
	} else {
		palette = warm
	}
	if c, ok := palette[name]; ok {
		return c
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t *AppTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *AppTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *AppTheme) Size(name fyne.ThemeSizeName) float32 {
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

// IsParchment returns true if the light parchment theme is active
func (t *AppTheme) IsParchment() bool { return t.ID == ThemeParchment }
