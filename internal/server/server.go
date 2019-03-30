// Package server defines rest api endpoints for pinger
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/saromanov/gleek/config"
	"github.com/saromanov/gleek/internal/storage"
	pb "github.com/saromanov/gleek/proto"
	log "go.uber.org/zap"
)

var tokenAuth *jwtauth.JWTAuth

type server struct {
	router  *chi.Mux
	address string
	st      *storage.Storage
}

// createAccount makes a new task
func (s *server) createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	account := &pb.Task{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *server) makeHandlers() {
	s.router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
	})

	s.router.Group(func(r chi.Router) {
		r.Post("/v1/tasks", s.createTask)
	})
}

func (s *server) startServer() {
	fmt.Printf("server is started at %s", s.address)
	srv := &http.Server{
		Addr:         s.address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	fmt.Printf("shutting down server")
}

// New makes http endpoints and handler
func New(st *storage.Storage, c *config.Config) {
	tokenAuth = jwtauth.New("HS256", []byte("testtoken"), nil)
	r := chi.NewRouter()
	s := &server{
		st:      st,
		router:  r,
		address: c.Address,
	}
	s.makeHandlers()
	s.startServer()
}
