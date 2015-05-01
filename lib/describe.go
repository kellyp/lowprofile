package awspm

import (
  "os/user"
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/vaughan0/go-ini"
  "path/filepath"
  "strings"
)

func describe(c *cli.Context) {
  Debugln("reading file ~/.aws/config")
  usr, _ := user.Current()
  path, pathError := filepath.Abs(fmt.Sprintf("%v/.aws/config", usr.HomeDir))
  Debugf("Error %v", pathError)
  config, iniError := ini.LoadFile(path)
  Debugf("Error %v", iniError)

  for name, section := range config {
    fmt.Printf("Section name: %s\n", name)
    if len(c.String("name")) == 0 {
      Debugf("Found section: %v", name)
      for key, value := range section {
        Debugf("%v: %v\n", key, value)
      }
    } else if strings.Contains(name, c.String("name")) {
      Debugf("Found section: %v", name)
      for key, value := range section {
        Debugf("%v: %v\n", key, value)
      }
    }
  }
}
