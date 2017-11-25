package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func main() {


	name := os.Args[1]

	tpl := fmt.Sprintf(`<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>Hello World!</title>
				</head>
				<body>
					<h1>` + name + `</h1>
				</body>
			</html>`)

	nf, err := os.Create("C:\\Users\\Dinesh Auti\\Desktop\\index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
