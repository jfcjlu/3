package script

import (
	"fmt"
	"go/token"
	"reflect"
	"strings"
)

// compileErr, and only compileErr will be caught by Compile and returned as an error.
type compileErr struct {
	pos token.Pos
	msg string
}

// implements error
func (c *compileErr) Error() string {
	return c.msg
}

// constructs a compileErr
func err(pos token.Pos, msg ...interface{}) *compileErr {
	return &compileErr{pos, fmt.Sprintln(msg...)}
}

// type string for value i
func typ(i interface{}) string {
	return reflect.TypeOf(reflect.ValueOf(i).Interface()).String()
}

func assert(test bool) {
	if !test {
		panic("assertion failed")
	}
}

// decodes a token position in source to a line number
// and returns the line number + line code.
func pos2line(pos token.Pos, src string) string {
	if pos == 0 {
		return ""
	}
	lines := strings.Split(src, "\n")
	line := 0
	for i, b := range src {
		if token.Pos(i) == pos {
			return fmt.Sprint("line ", line+1, ": ", strings.Trim(lines[line], " \t")) // lines count from 1
		}
		if b == '\n' {
			line++
		}
	}
	return fmt.Sprint("position", pos) // we should not reach this
}
