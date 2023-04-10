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
		template := template.Must(template.ParseFiles("./Page/index.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./Page/contact.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})
	http.HandleFunc("/equipes", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./Page/equipes.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})
	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./Page/blog.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./Page/articles.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	http.HandleFunc("/articlesB", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./Page/articleBoxe.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	http.HandleFunc("/articlesBasket", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./Page/articlebasket.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	http.HandleFunc("/articlesHandi", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./Page/articlehandi.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./Page/nos_services.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	http.HandleFunc("/articlesdanse", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./Page/articleDanse.html", "templates/header.html", "templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, "")
			return
		}
	})

	fs := http.FileServer(http.Dir("Static/"))
	http.Handle("/Static/", http.StripPrefix("/Static/", fs))
	fmt.Println("http://biggerthanme.fr:8080")
	http.ListenAndServe(":8080", nil)
}
