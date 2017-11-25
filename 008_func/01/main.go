package main

import (
	"text/template"
	"os"
	"log"
	"strings"
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

var fm = template.FuncMap{
	"uc":strings.ToUpper,
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("C:\\Users\\Dinesh Auti\\IdeaProjects\\go\\src\\github.com\\dineshauti\\webDevTraining\\008_func\\01\\tpl.gohtml"))
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
