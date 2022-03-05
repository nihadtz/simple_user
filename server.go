package main

import (
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/julienschmidt/httprouter"
	"github.com/nihadtz/simple_user/pkg/controllers"
	"github.com/nihadtz/simple_user/services"
	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"
)

var (
	users controllers.UserController

	runas   = kingpin.Flag("runas", "Define deployment type and which configuration profile to read (dev, test, prod...)").Short('r').Default("dev").String()
	migrate = kingpin.Flag("migrate", "Migrate database to level defined in configuration").Short('m').Bool()
)

func init() {
	kingpin.Parse()
}

func main() {

	services.NewConfigurer(*runas, "conf/conf.yaml")
	services.NewLogger()
	services.NewRenderer()

	if *migrate {
		services.MigrateDB(*runas)
	}

	services.NewAccess(*runas)

	server := negroni.Classic()
	mux := httprouter.New()

	recovery := negroni.NewRecovery()
	recovery.PrintStack = false

	server.Use(recovery)

	server.UseHandler(mux)

	users := controllers.NewUserController()

	mux.POST("/register", users.UserCreate)
	mux.GET("/user/:id", users.RetrieveUserByID)
	mux.PUT("/user", users.UserUpdate)

	server.Run(":" + services.Configuration.Port)
	graceful.Run(":"+services.Configuration.Port, 10*time.Second, server)
}
