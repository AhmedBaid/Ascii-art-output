package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Handler(argument, banner, fileName string){
	fileName=strings.ToLower(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error in opening the file  USAGE ", err)
		return
	}
	defer file.Close()
	SlicesAscii := []string{}
	MapAscii := map[rune][]string{}
	count := 0
	espace := ' '
	myscanner := bufio.NewScanner(file)
	for myscanner.Scan() {
		text := myscanner.Text()
		if text != "" {
			SlicesAscii = append(SlicesAscii, text)
			count++
		}
		if count == 8 {
			MapAscii[espace] = SlicesAscii
			espace++
			SlicesAscii = []string{}
			count = 0
		}
	}
	if len(argument) == 0 {
		return
	}

	if Isprintable(argument) {
		fmt.Println("Isprintable charachtere not allowed")
		return
	}

	Splitslice := strings.Split(argument, "\\n")
	var lastResult string

	if strings.ReplaceAll(argument, "\\n", "") == "" {
		for i := 0; i < strings.Count(argument, "\\n"); i++ {
			lastResult += "\n"
		}
	} else {
		lastResult = PrintAscii(Splitslice, MapAscii)
	}

	if banner != "" {
		err := os.WriteFile(banner, []byte(lastResult), 0o644)
		if err != nil {
			fmt.Println("Error", err)
			return
		}

	} else {
		fmt.Print(lastResult)
	}
}
