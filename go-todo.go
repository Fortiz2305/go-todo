package main

import "github.com/Fortiz2305/go-todo/todo"

func main() {
	err := todo.Run()
	if err != nil {
		panic(err)
	}
}
