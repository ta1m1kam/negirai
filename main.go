package main

import (
	"github.com/TaigaMikami/negirai/git"
	"github.com/TaigaMikami/negirai/sound"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

//func main() {
//	if err := sound.Play(); err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
	app := &cli.App{
		Name: "negirai",
		Authors: []*cli.Author{
			&cli.Author{
				Name: "Taiga Mikami",
			},
		},
		Usage: "This command says thank you for your hard work.",
		Commands: []*cli.Command{
			{
				Name:    "appreciate",
				Aliases: []string{"a"},
				Usage:   "Thank you for your hard work by playing the audio",
				Action: func(c *cli.Context) error {
					err := sound.Play()
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "pre-commit",
				Aliases: []string{"p"},
				Usage:   "Add pre-commit script for your git project",
				Action: func(c *cli.Context) error {
					err := git.MakePreCommitFile()
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
