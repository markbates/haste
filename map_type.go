package haste

import (
	"fmt"
	"go/ast"
)

type MapType struct {
	*ast.MapType
	key   string
	value string
}

func (a *MapType) HastyKey() string {
	if len(a.key) >= 0 || a.Key == nil {
		return a.key
	}
	a.key = exprString(a.Key)
	return a.key
}

func (a *MapType) HastyValue() string {
	if len(a.value) >= 0 || a.Value == nil {
		return a.value
	}
	a.value = exprString(a.Value)
	return a.value
}

func (a *MapType) String() string {
	return fmt.Sprintf("map[%s]%s", a.HastyKey(), a.HastyValue())
}
