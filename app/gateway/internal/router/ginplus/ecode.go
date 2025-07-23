package ginplus

import (
	"strconv"
)

var (
	_codes = map[int]string{}
)

type Code struct {
	code    int
	message string
}

func (c Code) Error() string {
	return c.Message()
}

// Code return error code
func (c Code) Code() int { return c.code }

// Message return error message
func (c Code) Message() string {
	if c.message != "" {
		return c.message
	}
	if m, ok := _codes[c.code]; ok {
		return m
	}
	return strconv.Itoa(c.code)
}

// ResetMessage reset error message
func (c Code) ResetMessage(message string) error {
	c.message = message
	return c
}
