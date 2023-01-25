package main

import (
	"io"
	"log"
	"os"

	"github.com/simon-co/fltr-cli/internal/apperr"
	"github.com/simon-co/fltr-cli/internal/command"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		log.Println(apperr.Parse(err))
	}
}

func run(args []string, output io.Writer) error {
	return command.Run(output, args)
}
