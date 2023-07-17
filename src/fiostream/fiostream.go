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
	create_file_pattern_force := regexp.MustCompile(`^Make\s*File\s*\(\s*force\s*"(.*?)"\s*\)$`)

	if create_file_pattern.MatchString(line) {
		return return_code_for_touch(create_file_pattern.FindStringSubmatch(line)[1], osname)
	} else if create_file_pattern_force.MatchString(line) {
		return return_force_file(create_file_pattern_force.FindStringSubmatch(line)[1], osname)
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
	make_folder_subdirs_pattern := regexp.MustCompile(`^Make\s*Folder\s*\(\s*enable\s+sub\s+dirs\s*"(.*?)"\s*\)$`)
	make_folder_force_pattern := regexp.MustCompile(`^Make\s*Folder\s*\(\s*force\s*"(.*?)"\s*\)$`)

	if make_folder_pattern.MatchString(line) {
		return return_code_for_mkdir(make_folder_pattern.FindStringSubmatch(line)[1], osname)
	} else if make_folder_subdirs_pattern.MatchString(line) {
		return return_code_for_subdirs(make_folder_subdirs_pattern.FindStringSubmatch(line)[1], osname)
	} else if make_folder_force_pattern.MatchString(line) {
		return return_code_for_force_mkdir(make_folder_force_pattern.FindStringSubmatch(line)[1], osname)
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

func return_code_for_subdirs(str string, osname string) string {
	if osname == "linux" {
		return fmt.Sprintf("mkdir -p \"%s\"", str)
	} else {
		return fmt.Sprintf("New-Item -ItemType Directory -Path \"%s\"", str)
	}
}

func return_code_for_force_mkdir(str string, osname string) string {
	if osname == "linux" {
		return fmt.Sprintf("rm -rf \"%s\" && mkdir \"%s\"", str, str)
	} else {
		return fmt.Sprintf("Remove-Item -Recurse -Force \"%s\"; New-Item -ItemType Directory -Path \"%s\"", str, str)
	}
}

func return_force_file(str string, osname string) string {
	if osname == "linux" {
		return fmt.Sprintf("rm \"%s\" && touch \"%s\"", str, str)
	} else {
		return fmt.Sprintf("Remove-Item \"%s\"; New-Item \"%s\" -ItemType File", str, str)
	}
}
