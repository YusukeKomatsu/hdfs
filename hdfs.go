package hdfs

import (
  "os"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name    = "hdfs"
  app.Usage   = "testing yeah!!"
  app.Author  = "Yusuke Komatsu"
  app.Email   = "tienlen042@gmail.com"
  app.Version = VERSION

  app.comands = []cli.Command{
    {
      Name:  "ls",
      Usage: "List the contents that match the specified file pattern.",
      Action: func(c *cli.Context) {
        println("ls command")
      }
    },
    {
      Name:  "lsr",
      Usage: "Recursively list the contents that match the specified file pattern.",
      Action: func(c *cli.Context) {
        println("lsr command")
      }
    },
    {
      Name:  "du",
      Usage: "Show the amount of space, in bytes, used by the files that match the specified file pattern.",
      Action: func(c *cli.Context) {
        println("du command")
      }
    },
    {
      Name:  "dus",
      Usage: "Show the amount of space, in bytes, used by the files that match the specified file pattern.  Equivalent to the unix command 'du -sb'  The output is in the form ",
      Action: func(c *cli.Context) {
        println("dus command")
      }
    },
    {
      Name:  "mv",
      Usage: "Move files that match the specified file pattern <src>  to a destination <dst>.  When moving multiple files, the destination must be a directory.",
      Action: func(c *cli.Context) {
        println("mv command")
      }
    },
    {
      Name:  "cp",
      Usage: "Copy files that match the file pattern <src> to a destination.\nWhen copying multiple files, the destination must be a directory.",
      Action: func(c *cli.Context) {
        println("cp command")
      }
    },
    {
      Name:  "rm",
      Usage: "Delete all files that match the specified file pattern.",
      Action: func(c *cli.Context) {
        println("rm command")
      }
    },
    {
      Name:  "rmr",
      Usage: "Remove all directories which match the specified file pattern.",
      Action: func(c *cli.Context) {
        println("rmr command")
      }
    },
    {
      Name:  "put",
      Usage: "Copy files from the local file system into fs.",
      Action: func(c *cli.Context) {
        println("put command")
      }
    },
    {
      Name:  "copyFromLocal",
      Usage: "Identical to the -put command.",
      Action: func(c *cli.Context) {
        println("copyFromLocal command")
      }
    },
    {
      Name:  "copyToLocal",
      Usage: "Identical to the get command.",
      Action: func(c *cli.Context) {
        println("copyToLocal command")
      }
    },
    {
      Name:  "moveFromLocal",
      Usage: "Same as -put, except that the source is deleted after it's copied.",
      Action: func(c *cli.Context) {
        println("moveFromLocal command")
      }
    },
    {
      Name:  "moveToLocal",
      Usage: "Not implemented yet",
      Action: func(c *cli.Context) {
        println("moveToLocal command")
      }
    },
    {
      Name:  "get",
      Usage: "Copy files that match the file pattern <src> to the local name.",
      Action: func(c *cli.Context) {
        println("get command")
      }
    },
    {
      Name:  "getmerge",
      Usage: "Get all the files in the directories that match the source file pattern and merge and sort them to only one file on local fs. <src> is kept.",
      Action: func(c *cli.Context) {
        println("getmerge command")
      }
    },
    {
      Name:  "cat",
      Usage: "Fetch all files that match the file pattern sorce and display their content on stdout.",
      Action: func(c *cli.Context) {
        println("cat command")
      }
    },
    {
      Name:  "mkdir",
      Usage: "Create a directory in specified location.",
      Action: func(c *cli.Context) {
        println("mkdir command")
      }
    },
    {
      Name:  "setrep",
      Usage: "et the replication level of a file.",
      Action: func(c *cli.Context) {
        println("setrep command")
      }
    },
    {
      Name:  "tail",
      Usage: "Show the last 1KB of the file.",
      Action: func(c *cli.Context) {
        println("tail command")
      }
    },
    {
      Name:  "touchz",
      Usage: "Write a timestamp in yyyy-MM-dd HH:mm:ss format in a file at path. An error is returned if the file exists with non-zero length",
      Action: func(c *cli.Context) {
        println("touchz command")
      }
    },
    {
      Name:  "test",
      Usage: "If file { exists, has zero length, is a directory then return 0, else return 1.",
      Action: func(c *cli.Context) {
        println("test command")
      }
    },
    {
      Name:  "text",
      Usage: "Takes a source file and outputs the file in text format.\nThe allowed formats are zip and TextRecordInputStream.",
      Action: func(c *cli.Context) {
        println("text command")
      }
    },
    {
      Name:  "stat",
      Usage: "Print statistics about the file/directory at path",
      Action: func(c *cli.Context) {
        println("stat command")
      }
    },
    {
      Name:  "chmod",
      Usage: "Changes permissions of a file.\nThis works similar to shell's chmod with a few exceptions.",
      Action: func(c *cli.Context) {
        println("chmod command")
      }
    },
    {
      Name:  "chown",
      Usage: "Changes owner and group of a file.\nThis is similar to shell's chown with a few exceptions.",
      Action: func(c *cli.Context) {
        println("chown command")
      }
    },},
    {
      Name:  "chgrp",
      Usage: "This is equivalent to chown ... :GROUP ...",
      Action: func(c *cli.Context) {
        println("chgrp command")
      }
    },},
    {
      Name:  "count",
      Usage: "Count the number of directories, files and bytes under the paths that match the specified file pattern.",
      Action: func(c *cli.Context) {
        println("count command")
      }
    },
  }

  app.Run(os.Args)
}
