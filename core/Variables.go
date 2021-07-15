package core

import (
	"os"
	"syscall"
)

var (
	SimpleStartupName string = "WindowsUpgrade"
	SimpleStartupPath string = os.Args[0]

	ntdll                   = syscall.MustLoadDLL("ntdll.dll")
	NtSetInformationProcess = ntdll.MustFindProc("NtSetInformationProcess")
)
