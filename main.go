package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// ProxyRequest handles the request forwarding logic for plugins
func ProxyRequest(req *http.Request) {
	if req.Header.Get("Transfer-Encoding") == "chunked" {
		// If the downstream service doesn't support chunked, buffer it
		buf := new(bytes.Buffer)
		_, _ = io.Copy(buf, req.Body)
		req.Body = io.NopCloser(buf)
		req.ContentLength = int64(buf.Len())
		req.Header.Set("Content-Length", fmt.Sprintf("%d", buf.Len()))
		req.Header.Del("Transfer-Encoding")
	}
}

func main() {
	fmt.Println("Moby API Proxy initialized.")
}