package crimson

import (
	"os/exec"
	"runtime"
	"strconv"
)

func OpenInBrowser() {
	OpenUrlInBrowser("http://127.0.0.1:" + strconv.Itoa(GetServerPort()))
}

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
	exec.Command(cmd, url)
}
