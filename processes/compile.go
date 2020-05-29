package processes

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Compile ... Take the sbin and compile it to pure binary
func Compile(fName string) string {
	sbin, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatal(err)
	}
	chunks := strings.Split(string(sbin), ":")
	trimmedChunks := chunks[1 : len(chunks)-1]
	var bit string
	binary := []string{}
	for _, chunk := range trimmedChunks {
		if strings.HasSuffix(chunk, "?") {
			bit = "0"
		} else {
			bit = "1"
		}
		instances, err := strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(chunk, "?"), "!"))
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < instances; i++ {
			binary = append(binary, bit)
		}
	}
	return strings.Join(binary, "")
}
