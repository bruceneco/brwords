package main

import "github.com/urfave/cli/v3"

var (
	jsonOutFlag = &cli.BoolFlag{
		Name:    "json",
		Aliases: []string{"j"},
		Usage:   "output as json",
		Value:   false,
	}
	etymologyFlag = &cli.BoolFlag{
		Name:  "et",
		Usage: "display etymologies",
		Value: false,
	}
	synonymFlag = &cli.BoolFlag{
		Name:  "sy",
		Usage: "display synonyms",
		Value: false,
	}
	examplesFlag = &cli.BoolFlag{
		Name:  "ex",
		Usage: "display examples",
		Value: false,
	}
	meaningsFlag = &cli.BoolFlag{
		HideDefault: true,
		Name:        "meaning",
		Aliases:     []string{"m", "d"},
		Usage:       "display meanings (default to true if no other display flag is specified)",
		Value:       false,
	}
)
