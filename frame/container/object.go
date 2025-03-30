package container

// Object 对象
type Object interface {
	Name() string
}

// BaseObject 基础对象
type BaseObject struct {
	_name string `json:"name"`
}

// Name 获取名称
func (b *BaseObject) Name() string {
	return b._name
}

// SetName 设置名称
func (b *BaseObject) SetName(name string) {
	b._name = name
}
