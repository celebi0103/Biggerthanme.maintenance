package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/equipes", equipesHandler)
	http.HandleFunc("/blog", blogHandler)
	http.HandleFunc("/articles", articlesHandler)
	http.HandleFunc("/articlesB", articlesBHandler)
	http.HandleFunc("/articlesBasket", articlesBasketHandler)
	http.HandleFunc("/articlesHandi", articlesHandiHandler)
	http.HandleFunc("/services", servicesHandler)
	http.HandleFunc("/articlesdanse", articlesDanseHandler)
	http.HandleFunc("/articlesfede", articlesFedeHandler)

	fs := http.FileServer(http.Dir("Static/"))
	http.Handle("/Static/", http.StripPrefix("/Static/", fs))

	fmt.Println("Listening on port 80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/index.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.MethodPost {
		template.Execute(w, nil)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/contact.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.MethodPost {
		template.Execute(w, nil)
		return
	}
}

func equipesHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/equipes.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.MethodPost {
		template.Execute(w, nil)
		return
	}
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/blog.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.MethodPost {
		template.Execute(w, nil)
		return
	}
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/articles.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.MethodPost {
		template.Execute(w, nil)
		return
	}
}

func articlesBHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/articleBoxe.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.MethodPost {
		template.Execute(w, nil)
		return
	}
}

func articlesBasketHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/articlebasket.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.MethodPost {
		template.Execute(w, nil)
		return
	}
}

func articlesHandiHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/articlehandi.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.MethodPost {
		template.Execute(w, nil)
		return
	}
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/nos_services.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.MethodPost {
		template.Execute(w, nil)
		return
	}
}

func articlesDanseHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/articleDanse.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.Method {
			template.Execute(w, nil)
		return
	}
}

func articlesFedeHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./Page/articlefederation.html", "templates/header.html", "templates/footer.html"))
	if r.Method != http.Method {
			template.Execute(w, nil)
		return
	}
}
http.ListenAndServe(":80", http.HandlerFunc(Handlefunc))
