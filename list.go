package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
	//"strings"
	"encoding/json"
	//"time"
)

func todo_list(tasks_file string) *commander.Command {
	list := func(cmd *commander.Command, args []string) error {
		file, err := ioutil.ReadFile(tasks_file)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(file))
		tasksList := make([]Task, 0)
		json.Unmarshal(file, &tasksList)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Task", "Date", "Status"})

		for i := 0; i < len(tasksList); i++ {
			data := [][]string{
				[]string{tasksList[i].Todo, tasksList[i].Date.Format("2006-01-12"), tasksList[i].Status},
			}
			for _, v := range data {
		    table.Append(v)
			}
		}

		table.Render()

		return nil
	}

	return &commander.Command{
		Run:       list,
		UsageLine: "list [options]",
		Short:     "list go_todo",
	}
}

//func (tc []Task) FromJson(jsonFile []byte) error {
//	var data = &tc.Pool
//	fmt.Println(string(jsonFile))
//	return json.Unmarshal(jsonFile, data)
//}
