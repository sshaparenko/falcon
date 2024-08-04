package colors

import "fmt"

func PrintGreen(text string) {
	fmt.Printf("\033[32m%s\033[0m\n", text)
}

func PrintYellow(text string) {
	fmt.Printf("\033[33m%s\033[0m\n", text)
}

func PrintRed(text string) {
	fmt.Printf("\033[31m%s\033[0m\n", text)
}

func PrintMagenta(text string) {
	fmt.Printf("\033[35m%s\033[0m\n", text)
}
