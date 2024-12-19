package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"worldboxing/internal/core"
	"worldboxing/lib/utils"
)

func Start() {
	write("Welcome to World Boxing!\n")

	for {
		input := read()
		if len(input) == 0 {
			continue
		}
		call, e := parseInput(input)
		if e != nil {
			throwError(e)
			continue
		}
		processCall(call)
	}
}

func processCall(call *Call) {
	write(fmt.Sprintf("Test execute %s", call.Command))
}

type Call struct {
	Raw     string
	Command string
	Args    []string
	Kwargs  map[string]string
}

func parseInput(input string) (*Call, *utils.Error) {
	call := Call{
		input,
		"",
		[]string{},
		map[string]string{},
	}

	fields := strings.Fields(input)
	if len(fields) == 0 {
		return nil, utils.NewError(core.CodeCliCallParseError)
	}
	for i, field := range fields {
		if i == 0 {
			call.Command = field
			continue
		}
		if strings.Contains(field, "=") {
			parts := strings.Split(field, "=")
			if len(parts) != 2 {
				return nil, utils.NewError(core.CodeCliCallParseError)
			}
			call.Kwargs[parts[0]] = parts[1]
			continue
		}
		call.Args = append(call.Args, field)
	}
	return &call, nil
}

func throwError(e *utils.Error) {
	message := e.Error()
	write(message)
}

func write(data string) {
	print(data)
}

func read() string {
	input := bufio.NewScanner(os.Stdin)
	write("> ")
	ok := input.Scan()
	if !ok {
		fmt.Print("\n")
		os.Exit(0)
	}
	return input.Text()
}
