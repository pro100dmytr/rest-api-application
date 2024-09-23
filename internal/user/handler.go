package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-practice/internal/apperror"
	"rest-api-practice/internal/handlers"
	"rest-api-practice/pkg/logging"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/user/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	h.logger.Info("from GetList")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is list of users"))
	return nil
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	h.logger.Info("from GetUserByUUID")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is user by id"))
	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	h.logger.Info("from CreateUser")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("this is create user"))
	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	h.logger.Info("from UpdateUser")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("this is update user"))
	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	h.logger.Info("from PartiallyUpdateUser")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("this is partially update user"))
	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	h.logger.Info("from DeleteUser")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("this is delete user"))
	return nil
}
