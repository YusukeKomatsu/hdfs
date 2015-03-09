package main

import (
  "os"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "hdfs"
  app.Usage = "testing yeah!!"
  app.Action = func(c *cli.Context) {
    println("Hello, human!!\n");
  }

  app.Run(os.Args)
}