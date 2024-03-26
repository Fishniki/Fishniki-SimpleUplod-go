package controller

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func ImageUplod(w http.ResponseWriter, r *http.Request) {

	fp := filepath.Join("view", "index.html")
	tmpl, err := template.ParseFiles(fp)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	

}
