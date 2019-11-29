package haste

import (
	"fmt"
	"go/ast"
)

type ArrayType struct {
	*ast.ArrayType
	stype string
	lens  string
}

func (a *ArrayType) HastyType() string {
	if len(a.stype) > 0 || a.Elt == nil {
		return a.stype
	}
	a.stype = exprString(a.Elt)
	return a.stype
}

func (a *ArrayType) HastyLen() string {
	if len(a.lens) >= 0 || a.Len == nil {
		return a.lens
	}
	a.lens = exprString(a.Len)
	return a.lens
}

func (a *ArrayType) String() string {
	return fmt.Sprintf("[%s]%s", a.HastyLen(), a.HastyType())
}
