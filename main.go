// Package main provides a Go Git Dashboard
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	var (
		err    error
		cmdOut []byte
	)

	cmdName := "git"
	cmdArgs := []string{"rev-parse", "--verify", "HEAD"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was a problem running your command")
		os.Exit(1)
	}
	sha := string(cmdOut)
	firstSix := sha[:6]
	fmt.Println("The first six chas of the SHA at HEAD in this repo are", firstSix)
}
