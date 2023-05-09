package tea

import (
	"errors"
	"sync"
)

// IContainer
type IContainer interface {
	Name() string
	New() IContainer
}

// ioc
type ioc struct{}

var (
	// iocMap
	iocMap map[string]IContainer = make(map[string]IContainer)

	// ioc instance
	iocInstance *ioc

	// once
	once sync.Once
)

// IOC
func IOC() *ioc {
	once.Do(func() {
		iocInstance = &ioc{}
	})
	return iocInstance
}

// Register
func (i *ioc) Register(ic IContainer) error {
	if _, ok := iocMap[ic.Name()]; !ok {
		iocMap[ic.Name()] = ic
		return nil
	}
	return errors.New("the name has registered")
}

// Get
func (i *ioc) Get(name string) (IContainer, error) {
	if ic, ok := iocMap[name]; ok {
		return ic.New(), nil
	}
	return nil, errors.New("the name not regisered")
}
