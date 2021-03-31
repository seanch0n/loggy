package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"time"
)

func playbackLogFile(playbackSpeed int, playbackFile string, delim string, outFile string) bool {
	file, err := os.Open(playbackFile)
	if err != nil {
		return false
	}
	defer file.Close()
	p := Processor{
		writer:  os.Stdout,
		delim:   delim,
		outFile: outFile,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		time.Sleep(time.Duration(playbackSpeed) * time.Second)
		p.processLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return true
}

func main() {
	addr := flag.String("addr", "localhost", "bind ip")
	port := flag.Int("port", 5555, "bind port")
	delim := flag.String("delim", "", "What delimeter do you want to use? Pass empty string for no filtering")
	outFile := flag.String("outFile", "logfile.log", "logfile name")
	playback := flag.Bool("playback", false, "do you want to playback a log file")
	playbackSpeed := flag.Int("playbackSpeed", 1, "how fast to play back, in seconds. 0 will dump it all at once")
	playbackFile := flag.String("playbackFile", "log.log", "log file to playback")
	flag.Parse()

	if *playback {
		playbackLogFile(*playbackSpeed, *playbackFile, *delim, *outFile)
	} else {
		servStr := *addr + ":" + strconv.Itoa(*port)
		tcpServ, err := NewServer(servStr, *delim, *outFile)

		if err != nil {
			panic(err)
		}

		tcpServ.Run()

	}

}
