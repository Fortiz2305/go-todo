# go-todo [![Go ReportCard](https://goreportcard.com/badge/github.com/Fortiz2305/go-todo)](https://goreportcard.com/report/github.com/Fortiz2305/go-todo)[![GoDoc](https://godoc.org/github.com/Fortiz2305/go-todo/todo?status.svg)](https://godoc.org/github.com/Fortiz2305/go-todo/todo)

A simple command line todo list written in Go.

## Installation

```
go get github.com/Fortiz2305/go-todo
```

## Usage

```
bash-3.2$ ./go-todo
./go-todo - go_todo

Commands:

    delete      Delete a task
    list        Show the task list
    save        Save a new task
    setStatus  Set the status of your task

    Use "./go-todo help <command>" for more information about a command.
```

## Options

### List tasks

```
bash-3.2$ ./go-todo list
+----+------------------------+-------------+--------+
| ID |          TASK          |    DATE     | STATUS |
+----+------------------------+-------------+--------+
|  1 | Make an awesome Go CLI | 2016-08-28  | OPEN   |
+----+------------------------+-------------+--------+
```

### Add new task

```
bash-3.2$ ./go-todo save "Fix a bug"

bash-3.2$ ./go-todo list
+----+------------------------+-------------+--------+
| ID |          TASK          |    DATE     | STATUS |
+----+------------------------+-------------+--------+
|  1 | Make an awesome Go CLI | 2016-08-28  | OPEN   |
|  2 | Fix a bug              | 2016-08-28  | OPEN   |
+----+------------------------+-------------+--------+
```

### Delete a task

```
bash-3.2$ ./go-todo delete 2

bash-3.2$ ./go-todo list
+----+------------------------+-------------+--------+
| ID |          TASK          |    DATE     | STATUS |
+----+------------------------+-------------+--------+
|  1 | Make an awesome Go CLI | 2016-08-28  | OPEN   |
+----+------------------------+-------------+--------+
```

### Change status

```
bash-3.2$ ./go-todo setStatus 1 "Pending to review"

bash-3.2$ ./go-todo list
+----+------------------------+-------------+-------------------+
| ID |          TASK          |    DATE     |      STATUS       |
+----+------------------------+-------------+-------------------+
|  1 | Make an awesome Go CLI | 2016-08-28  | Pending to review |
+----+------------------------+-------------+-------------------+
```

## Requirements

* Golang: http://golang.org/doc/install.html

