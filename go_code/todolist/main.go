package main

import (
	"go_code/todolist/model"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/list/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("view/list.html"))
		tpl.ExecuteTemplate(w, "list.html", model.GetModel())
	})

	http.HandleFunc("/add/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.PostFormValue("name")
			model.AddModel(name)
			http.Redirect(w, r, "/list/", http.StatusSeeOther)
		}
		tpl := template.Must(template.ParseFiles("view/add.html"))
		tpl.ExecuteTemplate(w, "add.html", nil)
	})

	http.ListenAndServe(":8080", nil)
}
