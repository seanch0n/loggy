package main

import (
	"bytes"
	"io"
	"testing"
)

func TestProcessLine(t *testing.T) {
	checkRet := func(t testing.TB, line string, delim string, ret string) {
		t.Helper()
		buf := &bytes.Buffer{}
		var out io.Writer
		out = buf
		p := Processor{out, delim}
		p.processLine(line)
		if buf.String() != ret {
			t.Error("UHHHHGH")
		}
	}

	t.Run("basic :", func(t *testing.T) {
		str := "Comms::ConnectClient: client Connected"
		checkRet(t, str, "::", str+"\n")
	})
	t.Run("No delim found", func(t *testing.T) {
		checkRet(t, "Comms::ConnectClient", "-", "")
	})

}
