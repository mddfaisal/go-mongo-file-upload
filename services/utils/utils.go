package utils

import (
	"fmt"
	"runtime"
)

// Server server
type Server struct{}

// Trace to trace errors
func Trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s:%d %s", file, line, f.Name())
}
