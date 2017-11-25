package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("C:\\Users\\Dinesh Auti\\IdeaProjects\\go\\src\\github.com\\dineshauti\\webDevTraining\\007_data-structures\\02_map\\template\\*"))
}

func main() {

	colors := map[string]string{
		"White":"#FF",
		"Black":"#00",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", colors)
	if err != nil {
		log.Fatalln(err)
	}

}
