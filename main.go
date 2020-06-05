package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"material_storage/models"
	"material_storage/repository"
	"net/http"
)

type application struct {
	servicePort         int
	materialsRepository repository.Material
	s                   *http.ServeMux
}

var (
	conf models.Config
	tpl  = template.Must(template.ParseFiles("templates/index.html"))
)

func init() {
	models.LoadConfig(&conf)
}

func main() {
	app := NewApplication(conf)
	app.initServer()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(`:%d`, app.servicePort), app.s))
}

func NewApplication(conf models.Config) *application {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.SQLDataBase.Server, "5433", conf.SQLDataBase.User, conf.SQLDataBase.Password, conf.SQLDataBase.Database)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &application{
		servicePort:         8000,
		materialsRepository: repository.NewMaterialRepository(db),
	}
}

func (app *application) initServer() {
	app.s = http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	app.s.Handle("/assets/", http.StripPrefix("/assets/", fs))
	app.s.HandleFunc("/materials", app.ListHandler)
	app.s.HandleFunc("/add", app.AddHandler)
}

func (app *application) ListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "private, max-age=0, no-cache")
	materials, err := app.materialsRepository.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = tpl.Execute(w, materials)
	if err != nil {
		log.Println(err)
	}
}
func (app *application) AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		description := r.FormValue("description")
		ref := r.FormValue("ref")
		app.materialsRepository.Add(title, description, ref)
		http.Redirect(w, r, "/materials", http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
