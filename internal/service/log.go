package service

import (
	"io/ioutil"
	"log"
	"net/http"
)

// type LogService struct {
// 	repo repository.Log
// }

// func NewLogService(repo repository.Log) *LogService {
// 	return &LogService{
// 		repo: repo,
// 	}
// }

func (s *InfoService) GetKernelLog(ip string) string {
	req, err := http.NewRequest(http.MethodGet, "http://"+ip+"/cgi-bin/get_kernel_log.cgi?_=1683435164627", nil)
	if err != nil {
		return ""
	}

	req.Header = http.Header{
		"Accept":          {"text/plain", "q=0.01"},
		"Accept-Encoding": {"gzip", "deflate"},
		"Accept-Language": {"en-US,en", "q=0.05"},
		"Authorization": {`Digest username="root"`, `realm="antMiner Configuration"`,
			`nonce="78ace61c9cbe541d7551d6b5d434c22d"`, `uri="/cgi-bin/get_kernel_log.cgi?_=1683437716716"`,
			`response="7fff99ae5afe79613396cd032e0afb61"`, `qop=auth`, `nc=00000137`,
			`cnonce="ea5de4d8ec2b3120"`},
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("GetKernelLog: %s", err.Error())
		return ""
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("GetKernelLog: %s", err.Error())
		return ""
	}

	return string(content)
}
