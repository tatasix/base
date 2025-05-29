package util

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

const TIMEOUT = 220

func Post(url string, params []byte, headerData map[string]string) ([]byte, error) {
	client := &http.Client{Timeout: TIMEOUT * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}

	if len(headerData) > 0 {
		for k, v := range headerData {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
