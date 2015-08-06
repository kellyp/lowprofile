package main

import (
	"fmt"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/lib"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "lowprofile"
	app.Usage = "Control AWS profiles"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "shall we debug?",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			lowprofile.Debug = true
			lowprofile.Debugln("Turning debug on.")
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "describe-profiles",
			Aliases: []string{"d"},
			Usage:   `Describes the list of AWS profile`,
			Before:  lowprofile.BeforeDescribeProfiles,
			Action:  lowprofile.DescribeProfiles,
		},
		{
			Name:    "describe-active-profile",
			Aliases: []string{"dap"},
			Usage:   `Describes the currently active AWS profile`,
			Action:  lowprofile.DescribeActiveProfile,
		},
		{
			Name:    "activate-profile",
			Aliases: []string{"ap"},
			Usage:   `Sets the currently active profile`,
			Before:  lowprofile.BeforeActivateProfile,
			Action:  lowprofile.ActivateProfile,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "profile",
					Usage: "name of the profile to activate",
					Value: "profile-name",
				},
			},
		},
		{
			Name:    "deactive-profile",
			Aliases: []string{"dp"},
			Usage:   `Deactivate the currently active AWS profile`,
			Before:  lowprofile.BeforeDeactivateProfile,
			Action:  lowprofile.DeactivateProfile,
		},
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	app.Run(os.Args)
}
