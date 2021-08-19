package matr

type sparseArray map[int]float64

type TransitionMatrix map[int]sparseArray
