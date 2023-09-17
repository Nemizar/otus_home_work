package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir %s: %w", dir, err)
	}

	if err = os.Chdir(dir); err != nil {
		return nil, fmt.Errorf("change dir %s: %w", dir, err)
	}

	res := make(Environment)

	for _, f := range dirEntries {
		str, needRemove, err := readFile(f)
		if err != nil {
			return nil, fmt.Errorf("read file %w", err)
		}

		env := EnvValue{
			Value:      str,
			NeedRemove: needRemove,
		}

		res[f.Name()] = env
	}

	return res, nil
}

func readFile(dirEntry os.DirEntry) (string, bool, error) {
	if strings.Contains(dirEntry.Name(), "=") {
		return "", true, fmt.Errorf("unsuported file name %s", dirEntry.Name())
	}

	envFile, err := os.Open(dirEntry.Name())
	if err != nil {
		log.Fatalf("open file %s: %v", dirEntry.Name(), err)
	}

	defer func(envFile *os.File) {
		err := envFile.Close()
		if err != nil {
			log.Fatalf("close file %s: %v", envFile.Name(), err)
		}
	}(envFile)

	fi, err := envFile.Stat()
	if err != nil {
		return "", true, fmt.Errorf("get stat file %s: %w", envFile.Name(), err)
	}

	if fi.Size() == 0 {
		return "", true, nil
	}

	reader := bufio.NewReader(envFile)
	line, _, err := reader.ReadLine()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return string(line), false, nil
		}

		return "", true, fmt.Errorf("read file %s: %w", envFile.Name(), err)
	}

	for i, b := range line {
		if b == 0x00 {
			line[i] = '\n'
		}
	}

	res := strings.TrimRight(string(line), " \t")

	return res, false, nil
}
