package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sibelephant/workout-plan-api/internal/database"
	"github.com/sibelephant/workout-plan-api/internal/models"
	"github.com/sibelephant/workout-plan-api/prisma/db"
)

//CreateWorkoutPlan handles the creation of a new workout plan
func CreateWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	var plan models.WorkoutPlan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdPlan, err := database.Client.WorkoutPlan.CreateOne(db.WorkoutPlan.Name.Set(plan.Name), db.WorkoutPlan.Description.Set(*plan.Description)).Exec(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdPlan)
}

// GetWorkout Plans handles fetching all workout plans
func GetWorkoutPlans(w http.ResponseWriter, r *http.Request) {
	plans, err := database.Client.WorkoutPlan.FindMany().Exec(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(plans)
}

// UpdateWorkoutPlan handles updating a workout plan by ID
func UpdateWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var plan models.WorkoutPlan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedPlan, err := database.Client.WorkoutPlan.FindUnique(
		db.WorkoutPlan.ID.Equals(id),
	).Update(
		db.WorkoutPlan.Name.Set(plan.Name),
		db.WorkoutPlan.Description.Set(*plan.Description),
	).Exec(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedPlan)
}

// DeleteWorkoutPlan handles deleting a workout plan by ID
func DeleteWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := database.Client.WorkoutPlan.FindUnique(
		db.WorkoutPlan.ID.Equals(id),
	).Delete().Exec(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetWorkoutPlanByID handles fetching a single workout plan by ID
func GetWorkoutPlanByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    plan, err := database.Client.WorkoutPlan.FindUnique(
        db.WorkoutPlan.ID.Equals(id),
    ).Exec(context.Background())
    
    if err != nil {
        http.Error(w, "Workout plan not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(plan)
}