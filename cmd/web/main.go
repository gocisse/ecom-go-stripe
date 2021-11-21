package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"
const cssVersion = "1"

// config is holding the configs of our app
// port to listen on
// env enviroment dev or production
// api what url to call for my backend api
// dsn  how to connect to the database
type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

// application the receiver of our app
type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", app.config.port),
		Handler:     app.route(),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	app.infoLog.Printf("Starting HTTP server in %s mode on port %d", app.config.env, app.config.port)

	return srv.ListenAndServe()
}

func main() {
	var cfg config
	
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen om")
	flag.StringVar(&cfg.env, "env", "development", "Application enviroment {developement |  production")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to api")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// make template cache
	tc := make(map[string]*template.Template)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		templateCache: tc,
		version:  version,
	}

	err := app.serve()
	if err != nil {
		
		app.errorLog.Println(err)
		panic(err)
	}
}
