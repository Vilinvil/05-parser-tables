package vilin_parser

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"regexp"
)

type CsvParser struct {
	regexpHeader *regexp.Regexp
	regexpMain   *regexp.Regexp
	scanner      *bufio.Scanner
	template     *template.Template
}

func (c *CsvParser) parseToHtml() (resHtml string, err error) {
	if c == nil {
		return "", fmt.Errorf("CsvParser == nil. Error is: %v", ErrorNilInParseToHtml)
	}
	if c.regexpHeader == nil {
		return "", fmt.Errorf("CsvParser.regexpHeader == nil. Error is: %v", ErrorNilInParseToHtml)
	}
	if c.regexpMain == nil {
		return "", fmt.Errorf("CsvParser.regexpMain == nil. Error is: %v", ErrorNilInParseToHtml)
	}
	if c.template == nil {
		return "", fmt.Errorf("CsvParser.template == nil. Error is: %v", ErrorNilInParseToHtml)
	}
	if c.scanner == nil {
		return "", fmt.Errorf("CsvParser.scanner == nil. Error is: %v", ErrorNilInParseToHtml)
	}

	isHeader := true
	countIteration := 0
	slTemplate := make([][]string, 0)
	for c.scanner.Scan() {
		curLine := c.scanner.Text()

		slTemplate = append(slTemplate, []string{})
		slElemInLine := make([]string, 0)

		if isHeader {
			slElemInLine = c.regexpHeader.FindStringSubmatch(curLine)
			isHeader = false
		} else {
			slElemInLine = c.regexpMain.FindStringSubmatch(curLine)
		}

		if len(slElemInLine) == 0 {
			return "", fmt.Errorf("get line without elements in parseToHtml()")
		}
		// Сut out the first element because it's a full regex
		slElemInLine = slElemInLine[1:]
		slTemplate[countIteration] = slElemInLine

		countIteration += 1
	}
	buf := bytes.Buffer{}
	err = c.template.Execute(&buf, slTemplate)
	if err != nil {
		return "", fmt.Errorf("%v in Parse() on CsvParser. Error is: %v", ErrorTemplate, err)
	}
	resHtml = buf.String()

	return resHtml, nil
}
