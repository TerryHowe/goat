package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

/* mazeHandler(res http.ResponseWriter, req *http.Request) string*/
func mazeHandler(req *http.Request, r render.Render) {
	data := struct {
		Left string
		Right string
		Backward string
		Forward string
		DirectionLong string
		X int
		Y int
		MazeView string
	}{
		Left: "left",
		Right: "right",
		Backward: "backward",
		Forward: "forward",
		DirectionLong: "North",
		X: 0,
		Y: 1,
		MazeView: "BLAH",
	}
	r.HTML(200, "room", data)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", mazeHandler)
	m.Run()
}