package main

import (
	logger "log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"dbrepo/userrepo"
	handlerlib "delivery/restapplication/packages/httphandlers"
	"delivery/restapplication/usercrudhandler"
)

func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	   https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {

	dbrepo := userrepo.NewUserInMemRepository()
	usersvc := userrepo.NewService(dbrepo)

	hndlr := usercrudhandler.NewUserCrudHandler(usersvc)

	pingHandler := &handlerlib.PingHandler{}
	logger.Println("---Setup done---")
	logger.Println("Starting service for you")
	h := mux.NewRouter()
	h.Handle("/ping/", pingHandler)
	h.Handle("/user/{id}", hndlr)
	h.Handle("/user/", hndlr)
	logger.Println("Resource Setup Done.")
	logger.Fatal(http.ListenAndServe(":8080", h))
}
