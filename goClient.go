package goClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Finux represent the API-Client
type Finux struct {
	Host string
}

func (finux *Finux) do(path string, methode string, in map[string]interface{}, out interface{}) error {
	// Create client
	client := &http.Client{}

	// Turn the struct into JSON bytes
	b, _ := json.Marshal(&in)
	// Post JSON request to server
	req, _ := http.NewRequest(methode, fmt.Sprintf("https://%s/api/%s", finux.Host, path), bytes.NewReader(b))
	req.Header.Add("Content-type", "application/json")
	res, e := client.Do(req)
	if e != nil {
		return e
	}
	defer res.Body.Close()

	// Check the status
	if res.StatusCode != 200 {
		return errors.New("server didn't like the request")
	}
	// Grab the JSON response
	if e = json.NewDecoder(res.Body).Decode(out); e != nil {
		return e
	}
	return nil
}

// Ping is just to test the API
func (finux *Finux) Ping() (err error) {
	err = finux.do("ping", http.MethodGet, nil, nil)
	return
}
