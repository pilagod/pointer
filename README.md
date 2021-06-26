# pointer [![Build Status](https://travis-ci.com/pilagod/pointer.svg?branch=master)](https://travis-ci.com/pilagod/pointer) [![Coverage Status](https://coveralls.io/repos/github/pilagod/pointer/badge.svg?branch=master)](https://coveralls.io/github/pilagod/pointer?branch=master)

Pointer utils for Golang

## Installation

```sh
go get -u github.com/pilagod/pointer
```

## Usage

```go
import (
    "time"

    "github.com/pilagod/pointer"
)

type Person struct {
    Name    *string
    Age     *int
    BornAt  *time.Time
}

func CreatePerson() Person {
    return Person{
        // use New to convert arbitrary value to pointer, and manually assert result to expected type
        Name: pointer.New("Andy").(*string),
        // use type helper to convert value to pointer without type assertion
        Age: pointer.Int(30),
        BornAt: pointer.Time(time.Now())
    }
}
```
