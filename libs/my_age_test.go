package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAge(t *testing.T) {
	result := MyAge(20)
	assert.Equal(t, "my age is 20", result)
}
