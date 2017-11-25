package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template // package template, Type Template

func init() {
	tpl = template.Must(template.ParseGlob("C:\\Users\\Dinesh Auti\\IdeaProjects\\go\\src\\github.com\\dineshauti\\webDevTraining\\004_parse-execute\\02_parseGlob\\template\\*"))
}

func main() {

	// Create a file
	nf, err := os.Create("C:\\Users\\Dinesh Auti\\IdeaProjects\\go\\src\\github.com\\dineshauti\\webDevTraining\\004_parse-execute\\02_parseGlob\\index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	// Copy the template to the created file
	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}


}
