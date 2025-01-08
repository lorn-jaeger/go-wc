package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

func CountLines(file []byte) int {
	var count int

	for _, byte := range file {
		if byte == '\n' {
			count++
		}
	}

	return count
}

func CountWords(file []byte) int {
	var count int
	insideWord := false

	for _, b := range file {
		if unicode.IsSpace(rune(b)) {
			if insideWord {
				insideWord = false
			}
		} else {
			if !insideWord {
				insideWord = true
				count++
			}
		}
	}
	return count
}

func CountChars(file []byte) int {
	return len([]rune(string(file)))
}

func CountBytes(file []byte) int {
	return len(file)
}

func ReadFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func ReadStdin() []byte {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return stdin
}

func FmtResults(results []int, filename string) string {
	maxWidth := 0
	for _, result := range results {
		width := len(fmt.Sprint("%d", result))
		if width > maxWidth {
			maxWidth = width
		}
	}

	output := ""

	for _, result := range results {
		output += fmt.Sprintf("%*d", maxWidth-1, result)
	}

	if filename != "" {
		output += " " + filename
	}

	return output
}

func main() {
	var filename string
	var file []byte
	var results []int

	lineFlag := flag.Bool("l", false, "Count lines in file")
	wordFlag := flag.Bool("w", false, "Count words in file")
	charFlag := flag.Bool("m", false, "Count chars in file")
	byteFlag := flag.Bool("c", false, "Count bytes in file")

	flag.Parse()
	args := flag.Args()

	if !*lineFlag && !*wordFlag && !*charFlag && !*byteFlag {
		*lineFlag, *wordFlag, *byteFlag = true, true, true
	}

	if len(args) > 0 {
		filename = args[0]
		file = ReadFile(filename)

	} else {
		filename = ""
		file = ReadStdin()
	}

	if *lineFlag {
		results = append(results, CountLines(file))
	}

	if *wordFlag {
		results = append(results, CountWords(file))
	}

	if *charFlag {
		results = append(results, CountChars(file))
	}

	if *byteFlag {
		results = append(results, CountBytes(file))
	}

	fmt.Println(FmtResults(results, filename))

}
