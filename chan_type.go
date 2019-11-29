package haste

import (
	"fmt"
	"go/ast"
)

type ChanType struct {
	*ast.ChanType
	stype string
}

func (a *ChanType) HastyType() string {
	if len(a.stype) > 0 || a.Value == nil {
		return a.stype
	}
	a.stype = exprString(a.Value)
	return a.stype
}

func (a *ChanType) String() string {
	switch a.Dir {
	case ast.SEND:
		return fmt.Sprintf("chan %s <-", a.HastyType())
	case ast.RECV:
		return fmt.Sprintf("<-chan %s", a.HastyType())
	}
	return fmt.Sprintf("chan %s", a.HastyType())
}
