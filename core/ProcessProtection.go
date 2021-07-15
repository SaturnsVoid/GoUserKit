package core

func SetInformationProcess(hProcess uintptr, processInformationClass int, processInformation int, processInformationLength int) {
	_, _, _ = NtSetInformationProcess.Call(hProcess, uintptr(processInformationClass), uintptr(processInformation), uintptr(processInformationLength))
}
