package haste

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/packages"
)

type Haste struct {
	pkgs  []*packages.Package
	funcs FuncDecls
}

func (h *Haste) Funcs() FuncDecls {
	if len(h.funcs) > 0 {
		return h.funcs
	}
	for _, p := range h.pkgs {
		for _, f := range p.Syntax {
			for _, d := range f.Decls {
				fn, ok := d.(*ast.FuncDecl)
				if !ok {
					continue
				}
				h.funcs = append(h.funcs, NewFuncDecl(fn))
			}
		}

	}

	return h.funcs
}

func New(s ...string) (*Haste, error) {
	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax}
	pkgs, err := packages.Load(cfg, s...)
	if err != nil {
		return nil, err
	}
	h := &Haste{
		pkgs: pkgs,
	}
	return h, nil
}

func exprString(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.SelectorExpr:
		s := &SelectorExpr{
			SelectorExpr: t,
		}
		return s.String()
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		s := &StarExpr{
			StarExpr: t,
		}
		return s.String()
	case *ast.Ellipsis:
		e := &Ellipsis{
			Ellipsis: t,
		}
		return e.String()
	case *ast.InterfaceType:
		i := &InterfaceType{
			InterfaceType: t,
		}
		return i.String()
	case *ast.ArrayType:
		a := &ArrayType{
			ArrayType: t,
		}
		return a.String()
	case *ast.MapType:
		m := &MapType{
			MapType: t,
		}
		return m.String()
	case *ast.ChanType:
		c := &ChanType{
			ChanType: t,
		}
		return c.String()
	case *ast.FuncType:
		f := &FuncType{
			FuncType: t,
		}
		return f.String()
	default:
		x := fmt.Sprintf("TYPE: %T\n", e)
		panic(x)
	}
	return ""
}
