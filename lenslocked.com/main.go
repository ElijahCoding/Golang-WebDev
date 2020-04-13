package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
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

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>NOT FOUND</h1>")
}

func main()  {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", handleFunc)
	r.HandleFunc("/faq", faq)
	//http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", r)
}