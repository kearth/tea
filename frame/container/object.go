package container

type Object interface {
	Name() string
}

type BaseObject struct {
	_name string
}

func (b *BaseObject) Name() string {
	return b._name
}

func (b *BaseObject) SetName(name string) {
	b._name = name
}
