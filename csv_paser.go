package main

import (
	"bufio"
	"bytes"
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
		// Сut out the first element because it's a full regex
		slElemInLine = slElemInLine[1:]
		slTemplate[countIteration] = slElemInLine

		countIteration += 1
	}
	buf := bytes.Buffer{}
	err = c.template.Execute(&buf, slTemplate)
	if err != nil {
		return "", err
	}
	resHtml = buf.String()

	return resHtml, err
}
