package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Alok-Majumder/inventory-assignment/internal/environment"
	"github.com/Alok-Majumder/inventory-assignment/internal/postgres"
	"github.com/Alok-Majumder/inventory-assignment/pkg/controller"
)

func main() {
	fmt.Println("Hello")

	ctx := context.Background()

	envs, err := environment.GetEnvironment(ctx) // Populate RunTime Variables
	if err != nil {
		log.Fatal("Environment Variable Not Set Up ", err)
	}

	postGresDB, err := postgres.NewPostGresConnPool(envs.DbUser, envs.DbPwd, envs.DbTCPHost, envs.DbPort, envs.DbName)
	if err != nil {
		log.Fatal("DB Connection Failed: ", err)
	}

	controller := controller.NewController(

		envs,
		ctx,
		postGresDB,
	)

	go func() {
		http.Handle("/products/", controller)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("Error starting the service ", err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
