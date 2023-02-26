package vilin_parser

import (
	"bufio"
	"html/template"
	"strings"
	"testing"
)

type TestPrnParse struct {
	Parser  *PrnParser
	TestErr string
}

func TestErrorParsePrnToHtml(t *testing.T) {
	testingCase := []TestPrnParse{
		{Parser: nil, TestErr: ErrorNilInParseToHtml},
		{Parser: &PrnParser{template: nil}, TestErr: ErrorNilInParseToHtml},
		{Parser: &PrnParser{slRanges: []columnRange{{}}, scanner: nil, template: template.New("")}, TestErr: ErrorNilInParseToHtml},
		{Parser: &PrnParser{slRanges: nil, scanner: &bufio.Scanner{}, template: template.New("")}, TestErr: ErrorSlRangeEmpty},
		{Parser: &PrnParser{slRanges: []columnRange{{}}, scanner: &bufio.Scanner{}, template: template.New("")}, TestErr: ErrorTemplate},
	}

	for number, val := range testingCase {
		_, resErr := val.Parser.parseToHtml()
		if !strings.Contains(resErr.Error(), val.TestErr) {
			t.Errorf("unexpected error in case [%v] \n Expected: %v \n Got: %v", number, val.TestErr, resErr)
		}
	}
}
