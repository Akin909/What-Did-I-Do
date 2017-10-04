// Package main provides a Go Git Dashboard
package main

import (
	"fmt"
	"os"
	"os/exec"

	ui "github.com/gizak/termui"
)

func gitCommand() (out string) {

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
	return "The first six chas of the SHA at HEAD in this repo are" + firstSix
}

func quit(ui.Event) {
	ui.StopLoop()
}

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	p, g := quitMenu()
	output := gitCommand()
	git := showGitOutput(output)
	ui.Render(p, g, git)
	ui.Loop()
	// Handle q to quit
	ui.Handle("/sys/kbd/q", quit)
}
