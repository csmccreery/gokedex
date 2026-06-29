package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type REPL interface {
	Start(prompt string) error
	Stop(msg string) error
	Help(msg string) error
	Execute(arg CommandArg) error
	SetCommands(commands map[string]Command) error
}

type Repl struct {
	Name        string
	Description string
	StopMessage string
	HelpMessage string
	History     map[string]string
	Commands    map[string]Command
}

type CommandArg struct {
	Name string
	Args []string
}

type Command struct {
	Name        string
	Description string
	Callback    func(arg CommandArg) error
}

func getFields(text string) (string, []string) {
	res := strings.Fields(text)
	for idx := range res {
		res[idx] = strings.ToLower(res[idx])
	}

	return res[0], res[1:]
}

func (r *Repl) Start(prompt string) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			fmt.Println()
			break
		}

		text := scanner.Text()
		if len(strings.TrimSpace(text)) == 0 {
			continue
		}

		name, args := getFields(text)
		switch name {
		case "exit":
			return r.Stop(r.StopMessage)
		case "help":
			err := r.Help(r.HelpMessage)
			if err != nil {
				return err
			}

			continue
		}

		arg := CommandArg{
			Name: name,
			Args: args,
		}

		err := r.Execute(arg)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}
	return nil
}

func (r *Repl) Stop(msg string) error {
	fmt.Printf("Stopping %s: %s", r.Name, msg)
	os.Exit(0)
	return nil
}

func (r *Repl) SetCommands(commands map[string]Command) error {
	r.Commands = commands
	return nil
}

func (r *Repl) Execute(arg CommandArg) error {
	if com, ok := r.Commands[arg.Name]; ok {
		err := com.Callback(arg)
		if err != nil {
			return err
		}
	}
	err := fmt.Errorf("Could not execute %s, not a command found in REPL", arg.Name)
	return err
}

func (r *Repl) Help(msg string) error {
	fmt.Printf("%s\n", msg)
	if r.Commands == nil {
		return fmt.Errorf("No commands registered to this REPL. Exiting")
	}
	for _, com := range r.Commands {
		fmt.Printf("%s : %s\n", com.Name, com.Description)
	}
	return nil
}
