package main

import (
	"go-api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.Handle("/api/workout/{workoutId}", handlers.WorkoutHandler).Methods("GET")
	r.Handle("/api/workout", handlers.WorkoutHandler).Methods("GET")
	r.Handle("/api/exercise/{exerciseId}", handlers.ExerciseHandler).Methods("GET")
	r.Handle("/api/exercise", handlers.ExerciseHandler).Methods("GET")
	r.Handle("/api/program/{programId}", handlers.ProgramHandler).Methods("GET")
	r.Handle("/api/workoutsession", handlers.WorkoutSessionPostHandler).Methods("POST")
	r.Handle("/api/workoutsession/{workoutSessionId}", handlers.WorkoutSessionGetHandler).Methods("GET")

	http.ListenAndServe(":3000", r)
}
