package utils

type InfoLogger interface {
	Info(msg string)
}

type NUMBER interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}
