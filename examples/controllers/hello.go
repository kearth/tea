package controllers

import (
	"io"
	"net/http"
)

func Hello(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "<html>Hello World!</html>\n")
}
