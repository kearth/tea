/**
* Tea Framework
* Version 1.0.1
*
* Copyright 2018 - 2023, Kearth
* Golang 1.21
 */
package tea

import "github.com/kearth/tea/frame"

// 喝
func Drink(t *frame.Tea) {
	t.PourIntoCup()
}

// 一杯茶
func ACupOfTea() *frame.Tea {
	return frame.GetSomeTea().AddHotWater().BrewForAFewTime()
}
