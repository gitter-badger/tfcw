package cmd

import (
	"fmt"
	"testing"

	"github.com/mvisonneau/go-helpers/assert"
)

func TestExit(t *testing.T) {
	err := exit(20, fmt.Errorf("test"))
	assert.Equal(t, err.Error(), "")
	assert.Equal(t, err.ExitCode(), 20)
}
