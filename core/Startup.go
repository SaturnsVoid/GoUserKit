package core

func BasicStartup() {
	run("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V " + SimpleStartupName + " /t REG_SZ /F /D " + SimpleStartupPath)
}
