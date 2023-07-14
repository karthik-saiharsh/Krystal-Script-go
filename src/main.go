/*
This is the official source code for Krystal script
Author: I.Karthik Saiharsh
lisence: MIT
Date of creation: 13 July 2023
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"main.go/writer"
)

func main() {
	args := os.Args
	/*
		You pass in 3 arguments into the Transpiler which is a binary named Krystal
		1) Path to your code
		2) Target Operating system
		3) Weather you want to run it or build it
		Example: ./Krystal main.KS linux run (or) ./Krystal main.KS windows build
	*/

	if len(args) != 3 {
		fmt.Println("\033[1;91m Error: \033[1;93m Expected 2 arguments(filename, target OS) recieved \033[1;0m", len(args))
		os.Exit(1)
	}

	krystal_code := args[1]
	target_platform := args[2]

	filename_pattern := regexp.MustCompile(`(.*?)\.KS$`) // Making sure that the filename has the right extension

	if !(target_platform == "windows" || target_platform == "linux") {
		fmt.Println("\033[1;91mError: \033[1;93mExpected a valid Operating System(windows/linux) recieved\033[1;0m", target_platform)
		os.Exit(1)
	} else if !(filename_pattern.MatchString(krystal_code)) {
		fmt.Println("\033[1;91mError: \033[1;93mExpected a valid Krystal Script, ending in *.KS recieved\033[1;0m", krystal_code)
		os.Exit(1)
	} else {
		fmt.Println("\033[1;92mOperating system set to\033[1;0m", target_platform)
		fmt.Println("\033[1;92mLoaded File\033[1;0m", krystal_code)
	}

	// Create an empty shell script and then keep appending to it
	flname_match := filename_pattern.FindStringSubmatch(krystal_code)
	flname := ""
	if target_platform == "linux" {
		flname += (flname_match[1] + ".sh")
	} else {
		flname += (flname_match[1] + ".bat")
	}
	os.Create(flname)
	// Create an empty shell script and then keep appending to it

	/*
		The way the transpiler works is, it reads the code file provided, line by line and appends each line as an element to a slice
		Then We iterate over the slice, element by element matching the line with some regex pattern
		And then depending on the OS we append code in either .sh or .bat file
		The actual appending part will be inside writer.go
	*/

	code := make([]string, 0, 100) // By default we make room for 100 lines of code

	file, err := os.Open(krystal_code)
	scanner := bufio.NewScanner(file) // We're gonna read the file line by line and append it to code
	if err != nil {
		fmt.Println("\033[1;91mUnable to Open the code file you passed in\033[1;0m")
		os.Exit(1)
	}

	for scanner.Scan() {
		code = append(code, strings.TrimSpace(scanner.Text()))
	}
	file.Close() // I didn't use a defer on purpose 🙃

	if len(code) == 0 {
		// This step is obvious, we can't transpile a blank file....
		fmt.Println("\033[1;91mYou have provided an empty file. What should I do with it?\033[1;0m")
		os.Exit(1)
	}

	/*
		Next I'm going to define a bunch of regex Patterns.
		Then I'm going to iterate over the slice "code" and parse each line to check if it matches the regex pattern
		If yes then I'll call functions from writer.go to append lines to file.


		NOTE: Since regex is used to parse code, the parsing capabilities are somewhat limited.
		Hence the syntax should be followed strictly, write one-lineers or terse code won't really work 🥲

		We'll convert each line to lowercase and use regex to check if it matches the overall format of a command.
		If yes, then the original line is passed on to writer.go where it'll use regex again to check more carefully
		if the syntax is met.
		This helps in providing useful errors to the programmer.
	*/

	// Basic Regex validations defined here
	display_pattern := regexp.MustCompile(`^display\s*\((.*?)\)$`)
	// Basic Regex validations defined here
	// Here we iterate over each line of the code and pass it into writer.go
	// We validate each line with the regex pattern using if else statements
	for index, line := range code {
		if display_pattern.MatchString(strings.ToLower(line)) {
			writer.Display(line, index, flname, target_platform)
		}
	}

}
