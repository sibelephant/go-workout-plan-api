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

// CreateWorkoutPlan handles the creation of a new workout plan
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

// AddExercise handles adding a new exercise to a workout plan
func AddExercise(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workoutPlanId := params["id"]

	var exercise models.Exercise
	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// First verify that the workout plan exists
	_, err := database.Client.WorkoutPlan.FindUnique(
		db.WorkoutPlan.ID.Equals(workoutPlanId),
	).Exec(context.Background())

	if err != nil {
		http.Error(w, "Workout plan not found", http.StatusNotFound)
		return
	}

	// Create the exercise and link it to the workout plan
	createdExercise, err := database.Client.Exercise.CreateOne(
		db.Exercise.Name.Set(exercise.Name),
		db.Exercise.Sets.Set(exercise.Sets),
		db.Exercise.Reps.Set(exercise.Reps),
		db.Exercise.WorkoutPlan.Link(
			db.WorkoutPlan.ID.Equals(workoutPlanId),
		),
	).Exec(context.Background())

	if err != nil {
		http.Error(w, "Failed to create exercise", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdExercise)
}

// GetExercises handles fetching all exercises for a specific workout plan
func GetExercises(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workoutPlanId := params["id"]

	// First verify that the workout plan exists
	_, err := database.Client.WorkoutPlan.FindUnique(
		db.WorkoutPlan.ID.Equals(workoutPlanId),
	).Exec(context.Background())

	if err != nil {
		http.Error(w, "Workout plan not found", http.StatusNotFound)
		return
	}

	// Get all exercises for the workout plan
	exercises, err := database.Client.Exercise.FindMany(
		db.Exercise.WorkoutPlanID.Equals(workoutPlanId),
	).Exec(context.Background())

	if err != nil {
		http.Error(w, "Failed to fetch exercises", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercises)
}

// DeleteExercise handles deleting an exercise by ID
func DeleteExercise(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	exerciseId := params["exerciseId"]

	_, err := database.Client.Exercise.FindUnique(
		db.Exercise.ID.Equals(exerciseId),
	).Delete().Exec(context.Background())

	if err != nil {
		http.Error(w, "Exercise not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
