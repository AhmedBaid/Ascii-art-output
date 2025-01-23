package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/functions"
)

func main() {
	argument := ""
	fileName := "files/standard.txt"
	banner := ""
	switch {
	case len(os.Args) == 2:
		argument = os.Args[1]

	case len(os.Args) == 3:
		if strings.HasPrefix(os.Args[1], "--output=") {
			argument = os.Args[2]
			if strings.HasSuffix(os.Args[1], ".txt") {
				banner = os.Args[1][9:]
			} else {
				fmt.Println("Invalid output file name  : usage => banner.txt")
				return
			}
		} else {
			argument = os.Args[1]
			if strings.HasSuffix(os.Args[2], ".txt") {
				fileName = "files/" + os.Args[2]
			} else {
				fileName = "files/" + os.Args[2] + ".txt"
			}
		}
	case len(os.Args) == 4:
		if strings.HasPrefix(os.Args[1], "--output=") {
			argument = os.Args[2]
			if strings.HasSuffix(os.Args[1], ".txt") {
				banner = os.Args[1][9:]
			} else {
				fmt.Println("Invalid output file name  : usage => banner.txt")
				return
			}
			if strings.HasSuffix(os.Args[3], ".txt") {
				fileName = "files/" + os.Args[3]
			} else {
				fileName = "files/" + os.Args[3] + ".txt"
			}
		}
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	functions.Handler(argument, banner, fileName)
}
