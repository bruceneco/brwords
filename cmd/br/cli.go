package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/bruceneco/brwords"
	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
	"os"
)

func main() {
	cmd := &cli.Command{
		Name:        "brwords",
		Usage:       "PT-BR words CLI",
		Aliases:     []string{"bw"},
		Description: "Shows up the definitions of a word",
		Flags: []cli.Flag{
			jsonOutFlag,
			etymologyFlag,
			synonymFlag,
			examplesFlag,
			meaningsFlag,
		},
		Action:                exec,
		ArgsUsage:             "[word]",
		EnableShellCompletion: true,
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err)
	}
}

func exec(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() > 0 {
		return defineWord(cmd)
	}
	return cli.ShowAppHelp(cmd)
}
func defineWord(cmd *cli.Command) error {
	loading := &loader{cmd: cmd}
	loading.show()
	s := brwords.NewScrap()
	word := cmd.Args().First()
	if word == "" {
		return errors.New("no word specified")
	}
	definition, err := s.Word(word)
	loading.hide()
	if err != nil {
		if errors.Is(err, brwords.ErrCantVisit) {
			fmt.Println(color.RedString("⚠️ Palavra não encontrada"))
			return nil
		}
		return err
	}
	switch {
	case cmd.Bool("json"):
		return loadJsonEncoder().Encode(definition)
	default:
		fmt.Print(formatTxtOutput(definition, cmd))
		return nil
	}
}
