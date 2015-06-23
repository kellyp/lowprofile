package awspm

import (
  "os"
  "bufio"
  "fmt"
  "log"
  "github.com/codegangsta/cli"
  "gopkg.in/mattes/go-expand-tilde.v1"
  "regexp"
  "strings"
)


func DeactivateProfile(c *cli.Context) {
  Debugln("checking shell")
  shell := os.Getenv("SHELL")
  Debugf("the shell is %s", shell)
  profile := c.String("profile")

  var filename string
  if strings.Contains(shell, zsh) {
    Debugln("checking for variable in ~/.zshrc")
    filename = zshrc
  } else if strings.Contains(shell, bash) {
    Debugln("checking for variable in ~/.bash_profile")
    filename = bash_profile
  } else {
    log.Fatalf("Sorry, %s is not supported", shell)
  }

  filename, err := tilde.Expand(filename)
  if err != nil {
      log.Fatal(err)
  }
  found, lines := scanFileForVariableAndComment(filename, profileVariable, profile)
  if found {
    writeFile(filename, lines)
  }
}

func scanFileForVariableAndComment(filename string, variable string, profile string) (bool, []string) {

  file, err := os.Open(filename)
  if err != nil {
      log.Fatal(err)
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
      log.Fatal(err)
  }


  return found, lines
}
