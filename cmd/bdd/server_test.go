package tests

import (
	"net/http"
	"os/exec"
	"time"
)

// TODO need to look for free port
const address = ":8080"

func startServer() (*exec.Cmd, string, error) {
	cmd := exec.Command("../paymentserver/paymentserver", "-l", address)
	err := cmd.Start()
	if err != nil {
		return cmd, "", err
	}
	url := "http://" + address
	for i := 1; i < 10; i++ {
		_, err = http.Get(url)
		if err == nil {
			return cmd, url, nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	stopServer(cmd)
	return cmd, url, nil
}

func stopServer(cmd *exec.Cmd) {
	defer cmd.Wait()
	cmd.Process.Kill()
}
