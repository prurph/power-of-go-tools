package main

import (
	"os"
	"shell"
)

func main() {
	session := shell.NewSession(os.Stdout, os.Stdin, os.Stderr)
	session.Run()
}
