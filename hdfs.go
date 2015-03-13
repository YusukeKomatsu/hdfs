package main

import (
  "bytes"
  "fmt"
  "os"
  "os/exec"
  "github.com/codegangsta/cli"
//  "github.com/davecgh/go-spew/spew"
  "github.com/fatih/color"
)

func main() {
  app := cli.NewApp()
  app.Name    = "hdfs"
  app.Usage   = "Hadoop-fs wrapper command"
  app.Author  = "Yusuke Komatsu"
  app.Email   = "tienlen042@gmail.com"
  app.Version = VERSION

  app.Commands = []cli.Command{
    {
      Name:  "ls",
      Usage: "List the contents that match the specified file pattern.",
      Action: cmdList,
    },
    {
      Name:  "lsr",
      Usage: "Recursively list the contents that match the specified file pattern.",
      Action: cmdListRecursive,
    },
    {
      Name:  "du",
      Usage: "Show the amount of space, in bytes, used by the files that match the specified file pattern.",
      Action: cmdDirectoryUsage,
    },
    {
      Name:  "dus",
      Usage: "Show the amount of space, in bytes, used by the files that match the specified file pattern.",
      Action: cmdDirectoryUsageSummary,
    },
    {
      Name:  "mv",
      Usage: "Move files that match the specified file pattern <src>  to a destination <dst>.",
      Action: cmdMove,
    },
    {
      Name:  "cp",
      Usage: "Copy files that match the file pattern <src> to a destination.",
      Action: cmdCopy,
    },
    {
      Name:  "rm",
      Usage: "Delete all files that match the specified file pattern.",
      Action: cmdRemove,
    },
    {
      Name:  "rmr",
      Usage: "Remove all directories which match the specified file pattern.",
      Action: cmdRemoveRecursive,
    },
    {
      Name:  "put",
      Usage: "Copy files from the local file system into fs.",
      Action: cmdPut,
    },
    {
      Name:  "copyFromLocal",
      Usage: "Identical to the -put command.",
      Action: cmdCopyFromLocal,
    },
    {
      Name:  "copyToLocal",
      Usage: "Identical to the get command.",
      Action: cmdCopyToLocal,
    },
    {
      Name:  "moveFromLocal",
      Usage: "Same as -put, except that the source is deleted after it's copied.",
      Action: cmdMoveFromLocal,
    },
    {
      Name:  "moveToLocal",
      Usage: "Not implemented yet",
      Action: cmdMoveToLocal,
    },
    {
      Name:  "get",
      Usage: "Copy files that match the file pattern <src> to the local name.",
      Action: cmdGetFiles,
    },
    {
      Name:  "getmerge",
      Usage: "Get all the files in the directories that match the source file pattern and merge and sort them to only one file on local fs.",
      Action: cmdGetMerge,
    },
    {
      Name:  "cat",
      Usage: "Fetch all files that match the file pattern sorce and display their content on stdout.",
      Action: cmdCat,
    },
    {
      Name:  "mkdir",
      Usage: "Create a directory in specified location.",
      Action: cmdMakeDirectory,
    },
    {
      Name:  "setrep",
      Usage: "Set the replication level of a file.",
      Action: cmdSetReplicationLevel,
    },
    {
      Name:  "tail",
      Usage: "Show the last 1KB of the file.",
      Action: cmdTail,
    },
    {
      Name:  "touchz",
      Usage: "Write a timestamp in yyyy-MM-dd HH:mm:ss format in a file at path.",
      Action: cmdTouch,
    },
    {
      Name:  "test",
      Usage: "If file { exists, has zero length, is a directory then return 0, else return 1.",
      Action: cmdTest,
    },
    {
      Name:  "text",
      Usage: "Takes a source file and outputs the file in text format.",
      Action: cmdTextOutput,
    },
    {
      Name:  "stat",
      Usage: "Print statistics about the file/directory at path",
      Action: cmdPrintStatistics,
    },
    {
      Name:  "chmod",
      Usage: "Changes permissions of a file. This works similar to shell's chmod with a few exceptions.",
      Action: cmdChangeMode,
    },
    {
      Name:  "chown",
      Usage: "Changes owner and group of a file. This is similar to shell's chown with a few exceptions.",
      Action: cmdChangeOwn,
    },
    {
      Name:  "chgrp",
      Usage: "This is equivalent to chown ... :GROUP ...",
      Action: cmdChangeGroup,
    },
    {
      Name:  "count",
      Usage: "Count the number of directories, files and bytes under the paths that match the specified file pattern.",
      Action: cmdCount,
    },
  }

  app.Run(os.Args)
}

func cmdExec(command string, args ...string) {
  if len(args) <= 0 {
    fmt.Println("Can't get argments.");
    os.Exit(1)
  }

  rootEnv := os.Getenv("HADOOP_BIN_PATH")

  if len(rootEnv) <= 0 {
    rootEnv = "/usr/local/hadoop/bin/hadoop";
  }

  argments := []string{"fs", command}
  argments = append(argments, args...)

  cmd := exec.Command(rootEnv, argments...)
  var stdout bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr

  err := cmd.Run()

  if err != nil {
    println(fmt.Sprint(err) + ": " + stderr.String())
    os.Exit(1)
  }

  println(stdout.String())
}

func cmdCat(c *cli.Context) {
  cmdExec("-cat", c.Args()...)
}

func cmdChangeGroup(c *cli.Context) {
  cmdExec("-chgrp", c.Args()...)
}

func cmdChangeMode(c *cli.Context) {
  cmdExec("-chmod", c.Args()...)
}

func cmdChangeOwn(c *cli.Context) {
  cmdExec("-chown", c.Args()...)
}

func cmdCopyFromLocal(c *cli.Context) {
  cmdExec("-copyFromLocal", c.Args()...)
}

func cmdCopyToLocal(c *cli.Context) {
  cmdExec("-copyToLocal", c.Args()...)
}

func cmdCount(c *cli.Context) {
  cmdExec("-count", c.Args()...)
}

func cmdCopy(c *cli.Context) {
  cmdExec("-cp", c.Args()...)
}

func cmdDirectoryUsage(c *cli.Context) {
  cmdExec("-du", c.Args()...)
}

func cmdDirectoryUsageSummary(c *cli.Context) {
  cmdExec("-dus", c.Args()...)
}

func cmdGetFiles(c *cli.Context) {
  cmdExec("-get", c.Args()...)
}

func cmdGetMerge(c *cli.Context) {
  cmdExec("-getmerge", c.Args()...)
}

func cmdList(c *cli.Context) {
  cmdExec("-ls", c.Args()...)
}

func cmdListRecursive(c *cli.Context) {
  cmdExec("-lsr", c.Args()...)
}

func cmdMakeDirectory(c *cli.Context) {
  cmdExec("-mkdir", c.Args()...)
}

func cmdMove(c *cli.Context) {
  cmdExec("-mv", c.Args()...)
}

func cmdMoveFromLocal(c *cli.Context) {
  cmdExec("-moveFromLocal", c.Args()...)
}

func cmdMoveToLocal(c *cli.Context) {
  cmdExec("-moveToLocal", c.Args()...)
}

func cmdPrintStatistics(c *cli.Context) {
  cmdExec("-stat", c.Args()...)
}

func cmdPut(c *cli.Context) {
  cmdExec("-put", c.Args()...)
}

func cmdRemove(c *cli.Context) {
  path := c.Args().First()
  color.Red("WARNING: Are you sure? (yes/no)\ntarget:%s\n", path)
  if askForConfirmation() {
    cmdExec("-rm", c.Args().First())
  }
}

func cmdRemoveRecursive(c *cli.Context) {
  path := c.Args().First()
  color.Red("WARNING: Are you sure? (yes/no)\ntarget:%s\n", path)
  if askForConfirmation() {
    cmdExec("-rmr", c.Args().First())
  }
}

func cmdSetReplicationLevel(c *cli.Context) {
  cmdExec("-setrep", c.Args()...)
}

func cmdTail(c *cli.Context) {
  cmdExec("-tail", c.Args()...)
}

func cmdTest(c *cli.Context) {
  cmdExec("-test", c.Args()...)
}

func cmdTextOutput(c *cli.Context) {
  cmdExec("-text", c.Args()...)
}

func cmdTouch(c *cli.Context) {
  cmdExec("-touchz", c.Args()...)
}