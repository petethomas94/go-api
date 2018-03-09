package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.Handle("/api/workout/{workoutId}", workoutHandler).Methods("GET")
	r.Handle("/api/workout", workoutHandler).Methods("GET")
	r.Handle("/api/exercise/{exerciseId}", exerciseHandler).Methods("GET")
	r.Handle("/api/exercise", exerciseHandler).Methods("GET")
	r.Handle("/api/program/{programId}", programHandler).Methods("GET")
	r.Handle("/api/workoutsession", workoutSessionHandler).Methods("POST")

	http.ListenAndServe(":3000", r)
}

var exampleWorkout = Workout{
	ID: 2,
	Exercises: []ExerciseGroup{
		ExerciseGroup{
			ExerciseID: 2,
			Sets: []Set{
				Set{Amrap: false, Reps: 5},
				Set{Amrap: false, Reps: 5},
				Set{Amrap: true, Reps: 5}}},
		ExerciseGroup{
			ExerciseID: 3,
			Sets: []Set{
				Set{Amrap: false, Reps: 5},
				Set{Amrap: false, Reps: 5},
				Set{Amrap: true, Reps: 5}}},
		ExerciseGroup{
			ExerciseID: 4,
			Sets: []Set{
				Set{Amrap: false, Reps: 5},
				Set{Amrap: false, Reps: 5},
				Set{Amrap: true, Reps: 5}}}}}

var exampleExercises = []Exercise{
	Exercise{ID: 1, Name: "Deadlift"},
	Exercise{ID: 2, Name: "Squat"},
	Exercise{ID: 3, Name: "Overhead Press"},
	Exercise{ID: 4, Name: "Bench Press"}}

var workoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	if workoutID, err := strconv.Atoi(parameters["workoutId"]); err == nil {
		workout := getWorkout(workoutID)
		payload, _ := json.Marshal(workout)
		writeResponse(w, payload)
	}
})

var exerciseHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	if exerciseID, err := strconv.Atoi(parameters["exerciseId"]); err == nil {
		if exercise, err := getExercise(exerciseID); err == nil {
			payload, _ := json.Marshal(exercise)
			writeResponse(w, payload)
		} else {

		}
	} else {
		payload, _ := json.Marshal(exampleExercises)
		writeResponse(w, payload)
	}
})

var programHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	if programID, err := strconv.Atoi(parameters["programId"]); err == nil {
		program := getProgram(programID)
		payload, _ := json.Marshal(program)
		writeResponse(w, payload)
	}
})

var workoutSessionHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	workoutSession := new(WorkoutSession)
	err = json.Unmarshal(body, workoutSession)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Print(workoutSession)

	w.Header().Set("Location", r.URL.Path+"/")
	w.WriteHeader(http.StatusCreated)
})

func getProgram(programID int) *Program {
	return &Program{ID: 1, Name: "Greyskull", Workouts: []int{1, 2, 3, 4, 5, 6}}
}

func getWorkout(workoutID int) *Workout {
	return &exampleWorkout
}

func getExercise(exerciseID int) (*Exercise, error) {
	for _, exercise := range exampleExercises {
		if exercise.ID == exerciseID {
			return &exercise, nil
		}
	}
	return nil, errors.New("Error: Exercise not found")
}

func writeError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusNotFound)
}

func writeResponse(w http.ResponseWriter, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
