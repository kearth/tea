package bootstrap

import (
	"fmt"

	"github.com/kearth/tea"
)

func Bootstrap(ctx tea.Context) tea.IError {
	fmt.Println("hello world")
	return nil
}
