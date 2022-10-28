package command

import "errors"

var (
    ErrNoCmd = errors.New("No command supplied")
    ErrInvalidCmd = errors.New("Invalid command")
)
