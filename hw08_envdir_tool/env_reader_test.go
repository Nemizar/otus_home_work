package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("check equal", func(t *testing.T) {
		e := Environment{
			"BAR":   {Value: "bar", NeedRemove: false},
			"EMPTY": {Value: "", NeedRemove: false},
			"FOO":   {Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": {Value: "\"hello\"", NeedRemove: false},
			"UNSET": {Value: "", NeedRemove: true},
		}

		result, _ := ReadDir("./testdata/env")

		require.Equal(t, e, result)
	})

	t.Run("check not equal", func(t *testing.T) {
		e := Environment{"MY_ENV": {Value: "MY_VALUE", NeedRemove: false}}

		result, _ := ReadDir("./testdata/env")

		require.NotEqual(t, e, result)
	})
}
