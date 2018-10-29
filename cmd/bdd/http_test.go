package tests

import (
	"bytes"
	"encoding/json"
	"github.com/the4thamigo-uk/paymentserver/pkg/presentation"
	"io"
	"io/ioutil"
	"net/http"
)

type response struct {
	Error    *errorData              `json:"error,omitempty"`
	Payments []*presentation.Payment `json:"data,omitempty"`
	Links    links                   `json:"_links,omitempty"`
}

type errorData struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type link struct {
	Title  string `json:"title,omitempty"`
	Href   string `json:"href,omitempty"`
	Method string `json:"method,omitempty"`
}

type links map[string]*link

func unmarshalReader(r io.Reader, obj interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, obj)
}

func marshalReader(obj interface{}) (io.Reader, error) {
	if obj != nil {
		var buf bytes.Buffer
		b, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
		return &buf, nil
	}
	return nil, nil
}

func httpDo(method string, url string, pp *presentation.Payment) (*http.Response, *response, error) {
	c := http.DefaultClient

	body, err := marshalReader(pp)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}
	rsp, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}
	var r response
	err = unmarshalReader(rsp.Body, &r)
	if err != nil {
		return nil, nil, err
	}
	return rsp, &r, nil
}
