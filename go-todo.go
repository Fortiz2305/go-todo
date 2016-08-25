package main

import (
  "go-todo/todo"
)

func main() {
  err := todo.Run()
  if err !=  nil {
    panic(err)
  }
}
