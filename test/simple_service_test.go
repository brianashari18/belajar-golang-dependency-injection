package test

import (
	"github.com/stretchr/testify/assert"
	"golang-dependency-injection/simple"
	"testing"
)

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}
