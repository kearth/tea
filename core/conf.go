package core

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
)

type ParserFunc func(b []byte, obj any) error

type Conf interface {
	Exists(string) bool
	Parse(string, any) error
	RegisterParserFunc(string, ParserFunc) error
}

const (
	FileTOML = ".toml"
	FileJSON = ".json"
)

var (
	TOMLParserFunc ParserFunc = toml.Unmarshal
	JSOMParserFunc ParserFunc = json.Unmarshal

	ErrNoExists = errors.New("conf no exists")
	ErrNoParser = errors.New("no parser found")
	ErrFileExt  = errors.New("not file extension")

	ParserFuncs = map[string]ParserFunc{
		FileTOML: TOMLParserFunc,
		FileJSON: JSOMParserFunc,
	}
)

// Exists
func Exists(fname string) bool {
	_, err := os.Stat(fname)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// Parse
func Parse(fname string, confObj any) error {
	var fn ParserFunc
	var ok bool
	if !Exists(fname) {
		return ErrNoExists
	}
	ext := path.Ext(fname)
	if fn, ok = ParserFuncs[ext]; !ok {
		return ErrNoParser
	}
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		return err
	}
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	b := make([]byte, fi.Size())
	_, err = f.Read(b)
	if err != nil {
		return err
	}
	return fn(b, confObj)
}

// RegisterParserFunc
func RegisterParserFunc(fileExt string, fn ParserFunc) error {
	if strings.HasPrefix(fileExt, ".") {
		ParserFuncs[fileExt] = fn
		return nil
	}
	return ErrFileExt
}
