package main

import (
	"net/http"
	"strings"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.FormValue("name"))
	w.Write([]byte("main" + name))
}

func test1Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test1"))
}

func test2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test2"))
}
