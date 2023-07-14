package writer

import (
	"fmt"
	"os"
	"regexp"
)

func write_to_script(flname string, content string) {
	f, err := os.OpenFile(flname, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("\033[1;91mUnable to write to file %s\033[1;0m", flname)
		os.Exit(1)
	}
	f.WriteString(content)
	f.Close()
}

func Display(line string, index int, flname string, osname string) {
	display_pattern := regexp.MustCompile(`^Display\s*\(\s*"(.*?)"\s*\)$`)
	display_pattern_literal := regexp.MustCompile(`^Display\s*\(\s*literal\s*"(.*?)"\s*\)$`)
	display_pattern_no_new_line := regexp.MustCompile(`^Display\s*\(\s*no\s+new\s+line\s*"(.*?)"\s*\)$`)

	if display_pattern.MatchString(line) {
		match := display_pattern.FindStringSubmatch(line)
		if osname == "linux" {
			content := "echo \"" + match[1] + "\"\n"
			write_to_script(flname, content)
		} else {
			// Win Code here
		}

	} else if display_pattern_literal.MatchString(line) {
		match := display_pattern_literal.FindStringSubmatch(line)

		if osname == "linux" {
			content := "echo '" + match[1] + "'\n"
			write_to_script(flname, content)
		} else {
			// Win Code here
		}

	} else if display_pattern_no_new_line.MatchString(line) {
		match := display_pattern_no_new_line.FindStringSubmatch(line)

		if osname == "linux" {
			content := "echo -n \"" + match[1] + "\"\n"
			write_to_script(flname, content)
		} else {
			// Win Code here
		}

	} else {
		er := "\n\n033[1;91m Syntax Error on line %d: \033[1;93m%s\033[1;0m\n"
		fmt.Printf(er, index+1, line)
		os.Remove(flname)
		os.Exit(1)
	}

}
