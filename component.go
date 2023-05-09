package tea

// ComponentType
type ComponentType string

const (
	CoreComponent      ComponentType = "core"
	UserComponent      ComponentType = "user"
	ThirdPartComponent ComponentType = "third-part"
)

// String
func (ct ComponentType) String() string {
	return string(ct)
}

// IComponent
type IComponent interface {
	Register(name string) error
	Type() ComponentType
}
