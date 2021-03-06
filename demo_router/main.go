package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an eemail to <a href=\"mailto:support@gmail.com\">support@gmail.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for : </h1><p>(Please email us if you keep being sent to an invalid page.)</p>")
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/hello/:name", Hello)

	http.ListenAndServe("localhost:3000", router)
}
