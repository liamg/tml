package tml

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseWithFormattingDisabled(t *testing.T) {
	DisableFormatting()
	defer EnableFormatting()
	output, err := Parse("plain <red>red <bold>bold red</bold></red> plain <green>green</green> plain")
	require.Nil(t, err)
	assert.Equal(t, "plain red bold red plain green plain", output)
}
