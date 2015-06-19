package main

import (
  "os"
  "github.com/codegangsta/cli"
  "github.com/kellyp/awspm/awspm"
)

func main() {
  app := cli.NewApp()
  app.Name = "low-profile"
  app.Usage = "do work"
  app.Version = "0.1"

  app.Flags = []cli.Flag {
    cli.BoolFlag{
      Name: "debug",
      Usage: "shall we debug?",
    },
  }
  app.Before =  func(c *cli.Context) error {
    if c.Bool("debug") {
      awspm.Debug = true
      awspm.Debugln("Turning debug on.")
    }
    return nil
  }

  app.Commands = []cli.Command{
    {
      Name:      "describe-profiles",
      Aliases:   []string{"d"},
      Usage:     `Describes the list of AWS profile`,
      Action: awspm.DescribeProfiles,
    },
    {
      Name:      "describe-active-profile",
      Aliases:   []string{"dap"},
      Usage:     `Describes the currently active AWS profile`,
      Action: awspm.DescribeActiveProfile,
    },
    {
      Name:      "activate-profile",
      Aliases:   []string{"ap"},
      Usage:     `Sets the currently active profile`,
      Action:    awspm.ActivateProfile,
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
      Action: awspm.DeactivateProfile,
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