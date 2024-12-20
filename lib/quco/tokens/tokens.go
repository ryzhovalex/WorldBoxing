package tokens

type Type int

var TextToTokenType = map[string]Type{
	"CREATE": Create,
	"SET":    Set,
	"DELETE": Delete,
	"THEN":   Then,
	"=":      Assignment,
	"NE":     Ne,
	"LT":     Lt,
	"LE":     Le,
	"GT":     Gt,
	"GE":     Ge,
	"IN":     In,
	"\"":     Quote,
	"TRUE":   True,
	"FALSE":  False,
	"(":      ContainerOpen,
	")":      ContainerClose,
	".":      Dot,
	",":      Comma,
}

const (
	Get Type = iota
	Create
	Set
	Delete
	Then
	// Generic name, could be a field name, or the string.
	Name
	// We don't have EQ operator since by default all operations are equality
	// (if applicable). Operator `=` in Quco is Assignment.
	Assignment
	Ne
	In
	Lt
	Le
	Gt
	Ge
	Quote
	Integer
	Float
	True
	False
	ContainerOpen
	ContainerClose
	Newline
	Dot
	Comma
)

type Token struct {
	Type  Type
	Value string
}
