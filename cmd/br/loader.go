package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

type loader struct {
	cmd *cli.Command
}

func (l loader) show() {
	if !l.cmd.Bool("json") {
		fmt.Println(color.YellowString("🔍 Pesquisando..."))
	}
}
func (l loader) hide() {
	if !l.cmd.Bool("json") {
		fmt.Print("\033[1A\033[K")
	}
}
