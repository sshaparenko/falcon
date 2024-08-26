package main

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/sshaparenko/falcon/pkg/colors"
	"github.com/sshaparenko/falcon/pkg/commands"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Not a valid command. Write falcon --help to see a list of avalable commands")
		os.Exit(1)
	}

	commands.Help(os.Args[1:])

	command := os.Args[1]
	flags := os.Args[2:]

	switch command {
	case "run":
		commands.Run(flags)
		fmt.Printf("%v: Listening to your terminals...\n", formatedTime())
	case "stop":
		fmt.Printf("%v: Saving history...\n", formatedTime())
		colors.PrintGreen("[+] Falcon has finished")
		os.Exit(0)
	case "pid":
		commands.Pid(flags)
	default:
		fmt.Printf("%s is unknown command. Write falcon --help to see a list of avalable commands\n", command)
	}
}

func formatedTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
