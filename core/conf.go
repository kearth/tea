package core

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	TOMLParserFunc ParserFunc = toml.Unmarshal // 解析toml文件
	JSOMParserFunc ParserFunc = json.Unmarshal // 解析json文件

	// 解析器
	ParserFuncs = map[string]ParserFunc{
		FileTOML: TOMLParserFunc,
		FileJSON: JSOMParserFunc,
	}
)

const (
	FileTOML = ".toml" // toml
	FileJSON = ".json" // json
)

// 配置文件解析器
type ParserFunc func(b []byte, obj any) error

// Conf 配置接口
type Conf interface {
	Exists(string) bool
	Parse(string, any) error
	RegisterParserFunc(string, ParserFunc) IError
}

// Exists
func Exists(fname string) bool {
	_, err := os.Stat(fname)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// Parse
func Parse(fname string, confObj any) IError {
	var fn ParserFunc
	var ok bool
	if !Exists(fname) {
		return ConfNotExists
	}
	ext := path.Ext(fname)
	if fn, ok = ParserFuncs[ext]; !ok {
		return NoParserFound
	}
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		return FrameworkCoreError.Wrap(err)
	}
	fi, err := f.Stat()
	if err != nil {
		return FrameworkCoreError.Wrap(err)
	}
	b := make([]byte, fi.Size())
	_, err = f.Read(b)
	if err != nil {
		return FrameworkCoreError.Wrap(err)
	}
	err = fn(b, confObj)
	if err != nil {
		return FrameworkCoreError.Wrap(err)
	}
	return nil
}

// RegisterParserFunc 注册解析器
func RegisterParserFunc(fileExt string, fn ParserFunc) IError {
	if strings.HasPrefix(fileExt, ".") {
		ParserFuncs[fileExt] = fn
		return nil
	}
	return NoParserFound
}
