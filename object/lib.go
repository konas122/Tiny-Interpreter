package object

type LibFunction func(args ...Object) Object

type LibFunc struct {
	Fn LibFunction
}

func (lb *LibFunc) Type() ObjectType { return LIBFUNC_OBJ }
func (lb *LibFunc) Inspect() string  { return "lib function" }
