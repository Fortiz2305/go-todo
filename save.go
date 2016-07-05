package main

import (
  "fmt"
  "github.com/gonuts/commander"
  "encoding/json"
  "bytes"
)

type Task struct {
  Id int64
  Todo string
  Date string
  Status string
}

func todo_save(filename string) *commander.Command {
  save := func(cmd *commander.Command, args []string) error {
    if len(args) == 0 {
      cmd.Usage()
      return nil
    }
    task := Task{1, "Programming in Go", "hoy", "OPEN"}
    task_json, _ := json.Marshal(task)
    task_json, _ = prettyprint(task_json)
    fmt.Printf("%s", task_json)
    return nil
  }

  return &commander.Command {
    Run: save,
    UsageLine: "save [message]",
    Short: "save go_todo",
  }
}

func prettyprint(b []byte) ([]byte, error) {
  var output bytes.Buffer
  err := json.Indent(&output, b, "", "  ")
  return output.Bytes(), err
}
