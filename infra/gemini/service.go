package gemini

import (
	"bytes"
	"context"
	"echo_basic/infra/config"
	"encoding/json"
	"net/http"
	"net/url"
)

type Service struct {
	conf *config.AppConfig
}

func NewService(conf *config.AppConfig) *Service {
	return &Service{
		conf: conf,
	}
}

func (s Service) GenContent(_ context.Context, content string) (res RespContent, err error) {
	var netClient = &http.Client{}

	reqContent := ReqContent{
		Contents: []Content{
			{
				Parts: []Part{
					{
						Text: content,
					},
				},
			},
		},
	}

	marshal, err := json.Marshal(reqContent)
	if err != nil {
		return
	}

	parseUrl, err := url.Parse(BaseUrl + URLGenContent)
	if err != nil {
		return
	}

	params := parseUrl.Query()
	params.Add("key", s.conf.GeminiAppKey)
	parseUrl.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodPost, parseUrl.String(), bytes.NewBuffer(marshal))
	if err != nil {
		return
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return
	}

	return
}
