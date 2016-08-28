# go-todo [![Go ReportCard](https://goreportcard.com/badge/github.com/Fortiz2305/go-todo)](https://goreportcard.com/report/github.com/Fortiz2305/go-todo)[![GoDoc](https://godoc.org/github.com/Fortiz2305/go-todo/todo?status.svg)](https://godoc.org/github.com/Fortiz2305/go-todo/todo)

A simple command line todo list written in Go.

## Usage

```
bash-3.2$ ./go-todo
./go-todo - go_todo

Commands:

    delete      Delete a task
    list        Show the task list
    save        Save a new task
    set_status  Set the status of your task

    Use "./go-todo help <command>" for more information about a command.
```

## Options

### List tasks

```
bash-3.2$ ./go-todo list
+------------------------+-------------+--------+
|          TASK          |    DATE     | STATUS |
+------------------------+-------------+--------+
| Make an awesome Go CLI | 2016-08-25  | OPEN   |
+------------------------+-------------+--------+
```

### Add new task

```
bash-3.2$ ./go-todo save "Fix a bug"

bash-3.2$ ./go-todo list
+------------------------+-------------+--------+
|          TASK          |    DATE     | STATUS |
+------------------------+-------------+--------+
| Make an awesome Go CLI | 2016-08-25  | OPEN   |
| Fix a bug              | 2016-08-25  | OPEN   |
+------------------------+-------------+--------+
```

### Delete a task

```
bash-3.2$ ./go-todo delete "Fix a bug"

bash-3.2$ ./go-todo list
+------------------------+-------------+--------+
|          TASK          |    DATE     | STATUS |
+------------------------+-------------+--------+
| Make an awesome Go CLI | 2016-08-25  | OPEN   |
+------------------------+-------------+--------+
```

### Change status

```
bash-3.2$ ./go-todo set_status "Make an awesome Go CLI" "Pending to review"

bash-3.2$ ./go-todo list
+------------------------+-------------+-------------------+
|          TASK          |    DATE     |      STATUS       |
+------------------------+-------------+-------------------+
| Make an awesome Go CLI | 2016-08-25  | Pending to review |
+------------------------+-------------+-------------------+
```

## Requirements

* Golang: http://golang.org/doc/install.html

