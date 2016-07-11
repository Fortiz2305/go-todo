package main

import (
  "fmt"
  "github.com/gonuts/commander"
  "encoding/json"
  "bytes"
  "os"
  "time"
  "math/rand"
)

type Task struct {
  Id int64
  Todo string
  Date time.Time
  Status string
}

func todo_save(filename string) *commander.Command {
  save := func(cmd *commander.Command, args []string) error {
    if len(args) == 0 {
      cmd.Usage()
      return nil
    }

    task := Task{rand.Int(), os.Args[2], time.Now(), "OPEN"}
    task_json, _ := json.Marshal(task)
    task_json, _ = prettyprint(task_json)
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
      panic(err)
    }

    defer file.Close()
    _, err = fmt.Fprintf(file, "%s\n", task_json)
    fmt.Printf("%s", task_json)
    if err != nil {
      panic(err)
    }
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
