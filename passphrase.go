package cryptdo

import (
	"errors"
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var ErrEmptyPassphase = errors.New("passphrase must not be empty")

func ReadPassphrase(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)
	defer fmt.Println()

	input, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	if len(input) == 0 {
		return "", ErrEmptyPassphase
	}

	return string(input), nil
}