package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// FlipDigit represents a single character in the clock
type FlipDigit struct {
	widget.BaseWidget
	char string
	text *canvas.Text
}

func NewFlipDigit(char string) *FlipDigit {
	fd := &FlipDigit{char: char}
	fd.text = canvas.NewText(char, theme.ForegroundColor())
	fd.text.TextSize = 48
	fd.text.Alignment = fyne.TextAlignCenter
	fd.ExtendBaseWidget(fd)
	return fd
}

func (f *FlipDigit) SetDigit(c string) {
	if f.char != c {
		f.char = c
		f.text.Text = c
		canvas.Refresh(f)
	}
}

func (f *FlipDigit) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewMax(f.text))
}

func updateTime(digits []*FlipDigit) {
	now := time.Now().Format("15:04:05")
	for i, c := range now {
		digits[i].SetDigit(string(c))
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Flip Clock")

	timeStr := time.Now().Format("15:04:05")
	var digits []*FlipDigit
	for _, c := range timeStr {
		d := NewFlipDigit(string(c))
		digits = append(digits, d)
	}

	row := container.NewHBox()
	for _, d := range digits {
		row.Add(d)
	}

	updateTime(digits)

	w.SetContent(container.NewCenter(row))
	go func() {
		for range time.Tick(time.Second) {
			updateTime(digits)
		}
	}()

	w.Resize(fyne.NewSize(400, 100))
	w.ShowAndRun()
}