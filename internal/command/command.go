package command

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Command struct {
    Name string
    Aliases []string
    runner
}

type runner interface {
    Run (output io.Writer, workingDir string, args []string) error
}

type CommandsMap map[string]*Command

var Commands = struct {
	CommandsMap
	CommandNames []string
	Mu           sync.Mutex
}{
	CommandsMap:  map[string]*Command{},
	CommandNames: []string{},
	Mu:           sync.Mutex{},
}

func (self *Command) RegisterCommand(){
    Commands.Mu.Lock()
    defer Commands.Mu.Unlock()
    Commands.CommandsMap[self.Name] = self
    Commands.CommandNames = append(Commands.CommandNames, self.Name)
    for _, alias := range self.Aliases{
        Commands.CommandsMap[alias] = self
    }
}

func (self CommandsMap) printValidCommands(output io.Writer){
    fmt.Fprintln(output, "Accepted Commands")
    for i, name := range Commands.CommandNames {
        fmt.Fprintf(output, "%d. %s\n", i, name)
    }
}

func Run(output io.Writer, args []string) error {
    if len(args) < 1 {
        fmt.Fprintln(output, ErrNoCmd)
        Commands.printValidCommands(output)
        return ErrNoCmd
    }
    cmd, ok := Commands.CommandsMap[args[0]]
    if !ok {
        fmt.Fprintln(output, ErrInvalidCmd)
        Commands.printValidCommands(output)
        return ErrInvalidCmd
    }
    cwd, _ := os.Getwd()
    return cmd.Run(output, cwd, args[1:])
}
