package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	tpl, err := template.ParseFiles("C:\\Users\\Dinesh Auti\\IdeaProjects\\go\\src\\github.com\\dineshauti\\webDevTraining\\004_parse-execute\\01_stdout\\tpl.gohtml")
	if err != nil {
		log.Fatalln("Can't parse HTML file", err)
	}

	nf, err := os.Create("C:\\Users\\Dinesh Auti\\IdeaProjects\\go\\src\\github.com\\dineshauti\\webDevTraining\\004_parse-execute\\01_stdout\\index.html")
	if err != nil {
		log.Fatalln("Error while creating the file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln("Can't execute the template", err)
	}
}
