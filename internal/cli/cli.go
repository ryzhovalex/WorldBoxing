package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"worldboxing/internal/code"
	"worldboxing/lib/utils"

	"github.com/fatih/color"
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
		e = executeCall(call)
		if e != nil {
			throwError(e)
			continue
		}
	}
}

type Context struct {
	Call *Call
	// database ...
}

type CommandFunction func(*Context) *utils.Error

var commands = map[string]CommandFunction{}

func RegisterCommand(command string, function CommandFunction) *utils.Error {
	_, ok := commands[command]
	if ok {
		return utils.NewError(code.CliCommandAlreadyRegistered)
	}
	commands[command] = function
	return nil
}

func executeCall(call *Call) *utils.Error {
	function, ok := commands[call.Command]
	if !ok {
		return utils.NewError(code.CliNoSuchCommand)
	}
	var ctx = Context{
		call,
	}
	return function(&ctx)
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
		return nil, utils.NewError(code.CliCallParsing)
	}
	for i, field := range fields {
		if i == 0 {
			call.Command = field
			continue
		}
		if strings.Contains(field, "=") {
			parts := strings.Split(field, "=")
			if len(parts) != 2 {
				return nil, utils.NewError(code.CliCallParsing)
			}
			call.Kwargs[parts[0]] = parts[1]
			continue
		}
		call.Args = append(call.Args, field)
	}
	return &call, nil
}

type Writer struct{}

var writer = &Writer{}

func (w *Writer) Write(data []byte) (int, error) {
	print(string(data))
	return len(data), nil
}

func throwError(e *utils.Error) {
	write("[")
	color.New(color.FgRed).Fprint(writer, fmt.Sprintf("Error %d", e.Code()))
	write(fmt.Sprintf("] %s\n", utils.TranslateCode(e.Code())))
}

func write(data string) {
	writer.Write([]byte(data))
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
