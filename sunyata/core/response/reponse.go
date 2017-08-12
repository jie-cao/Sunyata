package response

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/flosch/pongo2"
)

const ContentTypeHeader = "Content-Type"

func View(w http.ResponseWriter, viewName string, data interface{}) {
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	// Compile the template first (i. e. creating the AST)
	tpl, err := pongo2.FromFile("views/main/test.html")
	if err != nil {
		panic(err)
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	out, err := tpl.Execute(pongo2.Context{"name": "florian"})
	if err != nil {
		panic(err)

	}
	fmt.Fprintf(w, out)
	/*
		s1, error := template.ParseFiles("views/main/test.tmpl", "views/shared/header.tmpl")
		s1.ExecuteTemplate(w, "header", nil)
		s1.ExecuteTemplate(w, "content", data)
		if error == nil {
			s1.Execute(w, data) //执行模板的merger操作
		}
	*/
}

func ViewWithLayout(w http.ResponseWriter, LayoutName string, viewName string, data interface{}) {
	s1, error := template.ParseFiles("views/main/test.tmpl", "views/shared/header.tmpl")
	s1.ExecuteTemplate(w, "header", nil)
	s1.ExecuteTemplate(w, "content", data)
	if error == nil {
		s1.Execute(w, data) //执行模板的merger操作
	}
}

func Json(w http.ResponseWriter, data interface{}) {
	jsonResponse, ok := json.Marshal(data)
	if ok == nil {
		w.Header().Set(ContentTypeHeader, "application/json")
		fmt.Fprintf(w, string(jsonResponse))
	}
}
