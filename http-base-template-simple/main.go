package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"go-template-example/http-base-template/application/server"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	templatePattern := filepath.Join(wd, "/application/web/templates/*")
	baseTemplatePath := filepath.Join(wd, "/application/web/templates/base.gohtml")
	templates := server.NewTemplates(templatePattern, baseTemplatePath)
	for key, val := range templates {
		if key == "sign-in" {
			fmt.Printf("Template key %s, with name %s:\n", key, val.Name())
			fmt.Println(val.DefinedTemplates())

			for _, tmpl := range val.Templates() {
				fmt.Printf("Rendering Template %s\n", tmpl.Name())
				val.ExecuteTemplate(os.Stdout, tmpl.Name(), "")
			}
		}

	}

	templateRenderer := server.NewTemplateRenderer(templates, "base")

	webHandler := server.NewWebHandler(templateRenderer)

	httpStaticAssetsDir := http.Dir(fmt.Sprintf("%s/application/web/static/", wd))
	staticRoute := "/static/"
	staticAssetHandler := http.StripPrefix(
		staticRoute,
		http.FileServer(httpStaticAssetsDir),
	)

	router := chi.NewRouter()

	// Suggested basic middleware stack from chi's docs
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Handle(staticRoute+"*", staticAssetHandler)
	router.Get("/login", webHandler.GetLogin)
	router.Get("/register", webHandler.GetRegister)

	host := "localhost"
	port := "8080"

	srv := &http.Server{
		Handler:      router,
		Addr:         host + ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("running http server on port %s...\n", port)
	log.Fatal(srv.ListenAndServe())
}
