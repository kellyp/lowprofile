package lowprofile

import (
	"fmt"
	"github.com/kellyp/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/kellyp/lowprofile/Godeps/_workspace/src/github.com/vaughan0/go-ini"
	"github.com/kellyp/lowprofile/Godeps/_workspace/src/gopkg.in/mattes/go-expand-tilde.v1"
	"log"
	"os"
)

const AWS_DEFAULT_PROFILE = "AWS_DEFAULT_PROFILE"
const dot_aws_credentials = "~/.aws/credentials"

func DescribeProfiles(c *cli.Context) {
	Debugln("reading profiles from ~/.aws/credentials")
	filename, err := tilde.Expand(dot_aws_credentials)
	if err != nil {
		log.Fatal(err)
	}

	profiles := getProfiles(filename)
	for _, profile := range profiles {
		println(profile)
	}
}

func DescribeActiveProfile(c *cli.Context) {
	Debugln("reading variable AWS_DEFAULT_PROFILE")
	profile := os.Getenv(AWS_DEFAULT_PROFILE)
	if len(profile) > 0 {
		fmt.Printf("current profile is %s\n", profile)
	} else {
		fmt.Println("there is currently no active profile")
	}
}

func getProfiles(filename string) []string {
	config, err := ini.LoadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var profiles []string
	for profile := range config {
		profiles = append(profiles, profile)
	}
	return profiles
}
