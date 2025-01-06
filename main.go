package main

/*
	Main file with routes and nothing else.
*/

import (
	"log"
	"net/http"

	"github.com/dixxe/personal-website/service/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	log.Println(`
 ____  _ ___  ____  _ _____      ____  ____  _____ _____
/  _ \/ \\  \//\  \///  __/     /  __\/  _ \/  __//  __/
| | \|| | \  /  \  / |  \ _____ |  \/|| / \|| |  _|  \  
| |_/|| | /  \  /  \ |  /_\____\|  __/| |-||| |_//|  /_ 
\____/\_//__/\\/__/\\\____\     \_/   \_/ \|\____\\____\
	   	`)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)

	r.Get("/", controllers.GetIndexHandler)

	r.Get("/blog", controllers.GetShowBlog)
	r.Post("/post", controllers.PostCreatePost)
	r.Post("/post/delete", controllers.PostDeletePost)

	r.Post("/admin/login", controllers.PostAdminLogin)
	r.Get("/admin", controllers.GetAdminLogin)

	fs := http.FileServer(http.Dir("resources/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	log.Println("Website started")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln(err)
	}
}
