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

	if display_pattern.MatchString(line) {
		match := display_pattern.FindStringSubmatch(line)
		if osname == "linux" {
			content := "echo \"" + match[1] + "\"\n"
			write_to_script(flname, content)
		} else {
			content := "echo " + match[1] + "\n"
			write_to_script(flname, content)
		}
	} else {
		er := "\n\n033[1;91m Error on line %d: \033[1;93m%s\033[1;92m\nThe Correct syntax is Display(\"< Your Text here>\")\033[1;0m\n"
		fmt.Printf(er, index+1, line)
		os.Remove(flname)
		os.Exit(1)
	}

}
