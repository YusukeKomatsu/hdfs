package main

import (
  "os"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name   = "hdfs"
  app.Usage  = "testing yeah!!"
  app.Author = "Yusuke Komatsu"
  app.Email  = "tienlen042@gmail.com"
  app.Action = func(c *cli.Context) {
    println("Hello, human!!\n");
  }

  app.Run(os.Args)
}
