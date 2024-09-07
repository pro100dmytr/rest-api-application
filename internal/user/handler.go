package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-practice/internal/handlers"
	"rest-api-practice/pkg/logging"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/user/:uuid"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("from GetList")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is list of users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("from GetUserByUUID")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is user by id"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("from CreateUser")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("this is create user"))

}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("from UpdateUser")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("this is update user"))

}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("from PartiallyUpdateUser")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("this is partially update user"))

}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("from DeleteUser")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("this is delete user"))

}
