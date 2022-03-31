package test

type Orderable[T any] interface {
	Cmp(other T) int
}

type Number interface {
	uint8 | uint16 | uint32 | uint64
	int8 | int16 | int32 | int64
	float32 | float64 | int
}
