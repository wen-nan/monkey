package object

import "fmt"

type ObjectType string

// constants for each ObjectType
const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer integers data type.
// Whenever we encounter an integer literal in the source code
// we first turn it into an ast.IntegerLiteral and then,
// when evaluating that AST node, we turn it into an object.Integer,
// saving the value inside our struct and passing around a reference to this struct.
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Boolean data type.
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null data type.
// It represents the absence of any value.
type Null struct {
}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
