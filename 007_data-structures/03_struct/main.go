package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

type sages struct {
	sage
	Country string
}

func init() {
	tpl = template.Must(template.ParseGlob("C:\\Users\\Dinesh Auti\\IdeaProjects\\go\\src\\github.com\\dineshauti\\webDevTraining\\007_data-structures\\03_struct\\tpl.gohtml"))
}

func main() {

	buddha := sages{
		sage{
			Name:"Buddha",
			Motto:"The belief of no beliefs",
		},
		"India",
	}

	mlk := sages{
		sage{
			Name:"Martin Luther King",
			Motto:"Darkness cannot drive out darkness; only light can do that. Hate cannot drive out hate; only love can do that.",
		},
		"USA",
	}

	x := []sages{buddha, mlk}


	err := tpl.Execute(os.Stdout, x)
	if err != nil {
		log.Fatalln(err)
	}
}
