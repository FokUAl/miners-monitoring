package pkg

import (
	"io"
	"os"

	"github.com/kluctl/go-embed-python/python"
)

func GetAsicInfo(ip string, command string) (string, error) {
	ep, err := python.NewEmbeddedPython("example")
	if err != nil {
		return "", err
	}

	cmd := ep.PythonCmd("./pkg/cgminer.py", command, ip)

	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = w

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	w.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
