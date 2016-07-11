package main

import (
	"fmt"
	"github.com/gonuts/commander"
	//"github.com/olekukonko/tablewriter"
	"io/ioutil"
	//"os"
	//"strings"
	"encoding/json"
	"time"
)

type Tasks struct {
	Id int 											`json:"Id"`
	Todo string									`json:"Todo"`
	Date time.Time							`json:"Date"`
	Status string								`json:"Status"`
}

func todo_list(tasks_file string) *commander.Command {
	list := func(cmd *commander.Command, args []string) error {
		file, err := ioutil.ReadFile(tasks_file)
		if err != nil {
			panic(err)
		}

		var jsontype Tasks
		json.Unmarshal(file, &jsontype)
		fmt.Printf("Results: %+v", jsontype.Id)
		//jsonParser := json.NewDecoder(file)
		//if err = jsonParser.Decode(&tasks); err != nil {
	//		panic(err)
		//}

		//data := [][]string{
	//		[]string{tasks.Todo, tasks.Date.Format("2006-01-12"), tasks.Status},
	//	}

//		table := tablewriter.NewWriter(os.Stdout)
//		table.SetHeader([]string{"Task", "Date", "Status"})

	//	for _, v := range data {
	//		table.Append(v)
	//	}
	//	table.Render()

		return nil
	}

	return &commander.Command{
		Run:       list,
		UsageLine: "list [options]",
		Short:     "list go_todo",
	}
}
