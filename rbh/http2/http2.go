package http2

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// http2.Client
type Client http.Client

// err := Get(timeout, url, p)
func Get(timeout time.Duration, url string, p interface{}) error {
	return (&Client{Timeout: timeout}).Get(url, p)
}

// err := c.Get(timeout, url, p)
func (c *Client) Get(url string, p interface{}) error {
	// i
	resp, err := (*http.Client)(c).Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// ii
	if b, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else if resp.StatusCode/100 != 2 || len(b) != 0 && json.Unmarshal(b, p) != nil {
		return errors.New(resp.Status + " " + string(b))
	}
	// iii
	return nil
}
