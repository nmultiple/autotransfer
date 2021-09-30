package config

import (
	"io"
	"bufio"
	"errors"
	"os"
	"strings"
)

var (
	EmptyLine = errors.New("empty line")
	SyntaxError = errors.New("syntax error")
)

func parseLine(line string) (string, string, error) {
	line = strings.TrimSpace(line)

	if len(line) == 0 {
		return "", "", EmptyLine
	}

	eqPos := strings.IndexRune(line, '=')
	if eqPos < 0 {
		return "", "", SyntaxError
	}

	return line[:eqPos], line[eqPos+1:], nil
}

func LoadDotEnv() error {
	f, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, _, err :=  r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		key, value, err := parseLine(string(line))

		if err == EmptyLine {
			continue
		} else if err != nil {
			return err
		}

		os.Setenv(key, value)
	}

	return nil
}

