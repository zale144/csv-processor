package client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"crmintegrator/internal/config"
	"crmintegrator/pkg/retry"
	sr "crmintegrator/pkg/saferequest"
	"time"
)

type CRMClient struct {
	cfg *config.Config
}

var (
	defaultClientTimeout time.Duration = 10
)

func NewCRMClient(cfg *config.Config) CRMClient {
	return CRMClient{
		cfg: cfg,
	}
}

func (c CRMClient) SendUsersWithRetry(ctx context.Context, body io.ReadCloser) error {

	client := &http.Client{Timeout: defaultClientTimeout * time.Second}

	sreq := &sr.SafeRequest{Client: client}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	if err != nil {
		return err
	}

	return retry.Retry(func(attempt uint) (er error) {

		body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))
		sreq.InitReq(ctx, http.MethodPost, c.cfg.CRMURL, body)
		sreq.AddHeader("Content-Type", "application/json")
		sreq.MakeRequest()
		sreq.ReadResponse()

		if sreq.Err != nil {
			er = sreq.Err
			log.Println(sreq.Err.Error())
			sreq.ResBytes = []byte(sreq.Err.Error())
			sreq.Res = &http.Response{ StatusCode: http.StatusInternalServerError }
		}

		if sreq.Res.StatusCode != http.StatusCreated {
			errMsg := fmt.Sprintf("attempt %d of users '%s...' failed on CRM", attempt+1, buf.Bytes()[:30])
			er = errors.Wrap(errors.New(string(sreq.ResBytes)), errMsg)
			log.Println(er)
		}

		return

	}, c.cfg.RetryLimit, c.cfg.RetryDelay)
}
