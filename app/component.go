package app

type ComponentType string

const (
	CoreComponent      ComponentType = "core"
	UserComponent      ComponentType = "user"
	ThirdPartComponent ComponentType = "third-part"
)

func (ct ComponentType) String() string {
	return string(ct)
}

type IComponent interface {
	Register(name string) error
	Type() ComponentType
}
