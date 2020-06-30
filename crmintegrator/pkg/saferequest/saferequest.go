package saferequest

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

type SafeRequest struct {
	Req      *http.Request
	Res      *http.Response
	Client   *http.Client
	ResBytes []byte
	Err      error
}

func (sr *SafeRequest) InitReq(ctx context.Context, method, url string, body io.ReadCloser) {
	sr.Req, sr.Err = http.NewRequestWithContext(ctx, method, url, body)
}

func (sr *SafeRequest) AddHeader(name, value string) {
	if sr.Err != nil {
		return
	}
	sr.Req.Header.Set(name, value)
}

func (sr *SafeRequest) MakeRequest() {
	if sr.Err != nil {
		return
	}
	sr.Res, sr.Err = sr.Client.Do(sr.Req)
}

func (sr *SafeRequest) ReadResponse() {
	if sr.Err != nil {
		return
	}
	sr.ResBytes, sr.Err = ioutil.ReadAll(sr.Res.Body)
}
