// REFERENCE : https://pythonprogramming.net/go/templating-webapp-go-language-programming-tutorial/
// MORE ON : https://golang.org/pkg/html/template/

package main

import (
	"fmt"
	"net/http"
	"html/template"        // Go's templating package
)

type htmlpage struct {   // we can only pass one object to our template, so we're going to use a struct to pass multiple values.
    Title string
    News string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func htmlpageHandler(w http.ResponseWriter, r *http.Request) {
	p := htmlpage{Title: "Amazing News Aggregator", News: "some news"}   // page data
    t, _ := template.ParseFiles("template.html")                         // template files
    t.Execute(w, p)                                                      // execute template with page variables
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", htmlpageHandler)
	http.ListenAndServe(":8000", nil) 
}