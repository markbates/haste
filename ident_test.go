package haste

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Ident(t *testing.T) {
	r := require.New(t)

	i := NewIdent("foo")
	r.Equal("foo", i.Name)
}
