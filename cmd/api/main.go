package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sibelephant/workout-plan-api/internal/database"
	"github.com/sibelephant/workout-plan-api/internal/handlers"
)

func main() {
	//Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer database.Disconnect()

	//initialize the router
	r := mux.NewRouter()

	//Define API routes
	r.HandleFunc("/workout-plans", handlers.CreateWorkoutPlan).Methods("POST")
	r.HandleFunc("/workout-plans", handlers.GetWorkoutPlans).Methods("GET")
	r.HandleFunc("/workout-plans/{id}", handlers.GetWorkoutPlanByID).Methods("GET")
	r.HandleFunc("/workout-plans/{id}", handlers.UpdateWorkoutPlan).Methods("PUT")
	r.HandleFunc("/workout-plans/{id}", handlers.DeleteWorkoutPlan).Methods("DELETE")

	// Exercise routes
	r.HandleFunc("/workout-plans/{id}/exercises", handlers.AddExercise).Methods("POST")
	r.HandleFunc("/workout-plans/{id}/exercises", handlers.GetExercises).Methods("GET")
	r.HandleFunc("/exercises/{exerciseId}", handlers.DeleteExercise).Methods("DELETE")

	//start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
