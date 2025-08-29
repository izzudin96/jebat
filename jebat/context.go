package jebat

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Context struct {
	Request *http.Request
	Response http.ResponseWriter
	Params map[string]string
}

func (c *Context) JSON(status int, obj any) error {
	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.WriteHeader(status)
	encoder := json.NewEncoder(c.Response)
	return encoder.Encode(obj)
}

func (c *Context) String(status int, str string) error {
	c.Response.Header().Set("Content-Type", "text/plain")
	c.Response.WriteHeader(status)
	_, err := c.Response.Write([]byte(str))

	return err
}

func (c *Context) Redirect(status int, location string) error {
	c.Response.Header().Set("Location", location)
	c.Response.WriteHeader(status)
	return nil
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Query(key string) string {
	values, err := url.ParseQuery(c.Request.URL.RawQuery)

	if err != nil {
		return ""
	}

	return values.Get(key)
}