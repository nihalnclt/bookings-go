package handlers

import (
	"log"
	"net/http"

	"github.com/nihalnclt/bookings-go/internal/config"
	"github.com/nihalnclt/bookings-go/internal/driver"
	"github.com/nihalnclt/bookings-go/internal/forms"
	"github.com/nihalnclt/bookings-go/internal/models"
	"github.com/nihalnclt/bookings-go/internal/render"
	"github.com/nihalnclt/bookings-go/internal/repository"
	"github.com/nihalnclt/bookings-go/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ShowLogin(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "login.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository) PostShowLogin(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	// email := r.Form.Get("email")
	// password := r.Form.Get("password")

	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.IsEmail("email")

	if !form.Valid() {
		render.Template(w, r, "login.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
