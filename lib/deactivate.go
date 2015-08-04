package lowprofile

import (
	"bufio"
	"fmt"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/gopkg.in/mattes/go-expand-tilde.v1"
	"os"
	"regexp"
	"errors"
)

func BeforeDeactivateProfile(c *cli.Context) error {
	shell := os.Getenv("SHELL")
	if Shells()[shell] == "" {
		output := fmt.Sprintf("Sorry, %s is not a supported shell", shell)
		fmt.Println(output)
		return errors.New(output)
	}

	var filename, _ = tilde.Expand(Shells()[shell])
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		output := fmt.Sprintf("File %s not found", Shells()[shell])
		fmt.Println(output)
		return errors.New(output)
	}

	return nil
}

func DeactivateProfile(c *cli.Context) {
	Debugln("checking shell")
	shell := os.Getenv("SHELL")
	Debugf("the shell is %s", shell)

	profile := os.Getenv(AWS_PROFILE)
	if len(profile) > 0 {
		fmt.Printf("deactivating profile %s\n", profile)
	} else {
		fmt.Println("there is currently no active profile")
	}

	var filename, err = checkForShell()
	filename, err = tilde.Expand(filename)
	if err != nil {
		panic(err)
	}
	found, lines := scanFileForVariableAndComment(filename, profileVariable)
	if found {
		writeFile(filename, lines)
	}
}

func scanFileForVariableAndComment(filename string, variable string) (bool, []string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	found := false

	regex := regexp.MustCompile(fmt.Sprintf("\\#*\\s*(export\\s+%s=\\w*)", variable))
	replace := "# ${1}"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if regex.MatchString(text) {
			found = true
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
