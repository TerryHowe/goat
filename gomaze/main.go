package main

import (
	"github.com/go-martini/martini"
	"net/http"
)


func mazeHandler(res http.ResponseWriter, req *http.Request) string {
	return "<html><title>Hello world!</title><h2>Hello world!</h1></html>"
}

func main() {
	m := martini.Classic()
	m.Get("/", mazeHandler)
	m.Run()
}