package boxstrapper_test

import (
	"testing"
	. "github.com/gaffo/boxstrapper"
  	"github.com/stretchr/testify/assert"
  	"github.com/stretchr/testify/mock"
)

type MockDriver struct {
	mock.Mock
}

func (m *MockDriver) AddPackage(packageName string) error {
	ret := m.Called(packageName)
	r0 := ret.Error(0)
	return r0
}

func TestApCallsToDriver(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1"}
	driver := new(MockDriver)
	driver.On("AddPackage", "package1").Return(nil).Once()
	
	err := Ap(driver, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
}

func TestApCallsToDriver_MultiplePackages(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1", "package2"}
	driver := new(MockDriver)
	driver.On("AddPackage", "package1").Return(nil).Once()
	driver.On("AddPackage", "package2").Return(nil).Once()
	
	err := Ap(driver, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
}