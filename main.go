package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {

	Handlefunc()
}
func Handlefunc() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("Page/accueil.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	http.HandleFunc("/maintenance", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("Page/index.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("Page/nos_services.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	fs := http.FileServer(http.Dir("Static/"))
	http.Handle("/Static/", http.StripPrefix("/Static/", fs))
	fmt.Println("http://biggerthanme.fr:443")
	http.ListenAndServe("biggerthanme.fr:443", nil)
}
