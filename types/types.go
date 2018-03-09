package types

type Program struct {
	ID       int
	Name     string
	Workouts []int
}

type Workout struct {
	ID        int
	Exercises []ExerciseGroup
}

type ExerciseGroup struct {
	ExerciseID int
	Sets       []Set
}

type Set struct {
	Amrap bool
	Reps  int
}

type Exercise struct {
	ID   int
	Name string
}

type WorkoutSession struct {
	Workout Workout
	UserID  int
}
