package misc

import (
	"fmt"
	"regexp"
)

func Assign(line string, osname string) string {

	assign_pattern := regexp.MustCompile(`^Assign\s*\(\s*(.*?)\s*,\s*([a-zA-Z_]+)\s*\)$`)

	if assign_pattern.MatchString(line) {
		return vars(assign_pattern.FindStringSubmatch(line)[1], assign_pattern.FindStringSubmatch(line)[2], osname)
	} else {
		return "err"
	}
}

func vars(data string, varname string, osname string) string {
	if osname == "linux" {
		return fmt.Sprintf("%s=%s", varname, data)
	} else {
		return fmt.Sprintf("$%s = %s", varname, data)
	}
}
