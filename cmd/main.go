package main

import (
	"fmt"
	"github.com/StepanShevelev/task/pkg/api"
	cfg "github.com/StepanShevelev/task/pkg/config"
	mydb "github.com/StepanShevelev/task/pkg/db"
	"log"
	"net/http"
)

func main() {

	config := cfg.New()
	if err := config.Load("./configs", "config", "yml"); err != nil {
		log.Fatal(err)
	}

	db, err := mydb.New(config)
	if err != nil {
		log.Fatal(err)
	}
	db.SetDB()

	api.InitBackendApi(config)
	http.ListenAndServe(":"+config.Port, nil)
	fmt.Println("Done 0_o")
}
