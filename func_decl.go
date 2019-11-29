package haste

import (
	"bytes"
	"fmt"
	"go/ast"
	"strings"
)

type FuncDecl struct {
	*ast.FuncDecl
	name      string
	receivers Fields
	params    Fields
	returns   Fields
	stringer  string
	mstringer string
}

type FuncDecls []*FuncDecl

func (fns FuncDecls) Find(s string) (*FuncDecl, error) {
	for _, fn := range fns {
		if fn.MatchString() == s {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("could not find matching function %q", s)
}

func (f *FuncDecl) MatchString() string {
	if len(f.mstringer) > 0 {
		return f.mstringer
	}
	bb := &bytes.Buffer{}

	bb.WriteString("func ")

	recvs := f.HastyReceivers()
	if len(recvs) > 0 {
		var lines []string
		for _, rec := range recvs {
			lines = append(lines, rec.HastyType())
		}
		bb.WriteString(fmt.Sprintf("(%s) ", strings.Join(lines, ", ")))
	}

	bb.WriteString(fmt.Sprintf("%s(", f.HastyName()))

	var parms []string
	for _, p := range f.HastyParams() {
		parms = append(parms, p.HastyType())
	}

	bb.WriteString(strings.Join(parms, ", "))
	bb.WriteString(fmt.Sprintf(")"))

	rets := f.HastyReturns()
	if len(rets) == 0 {
		return bb.String()
	}
	bb.WriteString(" ")
	if len(rets) == 1 {
		bb.WriteString(rets[0].HastyType())
		return bb.String()
	}

	var rs []string
	for _, r := range rets {
		rs = append(rs, r.HastyType())
	}
	bb.WriteString(fmt.Sprintf("(%s)", strings.Join(rs, ", ")))
	f.mstringer = bb.String()
	return f.mstringer
}

func (f *FuncDecl) String() string {
	if len(f.stringer) > 0 {
		return f.stringer
	}
	bb := &bytes.Buffer{}

	bb.WriteString("func ")

	recvs := f.HastyReceivers()
	if len(recvs) > 0 {
		var lines []string
		for _, rec := range recvs {
			lines = append(lines, rec.String())
		}
		bb.WriteString(fmt.Sprintf("(%s) ", strings.Join(lines, ", ")))
	}

	bb.WriteString(fmt.Sprintf("%s(", f.HastyName()))

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
	f.stringer = bb.String()
	return f.stringer
}

func (f *FuncDecl) HastyName() string {
	if len(f.name) != 0 || f.Name == nil {
		return f.name
	}

	i := f.Name
	if i == nil {
		return ""
	}
	f.name = i.Name
	return f.name
}

// []{context.Context, []string}
func (f *FuncDecl) HastyParams() Fields {
	if len(f.params) > 0 || f.Type == nil {
		return f.params
	}
	for _, pf := range f.Type.Params.List {
		f.params = append(f.params, NewField(pf))
	}
	return f.params
}

func (f *FuncDecl) HastyReturns() Fields {
	if len(f.returns) > 0 || f.Type == nil || f.Type.Results == nil {
		return f.returns
	}
	for _, pf := range f.Type.Results.List {
		f.returns = append(f.returns, NewField(pf))
	}
	return f.returns
}

func (f *FuncDecl) HastyReceivers() Fields {
	if len(f.receivers) > 0 || f.Recv == nil {
		return f.receivers
	}
	for _, pf := range f.Recv.List {
		f.receivers = append(f.receivers, NewField(pf))
	}
	return f.receivers
}

func NewFuncDecl(f *ast.FuncDecl) *FuncDecl {
	return &FuncDecl{
		FuncDecl: f,
	}
}
