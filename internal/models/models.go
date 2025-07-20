package models

// WorkoutPlan represent the workout plan in the database
type WorkoutPlan struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Exercises []Exercise `json:"exercises,omitempty"`
}

//Exercse represents an exercise in a workout plan

type Exercise struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Sets int `json:"sets"`
	Reps int `json:"reps"`
	WorkoutPlan string `json:"workout_plan_id"`
}
