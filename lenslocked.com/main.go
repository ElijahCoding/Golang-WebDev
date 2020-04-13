package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>super welcome</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "Mail me <a href=\"mailto:support@google.com\">support</a>")
	}
}

func main()  {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)
}