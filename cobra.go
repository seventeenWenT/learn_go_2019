package main

import (
	"basic/cobra/app"
	"fmt"
	"os"
)

func main(){
	command := app.NewCobraCommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}