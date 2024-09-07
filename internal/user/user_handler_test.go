package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	router := httprouter.New()
	h := NewHandler()
	h.Register(router)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "this is list of users", rr.Body.String())
}

func TestGetUserByUUID(t *testing.T) {
	router := httprouter.New()
	h := NewHandler()
	h.Register(router)

	req, err := http.NewRequest("GET", "/user/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "this is user by id", rr.Body.String())
}

func TestCreateUser(t *testing.T) {
	router := httprouter.New()
	h := NewHandler()
	h.Register(router)

	req, err := http.NewRequest("POST", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "this is create user", rr.Body.String())
}

func TestUpdateUser(t *testing.T) {
	router := httprouter.New()
	h := NewHandler()
	h.Register(router)

	req, err := http.NewRequest("PUT", "/user/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestPartiallyUpdateUser(t *testing.T) {
	router := httprouter.New()
	h := NewHandler()
	h.Register(router)

	req, err := http.NewRequest("PATCH", "/user/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteUser(t *testing.T) {
	router := httprouter.New()
	h := NewHandler()
	h.Register(router)

	req, err := http.NewRequest("DELETE", "/user/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
