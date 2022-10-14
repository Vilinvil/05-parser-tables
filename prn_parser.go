package main

import (
	"bufio"
	"bytes"
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
	countIteration := 0
	slTemplate := make([][]string, 0)

	for p.scanner.Scan() {
		curLine := p.scanner.Text()
		curSlRune := []rune(curLine)

		slTemplate = append(slTemplate, []string{})

		for key, val := range p.slRanges {
			elText := strings.TrimSpace(string(curSlRune[val.lIndex:val.rIndex]))
			if key == len(p.slRanges)-2 {
				slTemplate[countIteration] = append(slTemplate[countIteration], elText)
			} else {
				slTemplate[countIteration] = append(slTemplate[countIteration], elText)
			}
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
