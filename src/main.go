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

	iostream "main.go/writer"
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

	// Create an empty shell script and then keep appending to it to it line by line
	flname_match := filename_pattern.FindStringSubmatch(krystal_code)
	flname := ""
	if target_platform == "linux" {
		flname += (flname_match[1] + ".sh")
	} else {
		flname += (flname_match[1] + ".ps1")
	}
	os.Create(flname)
	// Create an empty shell script and then keep appending to it line by line

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
		os.Remove(flname)
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
	comment_pattern := regexp.MustCompile(`^Com:\s*(.*?)$`)
	tilda_comment_pattern := regexp.MustCompile(`^~\s*Com:\s*(.*?)$`)
	display_pattern := regexp.MustCompile(`^display`)
	// Basic Regex validations defined here
	// Here we iterate over each line of the code and pass it into writer.go
	// We validate each line with the regex pattern using if else statements

	// Here are some variables that check or validate scopes
	// That is if a code is inside a for loop scope or it it is inside a multi line comment scope.
	is_multi_line_comment := false
	is_in_tildascope_win := false
	is_in_tildascope_linux := false
	// End of scope checking varialbes

	for index, line := range code {
		line := strings.TrimSpace(line)

		if comment_pattern.MatchString(line) || line == "" {
			// BY default comments do not get copied over to the generated script
			continue // Continue if a comment is found or if the line is empty

		} else if line == "MCom" {
			// Toggle multi line comment scope
			is_multi_line_comment = !is_multi_line_comment
			continue

		} else if line == "~win" {
			// Toggling tilda scope for powershell commands
			is_in_tildascope_win = !is_in_tildascope_win
			continue

		} else if line == "~linux" {
			// Toggling tilda scope for shell commands
			is_in_tildascope_linux = !is_in_tildascope_linux
			continue

		} else if is_in_tildascope_linux {
			// Code under tilda scope for linux should directly get copied into shell script
			// As code under tilda scope is expected to be valid regular shell commands
			// But code under ~linux should not get copied into windows powershell scripts
			if target_platform == "linux" {
				write_to_file(flname, line)
				continue
			}
			continue

		} else if is_in_tildascope_win {
			// Code under tilda scope for windows should directly get copied into powershell script
			// As code under tilda scope is expected to be valid regular powershell commands
			// But code under ~win should not get copied into linux shell scripts
			if target_platform == "windows" {
				write_to_file(flname, line)
				continue
			}
			continue

		} else if is_multi_line_comment {
			// If it's under a multi line scope, it's going to copy all lines under that scope into
			// the shell script
			write_to_file(flname, ("# " + line))
			continue

		} else if tilda_comment_pattern.MatchString(line) {
			// If the user wants the comment to get copied to the shell script
			// Then the comment must intentionally be prefixed with a ~
			write_to_file(flname, ("# " + tilda_comment_pattern.FindStringSubmatch(line)[1]))
			continue

		} else if display_pattern.MatchString(strings.ToLower(line)) {
			// If it matches the display pattern command
			content := iostream.Display(line, target_platform)
			if content != "err" {
				write_to_file(flname, content)
			} else {
				syntax_error(flname, index, line)
			}

		} else {
			// If the line does not match any regex pattern, it is a syntax error
			syntax_error(flname, index, line)
		}
	} // end of for loop

}

func write_to_file(flname string, content string) {
	// This function appends to the generated shell script
	file, _ := os.OpenFile(flname, os.O_APPEND|os.O_WRONLY, 0600)
	file.WriteString(content + "\n")
	file.Close()
}

func syntax_error(flname string, index int, line string) {
	// This function spits out a syntax error command and then deletes
	// the generated file
	fmt.Printf("\033[1;91mSyntax Error on line %d: \033[1;93m %s\033[1;0m\n", index+1, line)
	os.Remove(flname)
	os.Exit(1)
}
