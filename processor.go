package main

import (
	"fmt"
	"io"
	"regexp"
)

type Processor struct {
	writer io.Writer
	delim  string
}

// attempt to find the delimeter in our string. If it is found, we can print the line
func (p *Processor) processLine(line string) (err error) {
	match, _ := regexp.MatchString(p.delim, line)
	if match {
		fmt.Fprint(p.writer, line, "\n")
	}
	return nil
}
