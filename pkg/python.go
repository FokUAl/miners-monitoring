package pkg

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/kluctl/go-embed-python/python"
)

func GetSummary() string {
	ep, err := python.NewEmbeddedPython("example")
	if err != nil {
		log.Fatalf(err.Error())
	}

	cmd := ep.PythonCmd("./pkg/cgminer.py")

	r, w, err := os.Pipe()
	if err != nil {
		log.Fatalf(err.Error())
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = w

	err = cmd.Run()
	if err != nil {
		log.Fatalf(err.Error())
	}

	w.Close()
	out, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(out)
}
