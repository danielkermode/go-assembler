package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/danielkermode/assembler/convert"
)

var linenumber int

func main() {
	args := os.Args
	if len(args) < 2 {
		help()
		return
	}
	filename := args[1]
	// open input file
	fi, err := os.Open(filename)
	convert.Check(err)
	// close fi on exit and check for its returned error
	defer func() {
		closeErr := fi.Close()
		convert.Check(closeErr)
	}()

	// open output file
	output := strings.TrimSuffix(filename, filepath.Ext(filename))
	fo, err := os.Create(output + ".hack")
	convert.Check(err)
	// close fo on exit and check for its returned error
	defer func() {
		err := fo.Close()
		convert.Check(err)
	}()

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		linenumber++
		// next part strips comments and then skips line if it's empty
		text := stripComments(scanner.Text())
		if len(text) == 0 {
			continue
		}
		// convert the command
		result, err := convert.Convert(text, linenumber)
		convert.Check(err)
		_, writeErr := fo.WriteString(result + "\n")
		convert.Check(writeErr)
		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func stripComments(line string) string {
	comment := regexp.MustCompile("\\/\\/.*")
	space := regexp.MustCompile("\\s*")
	line = comment.ReplaceAllString(line, "")
	return space.ReplaceAllString(line, "")
}

func help() {
	fmt.Println("\nYou must enter a file for the assembler.")
}
