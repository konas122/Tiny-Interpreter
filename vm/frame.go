package vm

import (
	"Interpreter/code"
	"Interpreter/object"
)

type Frame struct {
	fn *object.CompiledFunction
	ip int
	bp int
}

func NewFrame(fn *object.CompiledFunction, bp int) *Frame {
	return &Frame{fn: fn, ip: -1, bp: bp}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
