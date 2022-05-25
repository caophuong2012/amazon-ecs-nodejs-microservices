package pointer

import (
	"time"
)

// OfString returns the pointer to the given string
func OfString(str string) *string {
	return &str
}

// OfInt returns the pointer to the given int
func OfInt(i int) *int {
	return &i
}

// OfTime returns the pointer to the given time
func OfTime(t time.Time) *time.Time {
	return &t
}

// OfInt64 returns the pointer to the given int64
func OfInt64(i64 int64) *int64 {
	return &i64
}

// OfInt32 returns the pointer to the given int32
func OfInt32(i32 int32) *int32 {
	return &i32
}

// OfFloat64 returns the pointer to the given OfFloat64
func OfFloat64(f64 float64) *float64 {
	return &f64
}

// OfBool returns the pointer to the given bool
func OfBool(b bool) *bool {
	return &b
}

// Of returns the pointer to the given value
func Of(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return &v
}
