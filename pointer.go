package pointer

import (
	"reflect"
	"time"
)

// New converts value to pointer
func New(value interface{}) interface{} {
	rv := reflect.ValueOf(value)
	if rv.Type().Kind() == reflect.Ptr {
		return value
	}
	rp := reflect.New(rv.Type())
	rp.Elem().Set(rv)
	return rp.Interface()
}

// String converts string to pointer
func String(value string) *string {
	return New(value).(*string)
}

// Int converts int to pointer
func Int(value int) *int {
	return New(value).(*int)
}

// Uint converts uint to pointer
func Uint(value uint) *uint {
	return New(value).(*uint)
}

// Float converts float to pointer
func Float(value float64) *float64 {
	return New(value).(*float64)
}

// Time converts time to pointer
func Time(value time.Time) *time.Time {
	return New(value).(*time.Time)
}
