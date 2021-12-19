package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Alok-Majumder/inventory-load/internal/environment"
	"github.com/Alok-Majumder/inventory-load/internal/inventory"
	"github.com/Alok-Majumder/inventory-load/internal/postgres"
)

const mnumberOfArguments = 2

func main() {

	if len(os.Args) < mnumberOfArguments {
		fmt.Println("Too few arguments.")
	}
	if len(os.Args) > mnumberOfArguments {
		fmt.Println("Too many arguments.")

	}

	ctx := context.Background()

	envs, err := environment.GetEnvironment(ctx) // Populate RunTime Variables
	if err != nil {
		log.Fatal("Environment Variable Not Set Up ", err)
	}

	postGresDB, err := postgres.NewPostGresConnPool(envs.DbUser, envs.DbPwd, envs.DbTCPHost, envs.DbPort, envs.DbName)
	if err != nil {
		log.Fatal("DB Connection Failed: ", err)
	}

	//ctx := context.Background()
	fmt.Println("Execution started...")
	fileName := os.Args[1]

	if strings.Contains(fileName, "inventory") {

		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal("File reading error", err)

		}
		inv := new(inventory.InventoriesSrc)

		err = json.Unmarshal(data, &inv)
		checkError(err)
		i := inventory.NewService(inventory.NewPostGresRepo(postGresDB))
		i.Process(inv)

	} else if strings.Contains(fileName, "product") {
		//TODO
		fmt.Println(fileName)

	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal("Error is ....", err)

	}

}
