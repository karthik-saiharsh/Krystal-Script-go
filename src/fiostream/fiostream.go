package fiostream

import (
	"fmt"
	"regexp"
)

func MakeFile(line string, osname string) string {
	/*
		This function manages the CreateFile() command.
		We use regex to see if the line matches the expected pattern.
		If yes we get the content inside the "" and do further processing
	*/
	create_file_pattern := regexp.MustCompile(`^Make\s*File\s*\(\s*"(.*?)"\s*\)$`)

	if create_file_pattern.MatchString(line) {
		return return_code_for_touch(create_file_pattern.FindStringSubmatch(line)[1], osname)
	} else {
		return "err"
	}

}

func MakeFolder(line string, osname string) string {
	/*
		This function manages the CreateFile() command.
		We use regex to see if the line matches the expected pattern.
		If yes we get the content inside the "" and do further processing
	*/
	make_folder_pattern := regexp.MustCompile(`^Make\s*Folder\s*\(\s*"(.*?)"\s*\)$`)

	if make_folder_pattern.MatchString(line) {
		return return_code_for_mkdir(make_folder_pattern.FindStringSubmatch(line)[1], osname)
	} else {
		return "err"
	}

}

func return_code_for_touch(str string, osname string) string {
	if osname == "linux" {
		return fmt.Sprintf("touch \"%s\"", str)
	} else {
		return fmt.Sprintf("if (Test-Path \"%s\") {(Get-Item \"%s\").LastWriteTime = Get-Date} else {New-Item \"%s\" -ItemType File}", str, str, str)
	}
}

func return_code_for_mkdir(str string, osname string) string {
	if osname == "linux" {
		return fmt.Sprintf("mkdir \"%s\"", str)
	} else {
		return fmt.Sprintf("New-Item -ItemType Directory -Path \"%s\"", str)
	}
}
