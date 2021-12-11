package handlers

import (
	"github.com/markopo/go-learning/pkg/render"
	"net/http"
)

// Home Page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About Page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}


