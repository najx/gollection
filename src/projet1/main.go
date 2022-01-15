package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	var i int = 0

	if src == "" {
		log.Fatal("src is null..")
	}

	if old == "" {
		log.Fatal("old is null..")
	}

	if new == "" {
		log.Fatal("old is null..")
	}

	file, err := os.Open(src)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		t := scanner.Text()
		found, res, count := ProcessLine(t, old, new)
		fmt.Printf("%s\n", res)
		if found == true {
			lines = append(lines, i)
			occ += count
		}
		i++
	}
	fmt.Printf("Number of occurences of Go: %d\n", occ)
	fmt.Printf("Number of lines: %d\n", i)
	return i, lines, err
}

func ProcessLine(Line, old, new string) (found bool, res string, occ int) {
	return strings.Contains(Line, old),
		strings.Replace(Line, old, new, 0),
		strings.Count(Line, old)
}

func main() {
	var src string = "file.txt"
	var old string = "Go"
	var new string = "Python"

	fmt.Printf("== Summary ==\n")
	_, value, err := FindReplaceFile(src, old, new)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Lines: %v\n", value)
	fmt.Printf("== End of Summary ==\n")
}
