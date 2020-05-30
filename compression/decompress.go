package compression

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Decompress ... Take the sbin and decompress it into pure binary
func Decompress(fName string) {
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
	file, err := os.OpenFile(
		"out",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, err := bitStringToBytes(strings.Join(binary, ""))
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(bytes)
}

func bitStringToBytes(s string) ([]byte, error) {
	b := make([]byte, (len(s)+(8-1))/8)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '1' {
			return nil, errors.New("value out of range")
		}
		b[i>>3] |= (c - '0') << uint(7-i&7)
	}
	return b, nil
}
