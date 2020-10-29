package restapis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/hooneun/golang-web-tutorial/app/models"
	"github.com/hooneun/golang-web-tutorial/app/models/dblayer"
)

// HandlerInterface !
type HandlerInterface interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

// Handler !
type Handler struct {
	db dblayer.DBLayer
}

// NewHandler create!
func NewHandler() (*Handler, error) {
	db, err := models.NewORM()
	if err != nil {
		return nil, err
	}

	return &Handler{
		db: db,
	}, nil
}

// GetUserByID Handler
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if h.db == nil {
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, err := h.db.GetUserByID(id)

	if err != nil {

	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

// CreateUser Handler
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if h.db == nil {
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

	}

	user, err = h.db.CreateUser(user)

	if err != nil {
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&user)
	// return user, db.Create(&user).Error
}
