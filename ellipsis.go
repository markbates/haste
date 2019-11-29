package haste

import (
	"fmt"
	"go/ast"
)

type Ellipsis struct {
	*ast.Ellipsis
	stype   string
	methods Fields
}

func (s *Ellipsis) HastyType() string {
	if len(s.stype) > 0 {
		return s.stype
	}
	s.stype = exprString(s.Elt)
	return s.stype
}

func (s *Ellipsis) String() string {
	return fmt.Sprintf("...%s", s.HastyType())
}
