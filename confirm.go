package main

import (
  "fmt"
  "log"
  "os"
  "sort"
)

func askForConfirmation() bool {
  var response string
  _, err := fmt.Scanln(&response)
  if err != nil {
    log.Fatal(err)
  }
  okResponses  := []string{"y", "Y", "yes", "Yes", "YES"}
  ngResponses    := []string{"n", "N", "no", "No", "NO"}
  if containsString(okResponses, response) {
    return true
  } else if containsString(ngResponses, response) {
    return false
  } else {
    fmt.Println("Please type yes or no and then press enter:")
    return askForConfirmation()
  }
}

fuc posString(slice []string, element string) int {
  for index, elem := range slice {
    if elem == element {
      return index
    }
  }
  return -1
}

func containsString(slice []string, element string) bool {
  return !(posString(slice, element) == -1)
}
