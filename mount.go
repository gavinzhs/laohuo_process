package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"net/http"
)

func mount(war string) {

	m := martini.Classic()
	m.Handlers(martini.Recovery())
	m.Use(gzip.All())
	m.Use(martini.Static(war, martini.StaticOptions{SkipLogging: true}))
	//    m.Use(render.Renderer(render.Options{
	//        Extensions: []string{".html", ".shtml"},
	//    }))
	//	m.Use(render.Renderer())
	//    m.Use(midTextDefault)

	//map web
	m.Use(func(w http.ResponseWriter, c martini.Context) {
		web := &Web{w: w}
		c.Map(web)
	})

	m.Group("/test", func(api martini.Router) {
		api.Get("", mainHandler)
		api.Get("/1", test1Handler)
		api.Get("/2", test2Handler)
	})

	http.Handle("/", m)
}
