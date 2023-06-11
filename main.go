package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := CreateCalcWindow(a)
	w.ShowAndRun()
}
