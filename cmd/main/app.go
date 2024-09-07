package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"rest-api-practice/internal/user"
	"rest-api-practice/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create logger")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()

	logger.Info("start application")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening on port 0,0,0,0:8080")
	logger.Fatal(server.Serve(listener))
}
