package main

import (
	"log"
	"net/http"

	"github.com/devsouzx/simple-go-mod/config"
	"github.com/devsouzx/simple-go-mod/handlers"
	"github.com/devsouzx/simple-go-mod/models"
	"github.com/gorilla/mux"
)

func main() {
	db := config.SetupDB()
	defer db.Close()

	_, err := db.Exec(models.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	//test2

	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(db)

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}