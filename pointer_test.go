package pointer

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOf(t *testing.T) {
	tests := []struct {
		name   string
		target any
	}{
		{
			name:   "string type",
			target: "test",
		},
		{
			name:   "int type",
			target: 2,
		},
		{
			name: "time type",
			target: func() time.Time {
				date, err := time.Parse(time.RFC3339, "2022-01-01T12:30:00Z")
				require.NoError(t, err)
				return date
			}(),
		},
		{
			name:   "uint32 type",
			target: uint32(4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Of(tt.target)
			require.True(t, reflect.ValueOf(result).Kind() == reflect.Ptr)
		})
	}
}

func TestValue(t *testing.T) {
	var v = 12.5
	if got := Value(Of(v), 0); got != v {
		t.Errorf("Value() = %v, want %v", got, &v)
	}
}

func TestValueWithDefault(t *testing.T) {
	var v = 12.5
	if got := Value(nil, v); got != v {
		t.Errorf("Value() = %v, want %v", got, &v)
	}
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		name string
		a    *int
		b    *int
		want bool
	}{
		{name: "a, b nil", want: true},
		{name: "a nill, b not", b: Of(1), want: false},
		{name: "a not, b nil", a: Of(1), want: false},
		{name: "a, b not equal", a: Of(1), b: Of(2), want: false},
		{name: "a, b equal", a: Of(1), b: Of(1), want: true},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.a, tt.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualValue(t *testing.T) {
	testCases := []struct {
		name string
		p    *int
		v    int
		want bool
	}{
		{name: "p nil", v: 123, want: false},
		{name: "p not nil, not eq", p: Of(999), v: 123, want: false},
		{name: "p not nil, eq", p: Of(123), v: 123, want: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualValue(tt.p, tt.v); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	var v = int8(1)
	if got := Int8(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("Int8() = %v, want %v", got, &v)
	}
}

func TestInt16(t *testing.T) {
	var v = int16(1)
	if got := Int16(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("Int16() = %v, want %v", got, &v)
	}
}

func TestInt32(t *testing.T) {
	var v = int32(1)
	if got := Int32(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("Int32() = %v, want %v", got, &v)
	}
}

func TestInt64(t *testing.T) {
	var v = int64(1)
	if got := Int64(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("Int64() = %v, want %v", got, &v)
	}
}

func TestUInt8(t *testing.T) {
	var v = uint8(1)
	if got := UInt8(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("UInt8() = %v, want %v", got, &v)
	}
}

func TestUInt16(t *testing.T) {
	var v = uint16(1)
	if got := UInt16(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("UInt16() = %v, want %v", got, &v)
	}
}

func TestUInt32(t *testing.T) {
	var v = uint32(1)
	if got := UInt32(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("UInt32() = %v, want %v", got, &v)
	}
}

func TestUInt64(t *testing.T) {
	var v = uint64(1)
	if got := UInt64(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("UInt64() = %v, want %v", got, &v)
	}
}

func TestFloat32(t *testing.T) {
	var v = float32(1)
	if got := Float32(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("Float32() = %v, want %v", got, &v)
	}
}

func TestFloat64(t *testing.T) {
	var v = float64(1)
	if got := Float64(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("Float64() = %v, want %v", got, &v)
	}
}

func TestComplex64(t *testing.T) {
	var v = complex64(1)
	if got := Complex64(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("Complex64() = %v, want %v", got, &v)
	}
}

func TestComplex128(t *testing.T) {
	var v = complex128(1)
	if got := Complex128(v); !reflect.DeepEqual(got, &v) {
		t.Errorf("Complex128() = %v, want %v", got, &v)
	}
}
