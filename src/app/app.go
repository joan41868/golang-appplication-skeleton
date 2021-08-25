package app

import (
	"encoding/json"
	"fmt"
	"github.com/joan41868/go-app-skeleton/src/model"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joan41868/go-app-skeleton/src/repository"
	"github.com/joan41868/go-app-skeleton/src/server"
	"github.com/joho/godotenv"
)

type Application struct {
	repository repository.Repository
	server     *server.Server
	logger     *log.Logger

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
		logger:     log.New(os.Stdout, fmt.Sprintf("[%s]", os.Getenv("APP_NAME")), 0),
	}

	return app
}

// Setup should perform additional app configuration/loading/etc. Any cross-package logic can be included here as well
func (app *Application) Setup() {
	// example:
	app.server.AddRoute("/hello", "GET", func(wr http.ResponseWriter, r *http.Request) {
		wr.Write([]byte("Hello world!"))
	})

	// example 2:
	app.server.AddRoute("/user", "POST", func(wr http.ResponseWriter, r *http.Request) {
		var usr model.Entity
		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			app.logger.Println(err)
		}
		saved, err := app.repository.Create(&usr)
		if err != nil {
			app.logger.Println(err)
		}
		json.NewEncoder(wr).Encode(saved)
	})
}

func (app Application) Start() {
	app.servicesWg.Add(1)
	go app.server.Start()

	app.servicesWg.Wait()
}
