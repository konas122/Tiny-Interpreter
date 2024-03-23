package object

import "fmt"

type ObjectType string

const (
	NULL_OBJ    = "NULL"
	BOOL_OBJ    = "BOOL"
	INTEGER_OBJ = "INTEGER"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// ================================================

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// ================================================

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOL_OBJ }

// ================================================

type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }
