package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var db, _ = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:4444)/books?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

// Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// User struct
type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"index:idx_email,unique"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Todos    []Todo
}

// Todo struct
type Todo struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID uint   `json:"user_id" gorm:"index:idx_user"`
	Desc   string `json:"desc" gorm:"index:idx_desc"`
}

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

// CreateUser !
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return
	}

	password, _ := HashPassword(user.Password)
	user.Password = password

	db.Create(&user)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

// GetUserByID !
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var user User
	db.Find(&user, id)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

// Chain middlewares
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

// HashPassword !
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash !
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	r := mux.NewRouter()

	db.Debug().Migrator().DropTable(&User{}, Todo{})
	db.Debug().AutoMigrate(&User{}, &Todo{})

	r.HandleFunc("/users", Chain(CreateUser, Logging())).Methods("POST")
	r.HandleFunc("/users/{id}", Chain(GetUserByID, Logging())).Methods("GET")

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
