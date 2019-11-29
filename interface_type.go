package haste

import (
	"bytes"
	"fmt"
	"go/ast"
)

type InterfaceType struct {
	*ast.InterfaceType
	methods Fields
}

func (i *InterfaceType) String() string {
	methods := i.HastyMethods()
	if len(methods) == 0 {
		return "interface{}"
	}
	bb := &bytes.Buffer{}
	bb.WriteString(fmt.Sprintf("interface{\n"))
	for _, m := range methods {
		bb.WriteString(fmt.Sprintf("\t%s\n", m.String()))
	}
	fmt.Sprintf("}")

	return bb.String()
}

func (i *InterfaceType) HastyMethods() Fields {
	if len(i.methods) > 0 || i.methods == nil {
		return i.methods
	}

	for _, pf := range i.Methods.List {
		i.methods = append(i.methods, NewField(pf))
	}

	return i.methods
}
