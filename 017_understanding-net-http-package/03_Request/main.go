package main

import (
	"net/http"
	"log"
	"html/template"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("C:\\Users\\Dinesh Auti\\IdeaProjects\\go\\src\\github.com\\dineshauti\\webDevTraining\\017_understanding-net-http-package\\03_Request\\index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
