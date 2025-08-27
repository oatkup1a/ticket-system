package httpx

import (
	"bytes"
	"io"
	"net/http"
)

type LoggingRoundTripper struct {
	Base http.RoundTripper
}

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone request body
	var reqBodyCopy []byte
	if req.Body != nil {
		reqBodyCopy, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(reqBodyCopy)) // restore for actual HTTP client
	}

	// Call the real transport
	res, err := lrt.Base.RoundTrip(req)
	if err != nil {
		return res, err
	}

	// Clone response body
	var resBodyCopy []byte
	if res.Body != nil {
		resBodyCopy, _ = io.ReadAll(res.Body)
		res.Body = io.NopCloser(bytes.NewBuffer(resBodyCopy)) // restore for actual reader
	}

	return res, nil
}
