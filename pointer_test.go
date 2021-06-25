package pointer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestPointer(t *testing.T) {
	suite.Run(t, &pointerSuite{})
}

type pointerSuite struct {
	suite.Suite
}

func (s *pointerSuite) TestNewString() {
	value := "hello"

	s.Equal(&value, New(value).(*string))
	s.Equal(&value, NewString(value))
}

func (s *pointerSuite) TestNewInt() {
	value := 123

	s.Equal(&value, New(value).(*int))
	s.Equal(&value, NewInt(value))
}

func (s *pointerSuite) TestNewUint() {
	value := uint(123)

	s.Equal(&value, New(value).(*uint))
	s.Equal(&value, NewUint(value))
}

func (s *pointerSuite) TestNewFloat() {
	value := 123.45

	s.Equal(&value, New(value).(*float64))
	s.Equal(&value, NewFloat(value))
}

func (s *pointerSuite) TestNewTime() {
	value := time.Now()

	s.Equal(&value, New(value).(*time.Time))
	s.Equal(&value, NewTime(value))
}

func (s *pointerSuite) TestNewStruct() {
	type Value struct {
		Value string
	}
	value := Value{"hello"}

	s.Equal(&value, New(value).(*Value))
}

func (s *pointerSuite) TestNewPointer() {
	value := "hello"

	s.Equal(&value, New(&value).(*string))
}
