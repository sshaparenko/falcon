package commands

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/creack/pty"
	"github.com/sshaparenko/falcon/pkg/colors"
	"github.com/sshaparenko/falcon/pkg/socket"
)

//go:embed menu.txt
var falcon_help string

//go:embed run.txt
var run_help string

//go:embed pid.txt
var pids_help string

//go:embed falcon.txt
var info string

func Run(args []string) {
	fs := flag.NewFlagSet("run", flag.ContinueOnError)
	sm := fs.Bool("s", false, "start falcon with silent mode")
	help := fs.Bool("help", false, "show the menu with available commands")
	if err := fs.Parse(args); err != nil {
		log.Fatalf("Failed to parse flags of command <run>: %s\n", err.Error())
	}
	if *help {
		colors.PrintMagenta(info)
		fmt.Println(run_help)
		os.Exit(0)
	}
	fmt.Printf("%v: Staring Falcon...\n", formatedTime())
	socket.Run()
	fmt.Printf("%v: Silent Mode: %t\n", formatedTime(), *sm)

}

func Help(args []string) {
	fs := flag.NewFlagSet("falcon", flag.ContinueOnError)
	help := fs.Bool("help", false, "show menu with available commands")
	if err := fs.Parse(args); err != nil {
		log.Fatalf("Failed to parse flags of command <falcon>: %s\n", err.Error())
	}
	if *help {
		colors.PrintMagenta(info)
		fmt.Println(falcon_help)
		os.Exit(0)
	}
}

func Pid(args []string) {
	fs := flag.NewFlagSet("pids", flag.ContinueOnError)

	trackedFlag := fs.Bool("tracked", false, "Show shell PIDs that are tracked by Falcon")
	allFlag := fs.Bool("all", false, "Show PIDs of all shells")
	// currentFlag := fs.Bool("current", false, "Show PID of a current shell")
	helpFlag := fs.Bool("help", false, "Show menu with available commands")

	if err := fs.Parse(args); err != nil {
		log.Fatalf("Failed to parse flags of command <pids>: %s\n", err.Error())
	}
	if *helpFlag {
		colors.PrintMagenta(info)
		fmt.Println(pids_help)
		os.Exit(0)
	}
	if *trackedFlag {
		fmt.Println(getAllPIDs())
		os.Exit(0)
	}
	if *allFlag {
		fmt.Println(getAllPIDs())
		os.Exit(0)
	}
	fmt.Println(os.Getppid())
}

func formatedTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getAllPIDs() []string {
	buff := make([]byte, 1024)
	c := exec.Command("pgrep", "bash")
	f, err := pty.Start(c)

	defer func() { _ = f.Close() }()

	if err != nil {
		log.Fatalf("Failed to run PTY command: %s\n", err.Error())
	}

	n, err := f.Read(buff)
	if err != nil {
		msg := fmt.Sprintf("Failed to get PTY command result: %s\n", err.Error())
		log.Fatal(msg)
	}

	return readPIDsFromBytes(buff[:n])
}

func readPIDsFromBytes(pids []byte) []string {
	result := []string{}
	var delim int
	for i, p := range pids {
		if p == 10 {
			pid := string(pids[delim : i-1])
			delim = i + 1
			result = append(result, pid)
		}
	}
	return result
}
