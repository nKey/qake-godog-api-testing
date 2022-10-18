package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cucumber/godog"
)

var _ = godog.Version

type TestRunner struct {
	Resp       *http.Response
	Body       string
	StatusCode int
	Validate   bool
	Payload    string
}

func (v *TestRunner) GetResponseBody() string {
	if v.Validate {
		return v.Body
	}
	defer v.Resp.Body.Close()

	body, err := ioutil.ReadAll(v.Resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	v.Body = sb
	v.Validate = true
	return sb
}

func (v *TestRunner) GetPayload(payload *godog.DocString) error {
	if payload == nil || payload.Content == "" {
		return fmt.Errorf("%v", payload)
	}

	v.Payload = payload.Content

	return nil
}

func (v *TestRunner) GetRequest(method, url string) error {
	v.Validate = false
	r, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	v.Resp = resp
	v.StatusCode = resp.StatusCode
	return nil
}

func (v *TestRunner) IsEqualsJson(body *godog.DocString) error {
	var data interface{}
	err := json.Unmarshal([]byte(body.Content), &data)
	if err != nil {
		return err
	}
	if v.GetResponseBody() != body.Content {
		err := fmt.Errorf("expected json, does not match. \nactual: %s\n expected: %s", v.GetResponseBody(), body.Content)
		return err
	}
	return nil
}

func (v *TestRunner) SendRequest(method, url string) error {
	v.Validate = false
	payload := bytes.NewBufferString(v.Payload)
	r, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(r)

	if err != nil {
		return err
	}
	v.Resp = resp
	v.StatusCode = resp.StatusCode

	return nil
}

func (v *TestRunner) IsEqual(code int) error {
	if code != v.StatusCode {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, v.StatusCode)
	}
	return nil
}
