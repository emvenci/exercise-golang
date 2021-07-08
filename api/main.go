package main

import (
	"exercise-backend/config"
	mysql "exercise-backend/infrastructure/db_mysql"
	"fmt"
	"log"
	"os"
)

type App struct{}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	// Load config
	if err = config.LoadConfig(); err != nil {
		return
	}
	// Handle Repository
	switch config.Conf.Repository {
	case "mysql":
		_, err = mysql.GetMysqlConn()
		if err != nil {
			return
		}
	default:
		log.Fatal("invalid repository")
	}
	app := App{}
	return app.Start()
}

func (a *App) Start() (err error) {
	return
}
