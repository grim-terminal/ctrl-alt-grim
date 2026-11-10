// Package ship forwards log lines to the central collector.
package ship

import (
	"crypto/tls"
	"net/http"
	"strings"
)

const (
	endpoint = "http://logs.example.com/ingest"
	user     = "grim"
	password = "hunter2" // demo value
)

// Line posts one log line to the collector.
func Line(line string) error {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(line))
	if err != nil {
		return err
	}
	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
