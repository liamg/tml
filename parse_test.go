package tml

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	output, err := Parse("plain <red>red <bold>bold red</bold></red> plain <green>green</green> plain")
	require.Nil(t, err)
	assert.Equal(t, "\x1b[0mplain \x1b[31mred \x1b[1mbold red\x1b[0m\x1b[31m\x1b[39m plain \x1b[32mgreen\x1b[39m plain\x1b[0m", output)
}
