package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/fuadsuleyman/netproject/internal/handlers"
	"github.com/fuadsuleyman/netproject/pkg/logging"
)

const (
	usersURL = "/users"
	userURL = "/users/:uuid"
)

// handler is local with lower case
type handler struct {
	logger logging.Logger
} 

// return interface
func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{logger: logger}
}

func (h *handler) Register(router *httprouter.Router){
	router.HandlerFunc(http.MethodGet, usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

// get all users
func (h *handler) GetList(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)
	w.Write([]byte("This is list of users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(200)
	w.Write([]byte("This is get user by uuid"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(201)
	w.Write([]byte("This is create user"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(200)
	w.Write([]byte("This is update user"))
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(200)
	w.Write([]byte("This is partially update user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(200)
	w.Write([]byte("This is delete user"))
}