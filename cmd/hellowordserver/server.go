package main

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Fprintf(writer,"hello world %s",request.FormValue("name"))
		fmt.Fprintln(writer,"<h1>hello world</h1>")
	})
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic(err)
	}
}
