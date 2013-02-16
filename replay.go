package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type replayAction struct {
	// Exactly one of these will be set.
	key byte
	f   func()
}

var (
	replayActions []replayAction
	replayDelay   = 50 * time.Millisecond // delay for key actions
)

func popReplay() (b byte, ok bool) {
	for len(replayActions) > 0 {
		a := replayActions[0]
		replayActions = replayActions[1:]
		if a.key != 0 {
			time.Sleep(replayDelay)
			return a.key, true
		}
		a.f()
	}
	return 0, false
}

func loadReplay(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Failed loading replay file %s: %v", filename, err)
		return
	}
	lr := newLineReader(bytes.NewReader(data))

	q := func(key byte) { replayActions = append(replayActions, replayAction{key: key}) }
	act := func(f func()) { replayActions = append(replayActions, replayAction{f: f}) }
	for {
		line, ok := lr.Next()
		if !ok {
			break
		}

		if strings.HasPrefix(line, "!") {
			// Settings.
			for _, kv := range strings.Split(line[1:], " ") {
				kv = strings.TrimSpace(kv)
				parts := strings.SplitN(kv, "=", 2)
				switch k, v := parts[0], parts[1]; k {
				case "delay":
					d, err := time.ParseDuration(v)
					if err != nil {
						log.Printf("Bad delay %q: %v", v, err)
						continue
					}
					act(func() {
						replayDelay = d
						debugf("replay delay set to %v", d)
					})
				case "seed":
					x, err := strconv.ParseUint(v, 0, 32)
					if err != nil {
						log.Printf("Bad seed %q: %v", v, err)
						continue
					}
					act(func() {
						seedrand(uint32(x))
						debugf("seed set to %d", x)
					})
				default:
					log.Printf("Unknown replay setting %q; ignoring", k)
				}
			}
			continue
		}

		// Keys.
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
			case '`':
				q('\033') // escape
			default:
				q(b)
			}
		}
	}

	// Do any initial actions right away.
	for len(replayActions) > 0 && replayActions[0].f != nil {
		replayActions[0].f()
		replayActions = replayActions[1:]
	}
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
