package haste

import (
	"bytes"
	"fmt"
	"go/ast"
	"strings"
)

type FuncType struct {
	*ast.FuncType
	params  Fields
	returns Fields
}

func (f *FuncType) HastyParams() Fields {
	if len(f.params) > 0 || f.Params == nil {
		return f.params
	}
	for _, pf := range f.Params.List {
		f.params = append(f.params, NewField(pf))
	}
	return f.params
}

func (f *FuncType) HastyReturns() Fields {
	if len(f.returns) > 0 || f.Results == nil {
		return f.returns
	}
	for _, pf := range f.Results.List {
		f.returns = append(f.returns, NewField(pf))
	}
	return f.returns
}

func (f *FuncType) String() string {
	bb := &bytes.Buffer{}

	bb.WriteString("func(")

	var parms []string
	for _, p := range f.HastyParams() {
		parms = append(parms, p.String())
	}

	bb.WriteString(strings.Join(parms, ", "))
	bb.WriteString(fmt.Sprintf(")"))

	rets := f.HastyReturns()
	if len(rets) == 0 {
		return bb.String()
	}
	bb.WriteString(" ")
	if len(rets) == 1 {
		bb.WriteString(rets[0].String())
		return bb.String()
	}

	var rs []string
	for _, r := range rets {
		rs = append(rs, r.String())
	}
	bb.WriteString(fmt.Sprintf("(%s)", strings.Join(rs, ", ")))
	return bb.String()
}
