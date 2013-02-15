package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var replayKeys []byte

func loadReplay(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Failed loading replay file %s: %v", filename, err)
		return
	}
	lr := newLineReader(bytes.NewReader(data))

	// First line is settings.
	line, _ := lr.Next()
	for _, kv := range strings.Split(line, " ") {
		kv = strings.TrimSpace(kv)
		parts := strings.SplitN(kv, "=", 2)
		switch k, v := parts[0], parts[1]; k {
		case "seed":
			x, err := strconv.ParseUint(v, 0, 32)
			if err != nil {
				log.Printf("Bad seed %q: %v", v, err)
				continue
			}
			debugf("seed set to %d", x)
			seedrand(uint32(x))
		default:
			log.Printf("Unknown replay setting %q; ignoring", k)
		}
	}

	// All remaining lines are commands.
	q := func(b byte) { replayKeys = append(replayKeys, b) }
	for {
		line, ok := lr.Next()
		if !ok {
			break
		}
		for line != "" {
			b := line[0]
			line = line[1:]
			if b != '\\' || line == "" {
				// unescaped char, or \ at end of line
				q(b)
				continue
			}
			b, line = line[0], line[1:]
			switch b {
			case 'n':
				q('\n')
			default:
				q(b)
			}
		}
	}
	debugf("replay sequence: %#v", replayKeys)
}

type lineReader struct {
	br *bufio.Reader
}

func newLineReader(r io.Reader) *lineReader {
	return &lineReader{br: bufio.NewReader(r)}
}

func (lr *lineReader) Next() (line string, ok bool) {
	for {
		line, err := lr.br.ReadString('\n')
		if err == io.EOF && line != "" {
			err = nil
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading replay file: %v", err)
			}
			return "", false
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasSuffix(line, "\n") {
			line = line[:len(line)-1]
		}
		return line, true
	}
	panic("unreachable")
}
