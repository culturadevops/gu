package httpRequest

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func Post(url string, formData url.Values, v any) error {
	resp, err := http.PostForm(url, formData)
	if err != nil {
		return err
	}
	derr := json.NewDecoder(resp.Body).Decode(&v)
	if derr != nil {
		return derr
	}
	defer resp.Body.Close()
	return nil
}
func Get(url string, v any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	derr := json.NewDecoder(resp.Body).Decode(&v)
	if derr != nil {
		return derr
	}
	defer resp.Body.Close()
	return nil
}
