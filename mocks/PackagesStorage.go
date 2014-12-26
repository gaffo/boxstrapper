package mocks

import "github.com/gaffo/boxstrapper"
import "github.com/stretchr/testify/mock"

type PackagesStorage struct {
	mock.Mock
}

func (m *PackagesStorage) ReadPackages() ([]*boxstrapper.Package, error) {
	ret := m.Called()

	r0 := ret.Get(0).([]*boxstrapper.Package)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *PackagesStorage) WritePackages(_a0 []*boxstrapper.Package, _a1 string) error {
	ret := m.Called(_a0, _a1)

	r0 := ret.Error(0)

	return r0
}
