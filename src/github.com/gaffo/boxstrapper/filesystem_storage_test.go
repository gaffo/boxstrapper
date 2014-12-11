package boxstrapper_test

import (
	"testing"
	. "github.com/gaffo/boxstrapper"
  	"github.com/stretchr/testify/assert"
)

func TestReadEmptyRepo(t *testing.T) {
	assert := assert.New(t)
	storage := NewFilesystemStorage("/nonexistent")

	data, err := storage.ReadPackages()
	assert.NotNil(err)
	assert.Equal("", data)
}