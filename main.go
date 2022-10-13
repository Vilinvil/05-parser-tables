package main

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func Parse(filePath string) error {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".csv":
		regexpHeader, err := regexp.Compile(`(?P<name>.+?)\,(?P<address>.+?)\,(?P<postcode>.+?)\,(?P<mobile>.+?)\,(?P<limit>.+?)\,(?P<birthday>.+)`)
		if err != nil {
			return err
		}

		regexpMain, err := regexp.Compile(`(?P<name>\".+\")\,(?P<address>.+)\,(?P<postcode>.+?)\,(?P<mobile>.+?)\,(?P<limit>.+?)\,(?P<birthday>.+)`)
		if err != nil {
			return err
		}

		scanner, err := ScannerFromFile(filePath)
		if err != nil {
			return err
		}

		templateMain := template.Must(template.New("").Parse(`<table>{{range .}}
	<tr>{{range .}}
		<td>{{.Text}}{{if .IsNotLast}},{{end}}</td>{{end}}
	</tr>{{end}}
</table>`))

		parser := CsvParser{regexpMain: regexpMain,
			regexpHeader: regexpHeader,
			scanner:      scanner,
			template:     templateMain}
		html, err := parser.parseToHtml()
		if err != nil {
			return err
		}
		fmt.Print(html)
	case ".prn":
		return fmt.Errorf("prn not yet working")
	default:
		return fmt.Errorf("func Parse can handle 'csv' and 'prn' files, extension %v is not implemented", ext)
	}

	return nil
}

func ScannerFromFile(filePath string) (scanner *bufio.Scanner, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("can`t open file in ScannerFromFile. Error is: %w", err)
	}
	defer func() {
		err = file.Close()
	}()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("can`t ReadAll file in ScannerFromFile. Error is: %w", err)
	}

	// Use charmap.ISO8859_1.NewDecoder() because the data contains characters in the encoding of the ISO8859_1
	reader := transform.NewReader(bytes.NewReader(fileBytes), charmap.ISO8859_1.NewDecoder())
	scanner = bufio.NewScanner(reader)

	return scanner, nil
}

func main() {
	err := Parse("./data_source/data.csv")
	if err != nil {
		log.Printf("Error is: %+v\n", err)
		return
	}
}
