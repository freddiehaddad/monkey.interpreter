// Monkey language object interface
package object

import "fmt"

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	NULL_OBJ    = "NULL"
	BOOLEAN_OBJ = "BOOLEAN"
	INTEGER_OBJ = "INTEGER"
)

// Null
type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

// Booleans
type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Integers
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
