package vilin_parser

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

// left border inclusive and right border not inclusive
type columnRange struct {
	lIndex int
	rIndex int
}

type PrnParser struct {
	slRanges []columnRange
	scanner  *bufio.Scanner
	template *template.Template
}

func (p *PrnParser) parseToHtml() (resHtml string, err error) {
	if p == nil {
		return "", fmt.Errorf("PrnParser == nil. Error is: %v", ErrorNilInParseToHtml)
	}
	if p.template == nil {
		return "", fmt.Errorf("PrnParser.template == nil. Error is: %v", ErrorNilInParseToHtml)
	}
	if p.scanner == nil {
		return "", fmt.Errorf("PrnParser.scanner == nil. Error is: %v", ErrorNilInParseToHtml)
	}
	if len(p.slRanges) == 0 {
		return "", fmt.Errorf("table can not have zero columns")
	}

	countIteration := 0
	slTemplate := make([][]string, 0)

	for p.scanner.Scan() {
		curLine := p.scanner.Text()
		curSlRune := []rune(curLine)

		slTemplate = append(slTemplate, []string{})

		for _, val := range p.slRanges {
			elText := strings.TrimSpace(string(curSlRune[val.lIndex:val.rIndex]))
			slTemplate[countIteration] = append(slTemplate[countIteration], elText)
		}

		countIteration += 1
	}
	buf := bytes.Buffer{}
	err = p.template.Execute(&buf, slTemplate)
	if err != nil {
		return "", err
	}
	resHtml = buf.String()

	return resHtml, err

}
