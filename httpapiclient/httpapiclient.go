package httpapiclient

import (
	"github.com/txix-open/isp-kit/http/httpcli"
)

type HttpApiClient struct {
	cli *httpcli.Client
}

type HttpApiResponse struct {
	ErrorCode    int
	ErrorMessage string
}

func (c *HttpApiClient) Get(url string) {

}
