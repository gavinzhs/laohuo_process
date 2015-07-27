package main
import (
    "net/http")

func main(){
    print("start")

    http.HandleFunc("/", mainHandler)
    http.HandleFunc("/test1", test1Handler)
    http.HandleFunc("/test2", test2Handler)

    http.ListenAndServe(":3000", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("main"))
}

func test1Handler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("test1"))
}

func test2Handler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("test2"))
}