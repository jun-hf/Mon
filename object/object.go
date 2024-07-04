package object

import (
	"bytes"
	"fmt"
	"strings"

	"mon.example/ast"
)

const (
	INTEGER_OBJECT      = "INTEGER"
	BOOLEAN_OBJECT      = "BOOLEAN"
	NULL_OBJECT         = "NULL"
	RETURN_VALUE_OBJECT = "RETURN_VALUE"
	ERROR_OBJECT        = "ERROR"
	FUNCTION_OBJECT = "FUNCTION"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer represent the interger value in the language when evaluating the ast
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return ObjectType(INTEGER_OBJECT) }

// Boolean represent the boolean value in the language when evaluating the ast
type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return ObjectType(BOOLEAN_OBJECT) }

// Null represetn the NULL value in the language, it means that there is no value
type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return ObjectType(NULL_OBJECT) }

// ReturnValue represent the value that is being returned
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() ObjectType { return ObjectType(RETURN_VALUE_OBJECT) }

// Error represent the error when evaluating the source text
type Error struct {
	Message string
}

func (e *Error) Inspect() string { return "ERROR " + e.Message }
func (e *Error) Type() ObjectType { return ObjectType(ERROR_OBJECT) }

type Function struct {
	Parameters []*ast.Identifier
	Body *ast.BlockStatement
	Env *Environment
}

func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
func (f *Function) Type() ObjectType { return FUNCTION_OBJECT }
