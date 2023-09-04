package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/RonkZeDonk/uogcal/pkg/collector"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "uogcalAdmin",
		Usage: "Control the UoGCal site from a CLI",
		Commands: []*cli.Command{
			{
				Name:      "refresh",
				ArgsUsage: "<update up to date (formatted yyyy-mm-dd)>",
				Aliases:   []string{"rf"},
				Usage:     "refresh all course sections before a certain date. used to add exams after they are announced",
				Action: func(ctx *cli.Context) error {
					date := ctx.Args().Get(0)
					term := ctx.Args().Get(1)
					if len(term) != 3 {
						return fmt.Errorf("a term should be 3 chars (example: F23)")
					}
					if term[0] != 'F' && term[0] != 'W' && term[0] != 'S' {
						return fmt.Errorf("first letter of term must be one of F, W, or S")
					}
					numeric, err := regexp.Compile(`^[0-9]{2}$`)
					if err != nil {
						panic(err)
					}
					if !numeric.MatchString(term[1:]) {
						return fmt.Errorf("last two chars of term must be numeric")
					}

					tDate, err := time.Parse(time.DateOnly, date)
					if err != nil {
						return fmt.Errorf("couldn't parse date (expected yyyy-mm-dd formatted date)")
					}

					if godotenv.Load() != nil {
						log.Fatalln("Error loading in the .env file")
					}
					err = collector.AddNewSections(tDate, term)
					if err != nil {
						log.Fatalln(err)
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
