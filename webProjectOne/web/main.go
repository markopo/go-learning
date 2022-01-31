package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/markopo/go-learning/pkg/config"
	"github.com/markopo/go-learning/pkg/handlers"
	"github.com/markopo/go-learning/pkg/render"
	"log"
	"net/http"
	"strings"
	"time"
)

const portNumber = ":6969"
const useCache = true
const inProduction = false

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = inProduction

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatalln("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = useCache

	repo := handlers.NewRepo(&app)

	render.NewTemplates(&app)
	handlers.NewHandlers(repo)

	fmt.Println(fmt.Sprintf("Starting the app at port%s", portNumber))

	if strings.Contains(portNumber, "69")  {
		fmt.Println("It's the dirty port!")
	}

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
