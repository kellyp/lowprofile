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

const zsh  = "zsh"
const zshrc = "~/.zshrc"
const bash = "bash"
const profileVariable = "AWS_DEFAULT_PROFILE"

func ActivateProfile(c *cli.Context) {
  Debugln("checking shell")
  shell := os.Getenv("SHELL")
  Debugf("the shell is %s", shell)
  profile := c.String("profile")

  if strings.Contains(shell, zsh) {
    Debugln("checking for variable in ~/.zshrc")
    filename, err := tilde.Expand(zshrc)
    if err != nil {
        log.Fatal(err)
    }
    found, lines := scanFileForVariable(filename, profileVariable, profile)
    if found {
      writeFile(filename, lines)
    }

  } else if strings.Contains(shell, bash) {

  } else {
    fmt.Printf("Sorry, %s is not supported", shell)
  }
}

func scanFileForVariable(filename string, variable string, profile string) (bool, []string) {

  file, err := os.Open(filename)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  var lines []string
  found := false
  Debugln(fmt.Sprintf("(export\\s+%s=)\\w*", variable))
  regex := regexp.MustCompile(fmt.Sprintf("(export\\s+%s=)\\w*", variable))
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
      log.Fatal(err)
  }


  return found, lines
}

func writeFile(filename string, lines []string) {
  file, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  Debugf("Writing to file %s", filename)

  w := bufio.NewWriter(file)
  for index := range lines {
    Debugf("Writing: %s", lines[index])
    fmt.Fprintln(w, lines[index])
  }

	w.Flush() // Don't forget to flush!
}
