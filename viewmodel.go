package main

import (
	"strconv"

	"fyne.io/fyne/v2/data/binding"
)

type command rune

const (
	None command = ' '
	Add  command = '+'
	Sub  command = '-'
	Mul  command = '*'
	Div  command = '/'
)

type CalcViewModel struct {
	mem             int
	cmd             command
	shouldClrResult bool
	result          binding.String
}

func (vm *CalcViewModel) calc(n int) {
	switch vm.cmd {
	case Add:
		vm.mem += n
	case Sub:
		vm.mem -= n
	case Mul:
		vm.mem *= n
	case Div:
		vm.mem /= n
	default:
		vm.mem = n
	}

	vm.result.Set(strconv.Itoa(vm.mem))
	vm.shouldClrResult = true
}

func (vm *CalcViewModel) PushNum(v int) {
	s, err := vm.result.Get()

	if vm.shouldClrResult || err != nil {
		s = "0"
		vm.shouldClrResult = false
	}
	s += strconv.Itoa(v)
	n, err := strconv.Atoi(s)
	if err == nil {
		vm.result.Set(strconv.Itoa(n))
	}
}

func (vm *CalcViewModel) PushCalc(c command) {
	v, err := vm.result.Get()
	if err != nil {
		return
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return
	}

	// 前の操作のコマンドに従って計算
	vm.calc(n)
	// 押されたボタンを次の操作のコマンドに代入
	vm.cmd = c
}

func (vm *CalcViewModel) PushClear() {
	vm.result.Set("0")
	vm.mem = 0
	vm.shouldClrResult = false
	vm.cmd = None
}

func (vm *CalcViewModel) PushEnter() {
	if vm.shouldClrResult {
		vm.shouldClrResult = false
		return
	}

	v, err := vm.result.Get()
	if err != nil {
		return
	}

	n, err := strconv.Atoi(v)
	if err != nil {
		return
	}
	vm.calc(n)
	vm.cmd = None
}

func NewCalcViewModel() *CalcViewModel {
	return &CalcViewModel{
		mem:             0,
		cmd:             None,
		shouldClrResult: false,
		result:          binding.NewString(),
	}
}
