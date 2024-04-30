package core

import (
	"sync"
)

/******************************************************************
 * 容器 Container 是最大粒度的抽象概念
 * 容器应该包括Name和New两个方法
 ******************************************************************/

var (
	// iocMap 容器映射
	iocMap map[string]IContainer = make(map[string]IContainer)

	// ioc instance 容器实例
	// 私有化，防止被外部修改
	iocInstance *ioc

	// once 保证只初始化一次
	once sync.Once
)

// IContainer 容器接口
type IContainer interface {
	// 名称
	Name() string
	// 新容器
	New() IContainer
}

// ioc 容器实现
type ioc struct {
	// 读写互斥锁
	lock sync.RWMutex
}

// IOC 获取容器实例
func IOC() *ioc {
	// 保证只初始化一次
	once.Do(func() {
		iocInstance = &ioc{}
	})
	return iocInstance
}

// Register 注册容器
// 容器不可重复注册，不可覆盖，不可修改
func (i *ioc) Register(ic IContainer) IError {
	// 锁
	i.lock.Lock()
	// 解锁
	defer i.lock.Unlock()
	// 判断容器是否已注册
	if _, ok := iocMap[ic.Name()]; !ok {
		// 注册容器
		iocMap[ic.Name()] = ic
		return nil
	}
	// 容器已注册
	return NameRegistered
}

// Get 获取容器
// 容器不存在返回错误
// 容器已注册返回新容器
func (i *ioc) Get(name string) (IContainer, IError) {
	// 锁
	i.lock.RLock()
	// 解锁
	defer i.lock.RUnlock()
	// 判断容器是否已注册
	if ic, ok := iocMap[name]; ok {
		// 基于容器创建新容器
		return ic.New(), nil
	}
	// 容器不存在
	return nil, NameNotRegistered
}
