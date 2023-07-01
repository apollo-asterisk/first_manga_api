package main

import (
	"database/sql"
	"fmt"
	"github.com/render_manga_api/controller"
	"github.com/render_manga_api/model/repository"
	"log"
	"net/http"
	"os"
)

var (
	tr = repository.NewTitleRepository()
	tc = controller.NewTitleController(tr)

	pr = repository.NewPageRepository()
	pc = controller.NewPageController(pr)

	ro = controller.NewRouter(tc, pc)
)

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DbHost"), os.Getenv("DbPort"), os.Getenv("DbUser"), os.Getenv("DbPassword"), os.Getenv("DbName"))
	repository.Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	defer repository.Db.Close()
	err = repository.Db.Ping()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to db!")

	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/home", ro.HandleTitlesRequest)
	http.HandleFunc("/viewer", ro.HandlePagesRequest)
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
