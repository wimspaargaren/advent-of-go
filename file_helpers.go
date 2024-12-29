package aoc

import (
	"os"
	"strings"
)

func MustReadFile(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	splitted := strings.Split(string(b), "\n")
	if splitted[len(splitted)-1] == "" {
		splitted = splitted[:len(splitted)-1]
	}
	return strings.Join(splitted, "\n")
}
