package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"shell"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	for {
		line, err := input.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		cmd, err := shell.CmdFromString(line)
		if err != nil {
			log.Fatal(err)
		}
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%s", out)
	}
}
