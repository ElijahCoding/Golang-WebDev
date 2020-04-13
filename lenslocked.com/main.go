package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func handleFunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>super welcome</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "Mail me <a href=\"mailto:support@google.com\">support</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404 page</h1>")
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main()  {
	router := httprouter.New()
	router.GET("/hello/:name", Hello)
	//http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", router)
}