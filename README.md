# The Tale of the Fool  Story Reader

A desktop GUI story reader built with Go + Fyne.

## Features

- **Chapter sidebar** — jump to any of the 16 chapters instantly
- **Reading panel** — clean dark theme, warm gold accents, poem/quote cards
- **Keyboard navigation** — `→` / `↓` next chapter, `←` / `↑` previous
- **Progress bar** — shows how far through the story you are
- **Scroll memory** — each chapter scrolls back to top on navigation

## sneak peek
<div align="center">
  <img src="https://github.com/notcandy001/storyforge/blob/main/showcase2.png" width="100%" />
  <img src="https://github.com/notcandy001/storyforge/blob/main/showcase.png" width="100%" />
</div>
## Setup

### Prerequisites

**Linux:**
```bash
sudo apt install gcc libgl1-mesa-dev xorg-dev
```

**macOS:**
```bash
xcode-select --install
```

**Windows:** Install TDM-GCC or MinGW-w64

### Run

```bash
git clone https://github.com/notcandy001/storyforge 
cd foolstory
go mod tidy
go run .
```

### Build binary

```bash
go build -o foolstory .
./foolstory
```

### Build with Fyne packager (optional, adds icon + metadata)

```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os linux -name "The Tale of the Fool"
```

## Project Structure

```
foolstory/
├── main.go                   # App entry, window setup, keybindings
├── internal/
│   ├── data/
│   │   └── story.go          # All 16 chapters hardcoded
│   └── ui/
│       ├── theme.go          # Custom dark Fyne theme (deep night + warm gold)
│       ├── titlebar.go       # Top bar with title/date/tag
│       ├── sidebar.go        # Left chapter list
│       └── reader.go         # Main reading panel + nav buttons
└── README.md
```

## Theme

| Color | Hex | Usage |
|---|---|---|
| Background | `#100e17` | Deep night |
| Foreground | `#e8e0d5` | Warm off-white |
| Accent | `#c49a6c` | Gold — chapter numbers, poems, nav |
| Muted | `#665e78` | Subtle text, dividers |
