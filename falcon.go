package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/sshaparenko/falcon/pkg/colors"
)

//go:embed common/falcon.txt
var common string

//go:embed common/menu.txt
var menu string

func main() {
	colors.PrintMagenta(common)
	colors.PrintYellow(menu)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadBytes(10)
		command := string(noDelim(input))

		if err != nil {
			colors.PrintRed("Cannot read user input")
			os.Exit(1)
		}

		switch command {
		case "start":
			fmt.Printf("%v: Staring Falcon...\n", formatedTime())
			fmt.Printf("%v: Listening to your terminals...\n", formatedTime())
		case "stop":
			fmt.Printf("%v: Saving history...\n", formatedTime())
			colors.PrintGreen("[+] Falcon has finished")
			os.Exit(0)
		case "help":
			colors.PrintYellow(menu)
		default:
			fmt.Println("Unknown command")
		}
	}
}

func noDelim(input []byte) []byte {
	return input[:len(input)-1]
}

func formatedTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
