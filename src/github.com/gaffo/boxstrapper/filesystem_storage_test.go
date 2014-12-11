package boxstrapper_test

import (
	"testing"
	. "github.com/gaffo/boxstrapper"
  	"github.com/stretchr/testify/assert"
  	"io/ioutil"
  	"os"
)

func TestReadPackages_EmptyRepo(t *testing.T) {
	assert := assert.New(t)
	storage := NewFilesystemStorage("/nonexistent")

	data, err := storage.ReadPackages()
	assert.NotNil(err)
	assert.Equal("", data)
}

func TestReadPackages_RepoWithPackagefile(t *testing.T) {
	assert := assert.New(t)
	_ = os.MkdirAll("tmp/r1", os.ModePerm)
	defer func () {
		_ = os.RemoveAll("tmp")
	}()
	_ = ioutil.WriteFile("tmp/r1/packages.bss", []byte(`contents`), os.ModePerm)

	storage := NewFilesystemStorage("tmp/r1")

	data, err := storage.ReadPackages()
	assert.Nil(err)
	assert.Equal(`contents`, data)
}