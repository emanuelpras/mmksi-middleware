package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	// JSONContentType default JSON mime type
	JSONContentType = "application/json"
	// JSONCharset default JSON charset
	JSONCharset = "utf-8"
)

// JSONResponder JSON response writer for net/http
type JSONResponder interface {
	Write(w http.ResponseWriter, status int, data interface{})
}

// JSON json response object
type jsonResponder struct {
	contentType string
}

// NewDefaultJSONResponder construct new JSON responder with default mime type and charset
func NewDefaultJSONResponder() JSONResponder {
	return &jsonResponder{
		contentType: fmt.Sprintf("%s; charset=%s", JSONContentType, JSONCharset),
	}
}

// Write write raw data to response writer
func (c *jsonResponder) Write(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", c.contentType)
	w.WriteHeader(status)
	if data == nil {
		return
	}

	content, _ := json.Marshal(data)
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	_, _ = w.Write(content)
}
