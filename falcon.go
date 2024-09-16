package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sshaparenko/falcon/pkg/colors"
	"github.com/sshaparenko/falcon/pkg/commands"
)

func main() {

	var (
		command string
		flags   []string
		logger  *log.Logger
	)

	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// checkForCommand()
	switch len(os.Args) {
	case 1:
		commands.Falcon(os.Args)
		return
	case 2:
		if os.Args[1] == "-help" || os.Args[1] == "--help" {
			commands.Falcon(os.Args)
			return
		}
	}
	// commands.Falcon(os.Args[1:])

	command = os.Args[1]
	flags = os.Args[2:]

	switch command {
	case "run":
		commands.Run(flags)
		logger.Println("Listening to your terminals...")
		//fmt.Printf("%v: Listening to your terminals...\n", formatedTime())
	case "stop":
		logger.Println("Saving history...")
		//fmt.Printf("%v: Saving history...\n", formatedTime())
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
