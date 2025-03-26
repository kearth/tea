package utils

import "github.com/gogf/gf/v2/os/gfile"

// Path 路径拼接
func Path(f ...string) string {
	return gfile.Join(f...)
}
