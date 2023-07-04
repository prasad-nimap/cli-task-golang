package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/thejunghare/task/cmd"
	"github.com/thejunghare/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	must(db.Init(dbPath))
	/*
		Above line can written in following way aswell
		err := db.Init(dbPath)
		if err != nil {
			panic(err)
		}
	*/

	/* Execute the root command: syntax -> cmd.cmdName.Execute() */
	must(cmd.RootCmd.Execute())
	/* 
		Above line could be write as...
		cmd.RootCmd.Execute()
	*/
}

/* Helper error function */
func must(err error) {
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
