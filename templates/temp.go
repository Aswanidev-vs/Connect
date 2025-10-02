package templates

import (
	"html/template"
	"log"
)

var Templates = map[string]*template.Template{}

func LoadTemplates() {
	var err error
	Templates["signup"], err = template.ParseFiles("templates/signup.html")
	if err != nil {
		log.Fatal("Error parsing signup.html:", err)
	}

	Templates["login"], err = template.ParseFiles("templates/login.html")
	if err != nil {
		log.Fatal("Error parsing login.html:", err)
	}

	Templates["index"], err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Error parsing index.html:", err)
	}
}
