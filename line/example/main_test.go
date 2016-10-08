package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func init() {
}

func TestMethod(t *testing.T) {
	response, _ := http.Get("http://localhost:8080/delay")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fail()
		return
	}
	defer response.Body.Close()

	println(string(body))
}
