package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lenslocked/views"
)

var homeView *views.View
var contactView *views.View

var homeTemplate *template.Template
var contactTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "This is the FAQ page ")
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "My own 404 page")
}

func main() {

	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	var h http.Handler = http.HandlerFunc(notfound)
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)
}
