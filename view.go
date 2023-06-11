package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func createNumButtons(vm *CalcViewModel) *fyne.Container {
	return container.New(
		layout.NewGridLayout(3),
		widget.NewButton(strconv.Itoa(7), func() { vm.PushNum(7) }),
		widget.NewButton(strconv.Itoa(8), func() { vm.PushNum(8) }),
		widget.NewButton(strconv.Itoa(9), func() { vm.PushNum(9) }),
		widget.NewButton(strconv.Itoa(4), func() { vm.PushNum(4) }),
		widget.NewButton(strconv.Itoa(5), func() { vm.PushNum(5) }),
		widget.NewButton(strconv.Itoa(6), func() { vm.PushNum(6) }),
		widget.NewButton(strconv.Itoa(1), func() { vm.PushNum(1) }),
		widget.NewButton(strconv.Itoa(2), func() { vm.PushNum(2) }),
		widget.NewButton(strconv.Itoa(3), func() { vm.PushNum(3) }),
		widget.NewButton(strconv.Itoa(0), func() { vm.PushNum(0) }),
	)
}
func createCalcButtons(vm *CalcViewModel) *fyne.Container {
	return container.New(
		layout.NewGridLayout(1),
		widget.NewButton("CL", func() { vm.PushClear() }),
		widget.NewButton("/", func() { vm.PushCalc(Div) }),
		widget.NewButton("*", func() { vm.PushCalc(Mul) }),
		widget.NewButton("+", func() { vm.PushCalc(Add) }),
		widget.NewButton("-", func() { vm.PushCalc(Sub) }),
	)
}

func CreateCalcWindow(a fyne.App) fyne.Window {
	vm := NewCalcViewModel()
	w := a.NewWindow("Calc")
	w.Resize(fyne.NewSize(300, 200))
	w.SetFixedSize(true)

	l := widget.NewLabelWithData(vm.result)
	l.Alignment = fyne.TextAlignTrailing

	numButtons := createNumButtons(vm)
	calcButtons := createCalcButtons(vm)
	enterButton := widget.NewButton("Enter", vm.PushEnter)

	w.SetContent(container.New(
		layout.NewBorderLayout(l, enterButton, nil, calcButtons),
		l, enterButton, numButtons, calcButtons,
	))

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyEnter, fyne.KeyReturn:
			vm.PushEnter()
		case fyne.KeyDelete:
			vm.PushClear()
		}
	})

	w.Canvas().SetOnTypedRune(func(v rune) {
		n, err := strconv.Atoi(string(v))
		if err == nil {
			vm.PushNum(n)
			return
		}
		vm.PushCalc(command(v))
	})

	return w
}
