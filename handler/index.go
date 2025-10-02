package handler

import (
	"net/http"

	tpl "github.com/Aswanidev-vs/Connect/templates"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.Templates["index"].Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
