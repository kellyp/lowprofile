package lowprofile

import (
	"bufio"
	"fmt"
	"github.com/kellyp/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/kellyp/lowprofile/Godeps/_workspace/src/gopkg.in/mattes/go-expand-tilde.v1"
	"os"
	"regexp"
	"strings"
)

func ActivateProfile(c *cli.Context) {
	if !c.IsSet("profile") {
		cli.ShowSubcommandHelp(c)
		panic("")
	}

	Debugln("checking shell")
	shell := os.Getenv("SHELL")
	Debugf("the shell is %s", shell)
	profile := c.String("profile")

	fmt.Printf("activating profile %s\n", profile)

	var filename string
	if strings.Contains(shell, zsh) {
		Debugln("checking for variable in ~/.zshrc")
		filename = zshrc
	} else if strings.Contains(shell, bash) {
		Debugln("checking for variable in ~/.bash_profile")
		filename = bash_profile
	} else {
		panic(fmt.Sprintf("Sorry, %s is not a supported shell", shell))
	}

	filename, err := tilde.Expand(filename)
	if err != nil {
		panic(err)
	}
	found, lines := scanFileForVariable(filename, profileVariable, profile)
	if !found {
		lines = append(lines, fmt.Sprintf("export %s=%s", profileVariable, profile))
	}

	writeFile(filename, lines)
}

func scanFileForVariable(filename string, variable string, profile string) (bool, []string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	found := false
	regex := regexp.MustCompile(fmt.Sprintf("\\#*\\s*(export\\s+%s=).*", variable))
	replace := fmt.Sprintf("${1}%s", profile)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if regex.MatchString(text) {
			found = true
			Debugln("We got one!")
			text = regex.ReplaceAllString(text, replace)
		}
		lines = append(lines, text)
		Debugln(text)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return found, lines
}
