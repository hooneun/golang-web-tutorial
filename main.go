package main

import (
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/hooneun/golang-web-tutorial/app/models"
	"github.com/hooneun/golang-web-tutorial/app/restapis"
)

var db, _ = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:4444)/books?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

// Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging !
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Info(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

// Chain middlewares
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

func main() {
	r := mux.NewRouter()
	h, err := restapis.NewHandler()

	if err != nil {
		log.Error(err)
	}

	db.Debug().Migrator().DropTable(&models.User{}, &models.Todo{})
	db.Debug().AutoMigrate(&models.User{}, &models.Todo{})

	r.HandleFunc("/users", Chain(h.CreateUser, Logging())).Methods("POST")
	r.HandleFunc("/users/{id}", Chain(h.GetUserByID, Logging())).Methods("GET")

	// user := User{
	// 	Name:     "johndoe",
	// 	Password: "secret",
	// 	Email:    "test@test.com",
	// }

	// db.Create(&user)
	// db.Delete(&user)

	// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	title := vars["title"]
	// 	page := vars["page"]
	// 	log.Info("/books/" + title + "/page/" + page)

	// 	fmt.Fprintf(w, "book: %s on page %s\n", title, page)
	// })

	// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	// bookrouter := r.PathPrefix("/books").Subrouter()
	// bookrouter.HandleFunc("/", AllBooks)
	// bookrouter.HandleFunc("/{title}", GetBook)

	http.ListenAndServe(":8080", r)
}

// log.Warn("TodoItem not found in database")
// 	log.WithFields(log.Fields{"Id": id}).Info("Deleting TodoItem")
