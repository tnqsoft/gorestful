package app

import (
	"fmt"
	"log"
	"net/http"

	"../config"
	"../handler"
	"../models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
	Config *config.Config
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {

	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	// defer db.Close()
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)

	a.Config = config
	a.DB = models.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/contacts", a.GetAllContacts)
	a.Post("/contacts", a.CreateContact)
	a.Get("/contacts/{id}", a.GetContact)
	a.Put("/contacts/{id}", a.UpdateContact)
	a.Delete("/contacts/{id}", a.DeleteContact)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Contact Data
func (a *App) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllContacts(a.DB, w, r)
}

func (a *App) CreateContact(w http.ResponseWriter, r *http.Request) {
	handler.CreateContact(a.DB, w, r)
}

func (a *App) GetContact(w http.ResponseWriter, r *http.Request) {
	handler.GetContact(a.DB, w, r)
}

func (a *App) UpdateContact(w http.ResponseWriter, r *http.Request) {
	handler.UpdateContact(a.DB, w, r)
}

func (a *App) DeleteContact(w http.ResponseWriter, r *http.Request) {
	handler.DeleteContact(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run() {
	log.Fatal(http.ListenAndServe(a.Config.APP.Hostname, a.Router))
}
