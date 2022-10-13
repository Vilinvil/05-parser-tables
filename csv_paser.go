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

	// Open file and create scanner on top of it
	countIteration := 0
	type TemplateStruct struct {
		Text      string
		IsNotLast bool
	}
	slTemplate := make([][]TemplateStruct, 0)
	for c.scanner.Scan() {

		curLine := c.scanner.Text()
		slTemplate = append(slTemplate, []TemplateStruct{})
		slElemInLine := make([]string, 0)

		if isHeader {
			slElemInLine = c.regexpHeader.FindStringSubmatch(curLine)
			isHeader = false
		} else {
			slElemInLine = c.regexpMain.FindStringSubmatch(curLine)
		}

		// Сut out the first element because it's a full regex
		slElemInLine = slElemInLine[1:]

		for key, val := range slElemInLine {
			if key == len(slElemInLine)-1 {
				slTemplate[countIteration] = append(slTemplate[countIteration], TemplateStruct{Text: val, IsNotLast: false})
			} else {
				slTemplate[countIteration] = append(slTemplate[countIteration], TemplateStruct{Text: val, IsNotLast: true})
			}
		}
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
