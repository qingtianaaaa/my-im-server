package program

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExitWithError(err error) {
	progName := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "%s exit with -1: %v\n", progName, err)
	os.Exit(-1)
}

func GetProgName() string {
	args := os.Args
	if len(args) > 0 {
		segments := strings.Split(args[0], "/") //os.Args[0] is the program name
		if len(segments) > 0 {
			return segments[len(segments)-1]
		}
	}
	return ""
}

func SIGTERMExit() {
	programName := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Warning %s receive process terminal SIGTERM signal\n", programName)
}