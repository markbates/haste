package haste

import "go/ast"

type Ident struct {
	*ast.Ident
}

func NewIdent(name string) *Ident {
	i := ast.NewIdent(name)
	return &Ident{
		Ident: i,
	}
}
