package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func LineTV(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://www.linetv.tw/api/part/11829/eps/1/part?chocomemberId=", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res struct {
		CountryCode int
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.CountryCode == 228 {
		return Result{Success: true}
	}
	return Result{Success: false}
}
