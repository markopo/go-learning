package handlers

import (
	"github.com/markopo/go-learning/pkg/config"
	"github.com/markopo/go-learning/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home Page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About Page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

