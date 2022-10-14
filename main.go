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

func writeStrToFile(filePath string, target string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("can`t Create file in writeStrToFile(). Error is: %w", err)
	}

	_, err = file.WriteString(target)
	if err != nil {
		return fmt.Errorf("can`t write in file in writeStrToFile(). Error is: %w", err)
	}

	return nil
}

// Used named return values to return error from defer
func dataFromFile(filePath string) (res []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("can`t open file in dataFromFile(). Error is: %w", err)
	}
	defer func() {
		err = file.Close()
	}()

	res, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("can`t ReadAll file in dataFromFile(). Error is: %w", err)
	}

	return res, nil
}

func scannerFromFile(filePath string) (scanner *bufio.Scanner, err error) {
	fileBytes, err := dataFromFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error is: %w in scannerFromFile()", err)
	}

	// Used charmap.ISO8859_1 because data.css and data.prn encoded iso8859-1
	reader := transform.NewReader(bytes.NewReader(fileBytes), charmap.ISO8859_1.NewDecoder())
	scanner = bufio.NewScanner(reader)

	return scanner, nil
}

func regexpFromFile(filePath string) (*regexp.Regexp, error) {
	fileBytes, err := dataFromFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error is: %w in regexpFromFile()", err)
	}

	resRegexp, err := regexp.Compile(string(fileBytes))
	if err != nil {
		return nil, fmt.Errorf("can`t Compile regexp from file in regexpFromFile(). Error is: %w", err)
	}

	return resRegexp, nil
}

func templateFromFile(filePath string) (*template.Template, error) {
	fileBytes, err := dataFromFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error is: %w in templateFromFile()", err)
	}

	resTemplate := template.New("templateMain")
	resTemplate, err = resTemplate.Parse(string(fileBytes))
	if err != nil {
		return nil, fmt.Errorf("can`t create template in templateFromFile. Error is: %w", err)
	}

	return resTemplate, nil
}

func Parse(filePath string) error {
	var resHtml string

	ext := filepath.Ext(filePath)
	switch ext {
	case ".csv":
		regexpHeader, err := regexpFromFile("regexp_source/regexp_Header")
		if err != nil {
			return fmt.Errorf("error is: %w in Parse() case(csv)", err)
		}

		regexpMain, err := regexpFromFile("regexp_source/regexp_Main")
		if err != nil {
			return fmt.Errorf("error is: %w in Parse() case(csv)", err)
		}

		scanner, err := scannerFromFile(filePath)
		if err != nil {
			return fmt.Errorf("error is: %w in Parse() case(csv)", err)
		}

		templateMain, err := templateFromFile("template_source/template_main")
		if err != nil {
			return fmt.Errorf("error is: %w in Parse() case(csv)", err)
		}

		parserCsv := CsvParser{regexpMain: regexpMain,
			regexpHeader: regexpHeader,
			scanner:      scanner,
			template:     templateMain}
		resHtml, err = parserCsv.parseToHtml()
		if err != nil {
			return fmt.Errorf("error is: %w in Parse() case(csv)", err)
		}

		err = writeStrToFile("result_html/csv_table.html", resHtml)
		if err != nil {
			return fmt.Errorf("error is: %w in Parse() case(csv)", err)
		}
	case ".prn":
		return fmt.Errorf("prn not yet working")
	default:
		return fmt.Errorf("func Parse can handle 'csv' and 'prn' files, extension %v is not implemented", ext)
	}

	return nil
}

func main() {
	err := Parse("./data_source/data.csv")
	if err != nil {
		log.Printf("Error is: %+v\n", err)
		return
	}
}
