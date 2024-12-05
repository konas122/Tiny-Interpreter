package object

import "fmt"

type LibFunction func(args ...Object) Object

type LibFunc struct {
	Fn LibFunction
}

func (lb *LibFunc) Type() ObjectType { return LIBFUNC_OBJ }
func (lb *LibFunc) Inspect() string  { return "lib function" }

// ==========================

var LibFuncs = []struct {
	Name    string
	LibFunc *LibFunc
}{
	{
		"len",
		&LibFunc{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}

				switch arg := args[0].(type) {
				case *String:
					return &Integer{Value: int64(len(arg.Value))}
				case *Array:
					return &Integer{Value: int64(len(arg.Elements))}
				default:
					return newError("argument to `len` not supported, got %s", args[0].Type())
				}
			},
		},
	},
	{
		"puts",
		&LibFunc{
			Fn: func(args ...Object) Object {
				for _, arg := range args {
					fmt.Println(arg.Inspect())
				}

				return nil
			},
		},
	},
	{
		"first",
		&LibFunc{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				if len(arr.Elements) > 0 {
					return arr.Elements[0]
				} else {
					return nil
				}
			},
		},
	},
	{
		"last",
		&LibFunc{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				length := len(arr.Elements)
				if length > 0 {
					return arr.Elements[length-1]
				} else {
					return nil
				}
			},
		},
	},
	{
		"rest",
		&LibFunc{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				length := len(arr.Elements)
				if length > 0 {
					newElements := make([]Object, length-1)
					copy(newElements, arr.Elements[1:length])
					return &Array{Elements: newElements}
				} else {
					return nil
				}
			},
		},
	},
	{
		"push",
		&LibFunc{
			Fn: func(args ...Object) Object {
				if len(args) != 2 {
					return newError("wrong number of arguments. got=%d, want=2", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				length := len(arr.Elements)

				newElements := make([]Object, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]

				return &Array{Elements: newElements}
			},
		},
	},
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func GetLibFuncByName(name string) *LibFunc {
	for _, def := range LibFuncs {
		if def.Name == name {
			return def.LibFunc
		}
	}

	return nil
}
