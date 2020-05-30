package main

import (
	"os"
	"strings"

	"github.com/Matt-Gleich/shortbin/compression"
)

func main() {
	fName := strings.TrimSpace(os.Args[1])
	if strings.HasSuffix(fName, ".sbin") {
		compression.Decompress(fName)
	} else {
		compression.Compress(fName)
	}
}
