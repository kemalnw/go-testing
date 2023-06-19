package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	word := HelloWorld("Kemal")

	assert.NotEmpty(t, word)
	assert.Equal(t, "Hello Kemal", word, "Result is must be 'Hello Kemal'")
}
