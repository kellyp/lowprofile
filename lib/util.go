package lowprofile

import (
  "os"
  "bufio"
  "fmt"
  "log"
  "strings"
  "errors"
)

const zsh  = "/bin/zsh"
const zshrc = "~/.zshrc"
const bash_profile = "~/.bash_profile"
const bash = "/bin/bash"
const profileVariable = "AWS_PROFILE"


var Debug bool = false

func Shells() map[string]string {
  return map[string]string {bash: bash_profile, zsh: zshrc}
}

func Debugln(str string) {
  if Debug {
    fmt.Printf("DEBUG: %v\n", str)
  }
}

func Debugf(str string, args ...interface{}) {
  if Debug {
    var format = fmt.Sprintf(str, args...)
    fmt.Printf("DEBUG: %v\n", format)
  }
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

	w.Flush()
}

func checkForShell()(string, error) {
	Debugln("checking shell")
	shell := os.Getenv("SHELL")
	Debugf("the shell is %s", shell)
	var err error
	var filename string
	if strings.Contains(shell, zsh) {
		Debugln("checking for variable in ~/.zshrc")
		filename = zshrc
	} else if strings.Contains(shell, bash) {
		Debugln("checking for variable in ~/.bash_profile")
		filename = bash_profile
	} else {
		 err = errors.New(fmt.Sprintf("Sorry, %s is not a supported shell", shell))
	}

	return filename, err
}
