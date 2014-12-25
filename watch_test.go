package boxstrapper_test

import (
	. "github.com/gaffo/boxstrapper"
	"github.com/gaffo/boxstrapper/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WatchNoArgs_DoesNothing(t *testing.T) {
	assert := assert.New(t)

	driver := new(mocks.Driver)
	storage := new(mocks.Storage)

	err := Watch(driver, storage, []string{})
	assert.Nil(err)
}
