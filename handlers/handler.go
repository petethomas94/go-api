package handlers

import (
	"encoding/json"
	"go-api/repository"
	. "go-api/types"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var WorkoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	if workoutID, err := strconv.Atoi(parameters["workoutId"]); err == nil {
		if workout, err := repository.GetWorkout(workoutID); err == nil {
			payload, _ := json.Marshal(workout)
			writeResponse(w, payload)
		} else {
			writeError(w, err)
		}
	}
})

var ExerciseHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	if exerciseID, err := strconv.Atoi(parameters["exerciseId"]); err == nil {
		if exercise, err := repository.GetExercise(exerciseID); err == nil {
			payload, _ := json.Marshal(exercise)
			writeResponse(w, payload)
		} else {
			writeError(w, err)
		}
	}
})

var ProgramHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	if programID, err := strconv.Atoi(parameters["programId"]); err == nil {
		if program, err := repository.GetProgram(programID); err == nil {
			payload, _ := json.Marshal(program)
			writeResponse(w, payload)
		} else {
			writeError(w, err)
		}
	}
})

var WorkoutSessionGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	if workoutSessionID, err := strconv.Atoi(parameters["workoutSessionId"]); err == nil {
		if workoutSession, err := repository.GetWorkoutSession(workoutSessionID); err == nil {
			payload, _ := json.Marshal(workoutSession)
			writeResponse(w, payload)
		} else {
			writeError(w, err)
		}
	}
})

var WorkoutSessionPostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	workoutSession := new(WorkoutSession)
	err = json.Unmarshal(body, workoutSession)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	repository.SaveWorkoutSession(workoutSession)

	w.Header().Set("Location", r.URL.Path+"/")
	w.WriteHeader(http.StatusCreated)
})

func writeError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusNotFound)
}

func writeResponse(w http.ResponseWriter, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
