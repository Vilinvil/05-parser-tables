package vilin_parser

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

const (
	ErrorNilInParseToHtml = "find nil in parseToHtml()"
	ErrorSlRangeEmpty     = "table can not have zero columns"
	ErrorTemplate         = "p.template.Execute not work"
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
		return nil, fmt.Errorf("%w in scannerFromFile()", err)
	}

	// Used charmap.ISO8859_1 because data.css and data.prn encoded iso8859-1
	reader := transform.NewReader(bytes.NewReader(fileBytes), charmap.ISO8859_1.NewDecoder())
	scanner = bufio.NewScanner(reader)

	return scanner, nil
}

func regexpFromFile(filePath string) (*regexp.Regexp, error) {
	fileBytes, err := dataFromFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%w in regexpFromFile()", err)
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
		return nil, fmt.Errorf("%w in templateFromFile()", err)
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

	scanner, err := scannerFromFile(filePath)
	if err != nil {
		return fmt.Errorf("%w in Parse()", err)
	}

	templateMain, err := templateFromFile("template_source/template_main")
	if err != nil {
		return fmt.Errorf("%w in Parse()", err)
	}

	ext := filepath.Ext(filePath)
	switch ext {
	case ".csv":
		regexpHeader, err := regexpFromFile("regexp_source/regexp_Header")
		if err != nil {
			return fmt.Errorf("%w in Parse() case(csv)", err)
		}

		regexpMain, err := regexpFromFile("regexp_source/regexp_Main")
		if err != nil {
			return fmt.Errorf("%w in Parse() case(csv)", err)
		}

		parserCsv := CsvParser{regexpMain: regexpMain,
			regexpHeader: regexpHeader,
			scanner:      scanner,
			template:     templateMain}
		resHtml, err = parserCsv.parseToHtml()
		if err != nil {
			return fmt.Errorf("%w in Parse() case(csv)", err)
		}

	case ".prn":
		slRanges := []columnRange{{0, 16},
			{16, 38},
			{38, 47},
			{47, 63},
			{63, 74},
			{74, 82}}

		parserPrn := PrnParser{
			slRanges: slRanges,
			scanner:  scanner,
			template: templateMain,
		}
		resHtml, err = parserPrn.parseToHtml()
		if err != nil {
			return fmt.Errorf("%w in Parse() case(prn)", err)
		}

	default:
		return fmt.Errorf("func Parse can handle 'csv' and 'prn' files, extension %v is not implemented", ext)
	}

	// Used ext[1:] because in extension first rune is '.'
	err = writeStrToFile(fmt.Sprintf("result_html/%v_table.html", ext[1:]), resHtml)
	if err != nil {
		return fmt.Errorf("%w in writeStrToFile() in Parse()", err)
	}

	return nil
}
