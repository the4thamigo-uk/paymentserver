package tests

import (
	"flag"
	"net/http"
	"os/exec"
	"time"
)

var address = flag.String("address", ":8080", "Address for the test server to listen on")
var serverFilename = flag.String("server", "./paymentserver", "Filename of the server executable")

func init() {
	flag.Parse()
}

func startServer() (*exec.Cmd, string, error) {
	cmd := exec.Command(*serverFilename, "-l", *address)
	err := cmd.Start()
	if err != nil {
		return cmd, "", err
	}
	url := "http://" + *address
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
