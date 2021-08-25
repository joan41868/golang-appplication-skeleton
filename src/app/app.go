package app

import (
	"log"
	"net/http"
	"sync"

	"github.com/joan41868/go-app-skeleton/src/repository"
	"github.com/joan41868/go-app-skeleton/src/server"
	"github.com/joho/godotenv"
)

type Application struct {
	repository *repository.Repository
	server     *server.Server

	servicesWg *sync.WaitGroup
}

func NewApp() *Application {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	app := &Application{
		server:     server.NewServer(),
		repository: nil,
		servicesWg: new(sync.WaitGroup),
	}

	return app
}

// Setup should perform additional app configuration/loading/etc. Any cross-package logic can be included here as well
func (app *Application) Setup() {
	// example:
	app.server.AddRoute("/hello", "GET", func(wr http.ResponseWriter, r *http.Request) {
		wr.Write([]byte("Hello world!"))
	})
}

func (app Application) Start() {
	app.servicesWg.Add(1)
	go app.server.Start()

	app.servicesWg.Wait()
}
