package lowprofile

import (
  "os"
  "bufio"
  "fmt"
  "log"
)

const zsh  = "zsh"
const zshrc = "~/.zshrc"
const bash_profile = "~/.bash_profile"
const bash = "bash"
const profileVariable = "AWS_DEFAULT_PROFILE"

var Debug bool = false

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

	w.Flush() // Don't forget to flush!
}
