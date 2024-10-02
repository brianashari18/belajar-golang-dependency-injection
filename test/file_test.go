package test

import (
	"github.com/stretchr/testify/assert"
	"golang-dependency-injection/simple"
	"testing"
)

func TestFileConnection(t *testing.T) {
	connection, cleanup := simple.InitializeConnection("Test.txt")
	assert.NotNil(t, connection)

	cleanup()
}
