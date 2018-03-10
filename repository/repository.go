package repository

import (
	"errors"
	. "go-api/types"
)

var workouts = make([]Workout, 0)
var exercises = make([]Exercise, 0)
var programs = make([]Program, 0)

func GetWorkout(ID int) (*Workout, error) {
	if len(workouts) == 0 {
		populateWorkouts()
	}
	for _, workout := range workouts {
		if workout.ID == ID {
			return &workout, nil
		}
	}
	return nil, errors.New("Workout not found")
}

func GetExercise(ID int) (*Exercise, error) {
	if len(exercises) == 0 {
		populateExercises()
	}
	for _, exercise := range exercises {
		if exercise.ID == ID {
			return &exercise, nil
		}
	}
	return nil, errors.New("Exercise not found")
}

func GetProgram(ID int) (*Program, error) {
	if len(programs) == 0 {
		populatePrograms()
	}
	for _, program := range programs {
		if program.ID == ID {
			return &program, nil
		}
	}
	return nil, errors.New("Program not found")
}

func populateWorkouts() {
	workouts = append(workouts, exampleWorkout)
}

func populatePrograms() {
	programs = append(programs, greyskull)
}

func populateExercises() {
	exercises = append(exercises, exampleExercises...)
}

var exampleExercises = []Exercise{
	Exercise{ID: 1, Name: "Deadlift"},
	Exercise{ID: 2, Name: "Squat"},
	Exercise{ID: 3, Name: "Overhead Press"},
	Exercise{ID: 4, Name: "Bench Press"}}

var greyskull = Program{ID: 1, Name: "Greyskull", Workouts: []int{1, 2, 3, 4, 5, 6}}

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
