package crimson

import (
	"os/exec"
	"runtime"
	"strconv"
)

// OpenInBrowser opens http://127.0.0.1:<port><page> in default browser
func OpenInBrowser() {
	OpenUrlInBrowser("http://127.0.0.1:" + strconv.Itoa(GetServerPort()) + GetBrowserOpenPage())
}

// OpenUrlInBrowser opens URL in default browser
func OpenUrlInBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "linux":
		// TODO: This case has not been test in Linux
		err = exec.Command("open", url).Start()
	default:
		// TODO: Add more cases of OS
		PrintWarning("Failed to open browser, unknown OS: " + runtime.GOOS)
		return
	}
	if err != nil {
		PrintError("Failed to open browser: " + err.Error())
	}
}
