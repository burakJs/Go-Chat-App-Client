package utils

import (
	"bufio"
	"fmt"
	"strings"
)

// ReadMessageFromConsole returns the text read from the console
// returns the empty string with error if there is an error
func readDataFromConsole(r *bufio.Reader) (string, error) {
	text, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	text = strings.Replace(text, "\n", "", -1)
	return text, nil
}

func ReadNameFromConsole(r *bufio.Reader) (string, error) {
	fmt.Print("Your name: ")
	return readDataFromConsole(r)
}

func ReadMessageFromConsole(r *bufio.Reader) (string, error) {
	fmt.Print("Your message: ")
	return readDataFromConsole(r)
}
