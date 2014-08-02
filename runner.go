package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"math/rand"
	"time"
	"path/filepath"
	"path"
	"strings"
	"regexp"
)

func usage() {
	fmt.Printf("usage: %s [template file]\n", os.Args[0])
	os.Exit(1)
}

func main() {
	// The template upon which we're basing our scam
	var template string

	// Maps a word type ("noun", etc) to a string list
	lists := make(map[string]*StringList)

	// Set up PRNG so our scams are unique per run
	rand.Seed( time.Now().UTC().UnixNano())

	if len(os.Args) != 2 {
		usage()
	}

	fname := os.Args[1]
	fileBytes, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("Error trying to read the template file '%s'\n\n", fname)
		usage()
	}

	template = string(fileBytes)

	// Read in all *.txt files to populate string lists
	//
	// TODO: Make this dynamic rather than based on specific local files
	dataFiles, err := filepath.Glob("./data/wordlists/*.txt")
	if err != nil {
		fmt.Println("Error trying to read word lists")
		os.Exit(1)
	}

	// Pull all 'wordlist' files and populate the StringList array
	for _, file := range dataFiles {
		fileBytes, err = ioutil.ReadFile(file)
		fileData := string(fileBytes)
		listname := strings.Replace(path.Base(file), ".txt", "", -1)
		lists[listname] = NewStringList()

		for _, str := range strings.Split(fileData, "\n") {
			if strings.TrimSpace(str) != "" {
				lists[listname].AddString(str)
			}
		}
	}

	// Throw out errors if any lists are empty
	for listname, list := range lists {
		if list.masterList.Len() == 0 {
			fmt.Printf("FATAL: List '%s' exists but has no data!\n", listname)
			os.Exit(1)
		}
	}

	// Read the template and populate data
	tvarRegex := regexp.MustCompile(`{{([^}]*)}}`)
	for {
		foundStrings := tvarRegex.FindStringSubmatch(template)
		if foundStrings == nil {
			break
		}

		// Set up a variable to hold the replacement value
		replacementValue := ""

		// Store the full match in an alias for easier replacing later
		fullMatch := foundStrings[0]

		// Handle possible variable assignments
		data := strings.Split(foundStrings[1], "->")
		listname := data[0]
		variable := ""
		if len(data) == 2 {
			variable = data[1]
		}

		// See if the list exists and warn if not
		list := lists[listname]
		if list == nil {
			fmt.Printf("ERROR: List '%s' needed but doesn't exist\n", listname)
		} else {
			replacementValue = list.RandomString()
		}

		if variable != "" {
			lists[variable] = NewStringList()
			lists[variable].AddString(replacementValue)
		}

		template = strings.Replace(template, fullMatch, replacementValue, 1)
	}

	fmt.Println(template)
}
