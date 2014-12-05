package mocks

import "github.com/stretchr/testify/mock"

type Driver struct {
	mock.Mock
}

func (m *Driver) AddPackage(packageName string) error {
	ret := m.Called(packageName)

	r0 := ret.Error(0)

	return r0
}
