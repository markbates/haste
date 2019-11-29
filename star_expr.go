package haste

import (
	"fmt"
	"go/ast"
)

func (s *StarExpr) HastyType() string {
	if len(s.stype) > 0 {
		return s.stype
	}
	s.stype = exprString(s.X)
	return s.stype
}

type StarExpr struct {
	*ast.StarExpr
	stype string
	name  string
}

func (s *StarExpr) HastyName() string {
	if len(s.name) > 0 {
		return s.name
	}
	s.name = exprString(s.X)
	return s.name
}

func (s *StarExpr) String() string {
	return fmt.Sprintf("*%s", s.HastyName())
}
