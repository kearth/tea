package main

import (
	"example/local/app/load"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea"
)

/*******************************
 * 框架入口
 *******************************/
func main() {
	tea.Drink(kctx.New(), load.LoadAll)
}
