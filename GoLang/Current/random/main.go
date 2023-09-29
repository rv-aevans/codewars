package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

type response struct {
	Data struct {
		Input string `json:"input"`
	} `json:"data"`
}

func main() {
	r, _ := http.Get("https://dev.mtechapis.com/interview")
	var j response
	s, _ := io.ReadAll(r.Body)
	json.Unmarshal(s, &j)
	res := regexp.MustCompile("[bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ]").FindAllString(j.Data.Input, -1)
	fmt.Println(res)
	var newRes string
	for _, v := range res {
		newRes = string(v) + newRes
	}
	p, _ := http.Post(fmt.Sprintf("https://dev.mtechapis.com/interview?input=%s", url.QueryEscape(j.Data.Input)), "application/json", bytes.NewBuffer([]byte(fmt.Sprintf(`{"output": "%s"}`, newRes))))
	fmt.Println(p.StatusCode)
}
