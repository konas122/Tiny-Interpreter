package evaluator

import (
	"Interpreter/object"
)

var LibFuncs = map[string]*object.LibFunc{
	"len":   object.GetLibFuncByName("len"),
	"first": object.GetLibFuncByName("first"),
	"last":  object.GetLibFuncByName("last"),
	"rest":  object.GetLibFuncByName("rest"),
	"push":  object.GetLibFuncByName("push"),
	"puts":  object.GetLibFuncByName("puts"),
}
