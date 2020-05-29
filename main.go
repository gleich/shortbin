package main

import (
	"os"
	"strings"

	"github.com/Matt-Gleich/shortbin/processes"
)

func main() {
	fName := strings.TrimSpace(os.Args[1])
	if strings.HasSuffix(fName, ".sbin") {
		processes.Compile(fName)
	}
}
