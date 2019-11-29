package haste

import (
	"fmt"
	"go/ast"
)

type Field struct {
	*ast.Field
	names []string
	ftype string
}

type Fields []*Field

func (f *Field) HastyType() string {
	if len(f.ftype) > 0 {
		return f.ftype
	}
	f.ftype = exprString(f.Type)
	return f.ftype
}

func (f *Field) String() string {
	n := f.HastyName()
	if len(n) == 0 {
		return f.HastyType()
	}
	return fmt.Sprintf("%s %s", n, f.HastyType())
}

// []string{"ctx"}
// []string{"args"}
func (f *Field) HastyNames() []string {
	if len(f.names) > 0 {
		return f.names
	}
	for _, i := range f.Names {
		f.names = append(f.names, i.Name)
	}
	return f.names
}

// ctx
// args
func (f *Field) HastyName() string {
	names := f.HastyNames()
	if len(names) == 0 {
		return ""
	}
	return names[0]
}

func NewField(f *ast.Field) *Field {
	return &Field{
		Field: f,
	}
}
