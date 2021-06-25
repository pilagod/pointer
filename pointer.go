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

// NewString converts string to pointer
func NewString(value string) *string {
	return New(value).(*string)
}

// NewInt converts int to pointer
func NewInt(value int) *int {
	return New(value).(*int)
}

// NewUint converts uint to pointer
func NewUint(value uint) *uint {
	return New(value).(*uint)
}

// NewFloat converts float to pointer
func NewFloat(value float64) *float64 {
	return New(value).(*float64)
}

// NewTime converts time to pointer
func NewTime(value time.Time) *time.Time {
	return New(value).(*time.Time)
}
