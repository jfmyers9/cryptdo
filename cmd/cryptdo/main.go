package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	flags "github.com/jessevdk/go-flags"

	"github.com/xoebus/cryptdo"
)

var opts struct {
	Passphrase string `short:"p" long:"passphrase" description:"passphrase for file encryption" required:"true"`

	Positional struct {
		Command string   `positional-arg-name:"COMMAND" required:"true"`
		Args    []string `positional-arg-name:"ARG"`
	} `positional-args:"yes"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	encryptedFiles, _ := filepath.Glob("*.enc")
	for _, file := range encryptedFiles {
		ciphertext, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalln(err)
		}

		plaintext, err := cryptdo.Decrypt(ciphertext, opts.Passphrase)
		if err != nil {
			log.Fatalln(err)
		}

		err = ioutil.WriteFile(decryptedName(file), plaintext, 0400)
		if err != nil {
			log.Fatalln(err)
		}
	}

	cmd := exec.Command(opts.Positional.Command, opts.Positional.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	for _, file := range encryptedFiles {
		plaintext, err := ioutil.ReadFile(decryptedName(file))
		if err != nil {
			log.Fatalln(err)
		}

		ciphertext, err := cryptdo.Encrypt(plaintext, opts.Passphrase)
		if err != nil {
			log.Fatalln(err)
		}

		newPath := file + ".new"
		err = ioutil.WriteFile(newPath, ciphertext, 0400)
		if err != nil {
			log.Fatalln(err)
		}

		if err = os.Rename(newPath, file); err != nil {
			log.Fatalln(err)
		}

		if err = os.Remove(decryptedName(file)); err != nil {
			log.Fatalln(err)
		}
	}
}

func decryptedName(file string) string {
	return file[:strings.LastIndex(file, ".enc")]
}
