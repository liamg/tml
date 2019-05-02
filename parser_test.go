package tml

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {

	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "no-tags",
			input:  "plain",
			output: "\x1b[0mplain\x1b[0m",
		},
		{
			name:   "basic",
			input:  "plain <red>red</red> plain",
			output: "\x1b[0mplain \x1b[31mred\x1b[39m plain\x1b[0m",
		},
		{
			name:   "nesting",
			input:  "plain <red>red <bold>bold red</bold></red> plain <green>green</green> plain",
			output: "\x1b[0mplain \x1b[31mred \x1b[1mbold red\x1b[0m\x1b[31m\x1b[39m plain \x1b[32mgreen\x1b[39m plain\x1b[0m",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			buffer := bytes.NewBufferString("")
			parser := NewParser(buffer)
			err := parser.Parse(strings.NewReader(testCase.input))
			require.Nil(t, err)
			assert.Equal(t, testCase.output, buffer.String())
		})
	}

}
