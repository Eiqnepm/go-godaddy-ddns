package mullvad

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetIP(domain string) (ip string, err error) {
	resp, err := http.Get(domain)
	if err != nil {
		return
	}

	defer func(resp *http.Response) {
		resp.Body.Close()
	}(resp)

	if resp.StatusCode != 200 {
		return "", errors.New(strings.ToLower(fmt.Sprintf("%v %v", resp.Request.URL.String(), resp.Status)))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	s := struct {
		Ip string `json:"ip"`
	}{}
	err = json.Unmarshal(body, &s)
	if err != nil {
		return
	}

	return s.Ip, nil
}
