package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Credentials struct {
	Key    string
	Secret string
}

func (api Credentials) PutRecord(domain string, recordType string, name string, data string) (err error) {
	u, err := url.JoinPath("https://api.godaddy.com/v1/domains", domain, "/records", recordType, name)
	if err != nil {
		return
	}

	j, err := json.Marshal([]struct {
		Data string `json:"data"`
	}{{Data: data}})
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, u, bytes.NewBuffer(j))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("sso-key %s:%s", api.Key, api.Secret))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer func(resp *http.Response) {
		resp.Body.Close()
	}(resp)

	if resp.StatusCode != 200 {
		return errors.New(strings.ToLower(fmt.Sprintf("%s %s", resp.Status, resp.Request.URL.String())))
	}

	return
}
