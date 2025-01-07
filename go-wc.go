package main

import (
	"fmt"
	"os"
)

func CountBytes(data string) int {
	return len(data)
}

func CountChars(data string) int {
	return len([]rune(data))
}

func CountLines(data string) int {
	const newline byte = '\n'
	count := 0

	for i := range len(data) {
		if data[i] == newline {
			count++
		}
	}
	if data[len(data)-1] != newline {
		count++
	}
	return count
}

func Delimeter(b byte) bool {
	const space byte = ' '
	const newline byte = '\n'

	if b == space || b == newline {
		return true
	}
	return false
}

func CountWords(data string) int {
	const space byte = ' '
	const newline byte = '\n'
	count := 0

	insideWord := false
	for i := range len(data) {
		if Delimeter(data[i]) && insideWord == false {
			insideWord = true
			count++
		}

		if Delimeter(data[i]) {
			insideWord = false
		}
	}

	return count

}

func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(CountBytes(string(data)))
	fmt.Println(CountChars(string(data)))
	fmt.Println(CountLines(string(data)))
	fmt.Println(CountWords(string(data)))

}
