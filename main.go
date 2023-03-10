package main

import (
	"fmt"
	"html/template"
	"net/http"

	"example.com/hello/models"
	"example.com/hello/utils"
)

var posts map[string]*models.Post

func indexHendler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Println(posts)
	t.ExecuteTemplate(w, "index", posts)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/writer.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "writer", nil)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/writer.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	id := r.FormValue("id")
	post, found := posts[id]
	if !found {
		http.NotFound(w, r)
	}

	t.ExecuteTemplate(w, "writer", post)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else {
		id = utils.GenereteId()
		post := models.NewPost(id, title, content)
		posts[post.Id] = post

	}

	http.Redirect(w, r, "/", 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
	}

	delete(posts, id)

	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("Listening on port :3000")

	posts = make(map[string]*models.Post, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHendler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/SavePost", savePostHandler)

	http.ListenAndServe(":3000", nil)
}

// 26.38
