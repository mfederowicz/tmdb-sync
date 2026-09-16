package cli

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
)

// OpenBrowser best-effort opens url in the user's default browser. Failure is
// non-fatal: the URL was already printed for the user to open manually.
func OpenBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	_ = cmd.Start()
}

// WaitForEnter blocks until the user presses Enter on stdin.
func WaitForEnter() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
