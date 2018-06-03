package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlDecode(t *testing.T) {
	LoadYaml("")
}

func TestCompleteYaml(t *testing.T) {
	input := "a: value1\n"
	data := LoadYaml(input)
	output := DumpYaml(data)
	assert.Equal(t, input, output)
}
