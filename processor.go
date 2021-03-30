package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

type Processor struct {
	writer  io.Writer
	delim   string
	outFile string
}

// attempt to find the delimeter in our string. If it is found, we can print the line
func (p *Processor) processLine(line string) (err error) {
	match, _ := regexp.MatchString(p.delim, line)
	if match {
		fmt.Fprint(p.writer, line, "\n")
		f, err := os.OpenFile(p.outFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(line + "\n"); err != nil {
			panic(err)
		}
	}
	return nil
}
