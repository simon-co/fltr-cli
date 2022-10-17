package main

import (
	"io"
	"os"

	"github.com/simon-co/fltr-cli/internal/cmd"
)

func main(){
    if err := run(os.Args[1:], os.Stdout); err != nil {
        panic(err)
    }
}

func run(args []string, output io.Writer) error{
    return nil
}
