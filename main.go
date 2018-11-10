package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/juju/mgosession"
	"github.com/urfave/negroni"
	mgo "gopkg.in/mgo.v2"

	"github.com/samayamnag/boilerplate/config"
	"github.com/samayamnag/boilerplate/app/http/controllers/v1"
	"github.com/samayamnag/boilerplate/app/http/middlewares"
	"github.com/samayamnag/boilerplate/app/repositories"
	"github.com/samayamnag/boilerplate/app/services"
)


func main() {
	config.Init()
	connStr := fmt.Sprintf("%s:%d", config.Get().MongoDBHost, config.Get().MongoDBPort)
	session, err := mgo.Dial(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	mPool := mgosession.NewPool(nil, session, config.Get().MongoDBConnPool)
	defer mPool.Close()

	r := mux.NewRouter()

	userRepo := repositories.NewMongoRepository(mPool, config.Get().MongoDBName)
	userService := services.NewUserService(userRepo)

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middlewares.Cors),
		negroni.NewLogger(),
	)

	// User handlers
	v1.MakeUserHandlers(r, *n, userService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(8080),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}


