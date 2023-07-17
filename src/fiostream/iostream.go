package iostream

import (
	"fmt"
	"regexp"
)

func Display(line string, osname string) string {
	/*
		This function manages the Display() command.
		We use regex to see if the line matches the expected pattern.
		If yes we get the content inside the "" and check to see if any further arguments are passed in
	*/
	display_pattern := regexp.MustCompile(`^Display\s*\(\s*"(.*?)"\s*\)$`)
	display_pattern_nonewline := regexp.MustCompile(`^Display\s*\(\s*no\s+new\s+line\s*"(.*?)"\s*\)$`)
	display_pattern_literal := regexp.MustCompile(`^Display\s*\(\s*literal\s*"(.*?)"\s*\)$`)
	display_pattern_literal_nonewline := regexp.MustCompile(`^Display\s*\(\s*literal\s+no\s+new\s+line\s*"(.*?)"\s*\)$`)

	if display_pattern.MatchString(line) {
		return return_regular_display(display_pattern.FindStringSubmatch(line)[1], osname)

	} else if display_pattern_nonewline.MatchString(line) {
		return return_nonewline_display(display_pattern_nonewline.FindStringSubmatch(line)[1], osname)

	} else if display_pattern_literal.MatchString(line) {
		return return_literal_display(display_pattern_literal.FindStringSubmatch(line)[1], osname)

	} else if display_pattern_literal_nonewline.MatchString(line) {
		return return_literal_nonewline_display(display_pattern_literal_nonewline.FindStringSubmatch(line)[1], osname)

	} else {
		return "err"
	}
}

// These functions are what actually return the line that gets appended at the EOF that is generated

func return_regular_display(str string, osname string) string {
	if osname == "linux" {
		return fmt.Sprintf("echo \"%s\"", str)
	} else {
		return fmt.Sprintf("Write-Output \"%s\"", str)
	}
}

func return_nonewline_display(str string, osname string) string {
	if osname == "linux" {
		return fmt.Sprintf("echo -n \"%s\"", str)
	} else {
		return fmt.Sprintf("Write-Output -NoNewline \"%s\"", str)
	}
}

func return_literal_display(str string, osname string) string {
	if osname == "linux" {
		return fmt.Sprintf("echo '%s'", str)
	} else {
		return fmt.Sprintf("Write-Output '%s'", str)
	}
}

func return_literal_nonewline_display(str string, osname string) string {
	if osname == "linux" {
		return fmt.Sprintf("echo -n '%s'", str)
	} else {
		return fmt.Sprintf("Write-Output -NoNewline '%s'", str)
	}
}
