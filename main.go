package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"schoolApplication/handlers"
	"schoolApplication/repository"

	"schoolApplication/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize the database connection
	connStr := "user=postgres dbname=schoolApp sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	repo := repository.NewStudentRepository(db)
	handler := handlers.NewStudentHandler(repo)
	repoC := repository.NewClassRepository(db)
	handlerC := handlers.NewClassHandler(repoC)

	r := mux.NewRouter()
	r.HandleFunc("/student", handler.CreateStudent).Methods("POST")
	r.HandleFunc("/students", handler.GetStudents).Methods("GET")
	r.HandleFunc("/students/{id:[0-9]+}", handler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id:[0-9]+}", handler.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/class", handlerC.CreateClass).Methods("POST")
	r.HandleFunc("/class", handlerC.GetClasses).Methods("GET")
	r.HandleFunc("/classes/{id:[0-9]+}", handlerC.UpdateClass).Methods("PUT")
	r.HandleFunc("/classes/{id:[0-9]+}", handlerC.DeleteClass).Methods("DELETE")

	// Apply the middleware
	r.Use(middleware.LoggingMiddleware)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	http.ListenAndServe(":8080", r)

}
