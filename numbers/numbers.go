package numbers

import (
	"fmt"
	"regexp"
)

var NumberTypes = map[string]NumberType{
	",": REAL,
	".": NATURAL,
	"E": EXPONENTIAL,
	"%": PORCENTAGE,
}

const (
	REAL        = "REAL"
	NATURAL     = "NATURAL"
	ILLEGAL     = "ILLEGAL"
	EXPONENTIAL = "EXPONENTIAL"
	PORCENTAGE  = "PORCENTAGE"
)

type NumberType string

type Number struct {
	Id     int
	Number string
	Type   string
}

func (n *Number) String() string {
	return fmt.Sprintf("No.:%d Cadena: %s Tipo: %s", n.Id, n.Number, n.Type)
}

func NewNumber(n string, id int) *Number {
	t := MatchNumberType(n)
	return &Number{Id: id, Number: n, Type: t}
}

func MatchNumberType(n string) string {
	//TODO Descarta todo lo que sean numeros, y guardar los caracteres que sean en una variable y ya luego meter esto a un switch
	//TODO Usar un hashmap para identificar el tipo de numero
	// TODO optimizar esto d manera sexy

	r, _ := regexp.Compile("[^0-9E,.%]")
	if r.Match([]byte(n)) {
		return ILLEGAL
	}

	r, _ = regexp.Compile("E\\d")
	if r.Match([]byte(n)) {
		return EXPONENTIAL
	}
	r, _ = regexp.Compile("%")
	if r.Match([]byte(n)) {
		return PORCENTAGE
	}
	r, _ = regexp.Compile("[\\.\\,]")
	if r.Match([]byte(n)) {
		r, _ = regexp.Compile("\\d+\\.\\d+\\.")
		if r.Match([]byte(n)) {
			return ILLEGAL
		}
		return REAL
	} else {
		r, _ = regexp.Compile("[^a-zA-Z]")
		if r.Match([]byte(n)) {
			return NATURAL
		} else {
			return ILLEGAL
		}
	}
}
