package crimson

import (
	"os/exec"
	"runtime"
	"strconv"
)

// OpenInBrowser opens http://127.0.0.1:port in default browser
func OpenInBrowser() {
	OpenUrlInBrowser("http://127.0.0.1:" + strconv.Itoa(GetServerPort()))
}

// OpenUrlInBrowser opens URL in default browser
func OpenUrlInBrowser(url string) {
	var cmd string
	switch runtime.GOOS {
	case "windows":
		cmd = "start"
	case "linux":
		cmd = "open"
	default:
		PrintWarning("Failed to open browser, unknown OS: " + runtime.GOOS)
		return
	}
	err := exec.Command(cmd, url).Start()
	if err != nil {
		PrintError("Failed to open browser: " + err.Error())
	}
}
