package haste

import (
	"fmt"
	"go/ast"
)

type SelectorExpr struct {
	*ast.SelectorExpr
	sel   string
	stype string
}

func (s *SelectorExpr) HastySelector() string {
	if len(s.sel) > 0 {
		return s.sel
	}
	s.sel = exprString(s.X)
	return s.sel

}

func (s *SelectorExpr) HastyType() string {
	if len(s.stype) > 0 {
		return s.stype
	}
	s.stype = exprString(s.Sel)
	return s.stype
}

func (s *SelectorExpr) String() string {
	return fmt.Sprintf("%s.%s", s.HastySelector(), s.HastyType())
}
