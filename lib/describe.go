package awspm

import (
  "os"
  "fmt"
  "github.com/codegangsta/cli"
)

const AWS_DEFAULT_PROFILE = "AWS_DEFAULT_PROFILE"

func DescribeActiveProfile(c *cli.Context) {
  Debugln("reading variable AWS_DEFAULT_PROFILE")
  profile := os.Getenv(AWS_DEFAULT_PROFILE)
  if len(profile) > 0 {
    fmt.Printf("current profile is %s\n", profile)
  } else {
    fmt.Println("there is currently no active profile")
  }
}
