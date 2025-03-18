package frame

import "github.com/kearth/tea/frame/load"

type Config struct {
}

type Tea struct {
	steps  []load.Step
	config *Config
}

// AddHotWater 添加热水
func (t *Tea) AddHotWater() *Tea {
	return t
}

// BrewForAFewTime 煮一段时间
func (t *Tea) BrewForAFewTime() *Tea {
	return t
}

func (t *Tea) PourIntoCup() {

}

// GetSomeTea 获取一些茶
func GetSomeTea() *Tea {
	return &Tea{}
}

// HTTPServer 获取HTTPServer实例
// panic 获取失败时抛出异常
// func HTTPServer() *httpServer {
// 	var hs IContainer
// 	var err error
// 	// 获取HTTPServer实例
// 	if hs, err = IOC().Get(new(httpServer).Name()); err != nil {
// 		panic(err)
// 	}
// 	return hs.(*httpServer)
// }
