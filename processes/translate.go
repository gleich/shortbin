package processes

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

// Translate ... Translate a binary file to a sbin file
func Translate(fName string) {
	bytes, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatal(err)
	}
	binaryChunks := []string{}
	for _, byteChunk := range bytes {
		var binaryChunkString string
		if strings.Contains(runtime.GOARCH, "64") {
			binaryChunkString = fmt.Sprintf("%064b", byteChunk)
		} else {
			binaryChunkString = fmt.Sprintf("%032b", byteChunk)
		}
		binaryChunks = append(binaryChunks, binaryChunkString)
	}

	// Actual Translation
	binaryString := strings.Join(binaryChunks, "")
	addedBits := []string{}
	sbinChunks := []string{}
	consecutiveInstances := 1
	for i, bit := range strings.Split(binaryString, "") {
		if i == 0 {
			addedBits = append(addedBits, bit)
			continue
		}
		addedBits = append(addedBits, bit)
		recent := addedBits[len(addedBits)-2]
		if recent == bit {
			consecutiveInstances++
		} else {
			var suffix string
			switch recent {
			case "0":
				suffix = "\\?"
			case "1":
				suffix = "\\!"
			}
			sbinChunks = append(sbinChunks, string(consecutiveInstances)+suffix)
			consecutiveInstances = 0
		}
		fmt.Printf("\ni: %#v\n", i)
		fmt.Printf("recent: %#v\n", recent)
		fmt.Printf("addedBits: %#v\n", addedBits)
		fmt.Printf("consecutiveInstances: %#v\n", consecutiveInstances)
		fmt.Printf("bit: %#v\n", bit)
	}
	fmt.Printf("sbinChunks: %#v\n", sbinChunks)
	f, err := os.Create(fName + ".sbin")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(":" + strings.Join(sbinChunks, ":") + ":")
}
