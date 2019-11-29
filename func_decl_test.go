package haste

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FuncDecls(t *testing.T) {
	ip := "github.com/markbates/haste/internal/ref"

	r := require.New(t)

	h, err := New(ip)
	r.NoError(err)

	funcs := h.Funcs()
	r.Len(funcs, 1)

	fn, err := funcs.Find("func Jim(context.Context, []string) error")
	r.NoError(err)

	r.Equal("func Jim(ctx context.Context, args []string) error", fn.String())
	r.Equal("func Jim(context.Context, []string) error", fn.MatchString())

	_, err = funcs.Find("i don't exist")
	r.Error(err)
}
