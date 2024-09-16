package commands

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/creack/pty"
	"github.com/sshaparenko/falcon/pkg/colors"
	"github.com/sshaparenko/falcon/pkg/terminal"
)

var (
	//go:embed menu.txt
	falcon_help string
	//go:embed run.txt
	run_help string
	//go:embed pid.txt
	pids_help string
	//go:embed falcon.txt
	info string
)

func Run(args []string) {
	helpFlag := NewFlag("help", false, "show the menu with available commands")

	fs, wrapper := BuildFlagSet("run", []*Flag{helpFlag})

	if err := fs.Parse(args); err != nil {
		log.Fatalf("Failed to parse flags of command <run>: %s\n", err.Error())
	}
	fName := wrapper.CheckActive()
	switch fName {
	case "help":
		colors.PrintMagenta(info)
		fmt.Println(run_help)
		return
	default:
		fmt.Printf("%v: Staring Falcon...\n", formatedTime())
		terminal.Run()
	}
}

func Help(args []string) {
	helpFlag := NewFlag("help", false, "show menu with available commands")

	fs, wrapper := BuildFlagSet("falcon", []*Flag{helpFlag})

	if err := fs.Parse(args); err != nil {
		log.Fatalf("Failed to parse flags of command <falcon>: %s\n", err.Error())
	}

	fName := wrapper.CheckActive()
	switch fName {
	case "help":
		colors.PrintMagenta(info)
		fmt.Println(falcon_help)
		return
	}
}

func Pid(args []string) {
	trackedFlag := NewFlag("tracked", false, "Show shell PIDs that are tracked by Falcon")
	allFlag := NewFlag("all", false, "Show PIDs of all shells")
	helpFlag := NewFlag("help", false, "Show menu with available commands")

	fs, wrapper := BuildFlagSet("pids", []*Flag{trackedFlag, allFlag, helpFlag})

	if err := fs.Parse(args); err != nil {
		log.Fatalf("Failed to parse flags of command <pids>: %s\n", err.Error())
	}

	fName := wrapper.CheckActive()
	switch fName {
	case "help":
		colors.PrintMagenta(info)
		fmt.Println(pids_help)
		return
	case "tracked":
		fmt.Println(getAllPIDs())
		return
	case "all":
		fmt.Println(getAllPIDs())
		return
	default:
		fmt.Println(os.Getppid())
	}
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
