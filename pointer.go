package pointer

// any types

func Of[T any](v T) *T {
	return &v
}

func Value[T any](v *T, def T) T {
	if v == nil {
		return def
	}

	return *v
}

func Equal[T comparable](a *T, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil {
		return *a == *b
	}

	return false
}

func EqualValue[T comparable](p *T, v T) bool {
	if p == nil {
		return false
	}

	return *p == v
}

// non autoresolveable types

func Int8(v int8) *int8 {
	return Of(v)
}

func Int16(v int16) *int16 {
	return Of(v)
}

func Int32(v int32) *int32 {
	return Of(v)
}

func Int64(v int64) *int64 {
	return Of(v)
}

func UInt8(v uint8) *uint8 {
	return Of(v)
}

func UInt16(v uint16) *uint16 {
	return Of(v)
}

func UInt32(v uint32) *uint32 {
	return Of(v)
}

func UInt64(v uint64) *uint64 {
	return Of(v)
}

func Float32(v float32) *float32 {
	return Of(v)
}

func Float64(v float64) *float64 {
	return Of(v)
}

func Complex64(v complex64) *complex64 {
	return Of(v)
}

func Complex128(v complex128) *complex128 {
	return Of(v)
}
