package main

import (
  "os"
  "github.com/codegangsta/cli"
  "github.com/kellyp/lowprofile/lowprofile"
)

func main() {
  app := cli.NewApp()
  app.Name = "lowprofile"
  app.Usage = "Control AWS profiles"
  app.Version = "0.1"

  app.Flags = []cli.Flag {
    cli.BoolFlag{
      Name: "debug",
      Usage: "shall we debug?",
    },
  }
  app.Before =  func(c *cli.Context) error {
    if c.Bool("debug") {
      lowprofile.Debug = true
      lowprofile.Debugln("Turning debug on.")
    }
    return nil
  }

  app.Commands = []cli.Command{
    {
      Name:      "describe-profiles",
      Aliases:   []string{"d"},
      Usage:     `Describes the list of AWS profile`,
      Action: lowprofile.DescribeProfiles,
    },
    {
      Name:      "describe-active-profile",
      Aliases:   []string{"dap"},
      Usage:     `Describes the currently active AWS profile`,
      Action: lowprofile.DescribeActiveProfile,
    },
    {
      Name:      "activate-profile",
      Aliases:   []string{"ap"},
      Usage:     `Sets the currently active profile`,
      Action:    lowprofile.ActivateProfile,
      Flags: []cli.Flag {
        cli.StringFlag{
          Name: "profile",
          Usage: "name of the profile to activate",
        },
      },
    },
    {
      Name:      "deactive-profile",
      Aliases:   []string{"dp"},
      Usage:     `Deactivate the currently active AWS profile`,
      Action: lowprofile.DeactivateProfile,
    },
    // {
    //   Name:      "template",
    //   Aliases:     []string{"r"},
    //   Usage:     "options for task templates",
    //   Flags: []cli.Flag {
    //     cli.BoolFlag{
    //       Name: "fart",
    //       Usage: "shall we fart?",
    //     },
    //   },
    //   Subcommands: []cli.Command{
    //     {
    //       Name:  "add",
    //       Usage: "add a new template",
    //       Action: func(c *cli.Context) {
    //           println("new task template: ", c.Args().First())
    //       },
    //     },
    //     {
    //       Name:  "remove",
    //       Usage: "remove an existing template",
    //       Action: func(c *cli.Context) {
    //         println("removed task template: ", c.Args().First())
    //       },
    //     },
    //   },
    // },
  }

  app.Run(os.Args)
}
