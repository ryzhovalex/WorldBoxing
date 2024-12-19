package cli

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"worldboxing/internal/code"
	"worldboxing/lib/orwynn"
	"worldboxing/lib/utils"

	"github.com/fatih/color"
)

func Init() (orwynn.Transport, *utils.Error) {
	return &Transport{}, nil
}

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

// Single-connection CLI transport. Always has a single active connection.
type Transport struct {
	connection orwynn.Connection
}

func (transport *Transport) GetMaxConnectionSize() int {
	return 1
}
func (transport *Transport) GetConnectionSize() int {
	return 1
}
func (transport *Transport) GetConnection(
	id utils.Id,
) (orwynn.Connection, *utils.Error) {
	return transport.connection, nil
}
func (transport *Transport) Accept() (orwynn.Connection, *utils.Error) {
	if transport.connection == nil {
		transport.connection = &Connection{
			id: 0,
		}
	}
	return transport.connection, nil
}
func (transport *Transport) Close() {}

type Connection struct {
	id        utils.Id
	transport orwynn.Transport
}

func (connection *Connection) Id() utils.Id {
	return connection.id
}
func (connection *Connection) GetTransport() orwynn.Transport {
	return connection.transport
}
func (connection *Connection) Send(data []byte) *utils.Error {
	_, e := writer.Write(data)
	if e != nil {
		return utils.DefaultError()
	}
	return nil
}
func (connection *Connection) Recv() ([]byte, *utils.Error) {
	return []byte(read()), nil
}
func (connection *Connection) Close() {
}

type CommandFunction func(*Context) *utils.Error

var commands = map[string]CommandFunction{}

func RegisterCommand(command string, function CommandFunction) *utils.Error {
	_, ok := commands[strings.ToLower(command)]
	if ok {
		return utils.NewError(code.CliCommandAlreadyRegistered)
	}
	commands[command] = function
	return nil
}

func executeCall(call *Call) *utils.Error {
	function, ok := commands[strings.ToLower(call.Command)]
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

var CommandRegex = regexp.MustCompile(`^([A-z0-9]+\.)?[A-z0-9]+$`)
var KwargKeyRegex = regexp.MustCompile(`^[A-z0-9]+$`)

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
			command := field
			if !CommandRegex.MatchString(command) {
				return nil, utils.NewError(code.CliCallParsing)
			}
			call.Command = command
			continue
		}
		if strings.Contains(field, "=") {
			parts := strings.Split(field, "=")
			if len(parts) != 2 {
				return nil, utils.NewError(code.CliCallParsing)
			}
			kwargKey := parts[0]
			if !CommandRegex.MatchString(kwargKey) {
				return nil, utils.NewError(code.CliCallParsing)
			}
			call.Kwargs[kwargKey] = parts[1]
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
