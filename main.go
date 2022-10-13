package vilin_parser

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"os"
)

func ScannerFromFile(filePath string) (scanner *bufio.Scanner, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("can`t open file in ScannerFromFile. Error is: %w", err)
	}
	defer func() {
		err = file.Close()
	}()

	// Use charmap.ISO8859_1.NewDecoder() because the data contains characters in the encoding of the ISO8859_1
	reader := transform.NewReader(file, charmap.ISO8859_1.NewDecoder())
	scanner = bufio.NewScanner(reader)

	return scanner, nil
}

func main() {

}
