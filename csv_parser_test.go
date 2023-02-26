package vilin_parser

import (
	"bufio"
	"html/template"
	"regexp"
	"strings"
	"testing"
)

type TestCsvParse struct {
	Parser   *CsvParser
	TestHtml string
	TestErr  string
}

func TestErrorParseCsvToHtml(t *testing.T) {
	testingCase := []TestCsvParse{
		{Parser: nil, TestHtml: "", TestErr: ErrorNilInParseToHtml},
		{Parser: &CsvParser{regexpHeader: nil}, TestHtml: "", TestErr: ErrorNilInParseToHtml},
		{Parser: &CsvParser{regexpHeader: &regexp.Regexp{}, regexpMain: nil}, TestHtml: "", TestErr: ErrorNilInParseToHtml},
		{Parser: &CsvParser{regexpHeader: &regexp.Regexp{}, regexpMain: &regexp.Regexp{}, template: nil}, TestHtml: "", TestErr: ErrorNilInParseToHtml},
		{Parser: &CsvParser{regexpHeader: &regexp.Regexp{}, regexpMain: &regexp.Regexp{}, template: template.New(" "), scanner: nil}, TestHtml: "", TestErr: ErrorNilInParseToHtml},
		{Parser: &CsvParser{regexpHeader: &regexp.Regexp{}, regexpMain: &regexp.Regexp{}, template: template.New(" "), scanner: &bufio.Scanner{}}, TestHtml: "", TestErr: ErrorTemplate},
	}

	for number, val := range testingCase {
		_, resErr := val.Parser.parseToHtml()
		if !strings.Contains(resErr.Error(), val.TestErr) {
			t.Errorf("unexpected error in case [%v] \n Expected: %v \n Got: %v", number, val.TestErr, resErr)
		}
	}
}
