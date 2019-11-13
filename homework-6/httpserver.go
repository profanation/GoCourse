package main

import (
	"fmt"
    "io"
    "net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	//reqParam:= req.RequestURI
	reqParam:=req.URL.Query().Get("name")
	println(reqParam)
    io.WriteString(res, reqParam)

}


func main() {
    fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

    http.ListenAndServe(":8080", nil)
}