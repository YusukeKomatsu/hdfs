package main

import (
  "bytes"
  "fmt"
  "os"
  "os/exec"
  "github.com/codegangsta/cli"
//  "github.com/davecgh/go-spew/spew"
)

func main() {
  app := cli.NewApp()
  app.Name    = "hdfs"
  app.Usage   = "testing yeah!!"
  app.Author  = "Yusuke Komatsu"
  app.Email   = "tienlen042@gmail.com"
  app.Version = VERSION

  app.Commands = []cli.Command{
    {
      Name:  "ls",
      Usage: "List the contents that match the specified file pattern.",
      Action: func(c *cli.Context) {
        cmdExec("-ls", c.Args().First())
      },
    },
    {
      Name:  "lsr",
      Usage: "Recursively list the contents that match the specified file pattern.",
      Action: func(c *cli.Context) {
        cmdExec("-lsr", c.Args().First())
      },
    },
    {
      Name:  "du",
      Usage: "Show the amount of space, in bytes, used by the files that match the specified file pattern.",
      Action: func(c *cli.Context) {
        println("Sorry, du command is still under construction")
      },
    },
    {
      Name:  "dus",
      Usage: "Show the amount of space, in bytes, used by the files that match the specified file pattern.",
      Action: func(c *cli.Context) {
        println("Sorry, dus command is still under construction")
      },
    },
    {
      Name:  "mv",
      Usage: "Move files that match the specified file pattern <src>  to a destination <dst>.",
      Action: func(c *cli.Context) {
        println("Sorry, mv command is still under construction")
      },
    },
    {
      Name:  "cp",
      Usage: "Copy files that match the file pattern <src> to a destination.",
      Action: func(c *cli.Context) {
        println("Sorry, cp command is still under construction")
      },
    },
    {
      Name:  "rm",
      Usage: "Delete all files that match the specified file pattern.",
      Action: func(c *cli.Context) {
        path := c.Args().First()
        fmt.Printf("WARNING: Are you sure? (yes/no)\ntarget:%s\n", path)
        if askForConfirmation() {
          cmdExec("-rm", c.Args().First())
        }
      },
    },
    {
      Name:  "rmr",
      Usage: "Remove all directories which match the specified file pattern.",
      Action: func(c *cli.Context) {
        path := c.Args().First()
        fmt.Printf("WARNING: Are you sure? (yes/no)\ntarget:%s\n", path)
        if askForConfirmation() {
          cmdExec("-rmr", c.Args().First())
        }
      },
    },
    {
      Name:  "put",
      Usage: "Copy files from the local file system into fs.",
      Action: func(c *cli.Context) {
        println("Sorry, put command is still under construction")
      },
    },
    {
      Name:  "copyFromLocal",
      Usage: "Identical to the -put command.",
      Action: func(c *cli.Context) {
        println("Sorry, copyFromLocal command is still under construction")
      },
    },
    {
      Name:  "copyToLocal",
      Usage: "Identical to the get command.",
      Action: func(c *cli.Context) {
        println("Sorry, copyToLocal command is still under construction")
      },
    },
    {
      Name:  "moveFromLocal",
      Usage: "Same as -put, except that the source is deleted after it's copied.",
      Action: func(c *cli.Context) {
        println("Sorry, moveFromLocal command is still under construction")
      },
    },
    {
      Name:  "moveToLocal",
      Usage: "Not implemented yet",
      Action: func(c *cli.Context) {
        println("Sorry, moveToLocal command is still under construction")
      },
    },
    {
      Name:  "get",
      Usage: "Copy files that match the file pattern <src> to the local name.",
      Action: func(c *cli.Context) {
        println("Sorry, get command is still under construction")
      },
    },
    {
      Name:  "getmerge",
      Usage: "Get all the files in the directories that match the source file pattern and merge and sort them to only one file on local fs.",
      Action: func(c *cli.Context) {
        println("Sorry, getmerge command is still under construction")
      },
    },
    {
      Name:  "cat",
      Usage: "Fetch all files that match the file pattern sorce and display their content on stdout.",
      Action: func(c *cli.Context) {
        println("Sorry, cat command is still under construction")
      },
    },
    {
      Name:  "mkdir",
      Usage: "Create a directory in specified location.",
      Action: func(c *cli.Context) {
        println("Sorry, mkdir command is still under construction")
      },
    },
    {
      Name:  "setrep",
      Usage: "et the replication level of a file.",
      Action: func(c *cli.Context) {
        println("Sorry, setrep command is still under construction")
      },
    },
    {
      Name:  "tail",
      Usage: "Show the last 1KB of the file.",
      Action: func(c *cli.Context) {
        println("Sorry, tail command is still under construction")
      },
    },
    {
      Name:  "touchz",
      Usage: "Write a timestamp in yyyy-MM-dd HH:mm:ss format in a file at path.",
      Action: func(c *cli.Context) {
        println("Sorry, touchz command is still under construction")
      },
    },
    {
      Name:  "test",
      Usage: "If file { exists, has zero length, is a directory then return 0, else return 1.",
      Action: func(c *cli.Context) {
        println("Sorry, test command is still under construction")
      },
    },
    {
      Name:  "text",
      Usage: "Takes a source file and outputs the file in text format.",
      Action: func(c *cli.Context) {
        println("Sorry, text command is still under construction")
      },
    },
    {
      Name:  "stat",
      Usage: "Print statistics about the file/directory at path",
      Action: func(c *cli.Context) {
        println("stat command is still under construction")
      },
    },
    {
      Name:  "chmod",
      Usage: "Changes permissions of a file. This works similar to shell's chmod with a few exceptions.",
      Action: func(c *cli.Context) {
        println("chmod command is still under construction")
      },
    },
    {
      Name:  "chown",
      Usage: "Changes owner and group of a file. This is similar to shell's chown with a few exceptions.",
      Action: func(c *cli.Context) {
        println("chown command is still under construction")
      },
    },
    {
      Name:  "chgrp",
      Usage: "This is equivalent to chown ... :GROUP ...",
      Action: func(c *cli.Context) {
        println("chgrp command is still under construction")
      },
    },
    {
      Name:  "count",
      Usage: "Count the number of directories, files and bytes under the paths that match the specified file pattern.",
      Action: func(c *cli.Context) {
        println("count command is still under construction")
      },
    },
  }

  app.Run(os.Args)
}

func cmdExec(command string, argment string) {
  if len(argment) <= 0 {
    fmt.Println("Can't get argment.");
    os.Exit(1)
  }

  rootEnv := os.Getenv("HADOOP_BIN_PATH")

  if len(rootEnv) <= 0 {
    rootEnv = "/usr/local/hadoop/bin/hadoop";
  }
  cmd := exec.Command(rootEnv, "fs" , command, argment)
  var stdout bytes.Buffer
  cmd.Stdout = &stdout

  err := cmd.Run()

  if err != nil {
    println(err)
    os.Exit(1)
  }

  println(stdout.String())
}
