package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"net/http"
)

/*
 web utilities
*/
type Web struct {
	w http.ResponseWriter
}

func (p *Web) Json(code int, data interface{}) (int, string) {
	b, err := json.Marshal(data)
	chk(err)
	p.w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return code, string(b)
}

func (p *Web) Code(code int) (int, string) {
	return code, http.StatusText(code)
}

func dup(err error) bool {
	return mgo.IsDup(err)
}

func notFound(err error) bool {
	return err == mgo.ErrNotFound
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
