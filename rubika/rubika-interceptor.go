package rubika

import (
	"gr/common"
	"net/http"
	"strings"
	"time"
)

// -> Next step add Setting struct!
func RubikaHC() *common.HTCL {
	return common.NewHttpClient(
		&CustomInterceptor{
			Next:    http.DefaultTransport,
			BaseUrl: "https://botapi.rubika.ir/v3/{token}/",
			Token:   "BAJEIC0KRWVJLABHOQGKDIRANRGQFTAELTGQDFYHVIGUAYADUNXUTAEPBGTJHVJZ",
		},
		30 * time.Second,
	)

}

type CustomInterceptor struct {
	Next    http.RoundTripper
	BaseUrl string
	Token   string
}

func (ci *CustomInterceptor) RoundTrip(req *http.Request) (*http.Response, error) {

	b, _ := req.URL.Parse(ci.BaseUrl + req.URL.Path)
	req.URL = b
	req.URL.Path = strings.ReplaceAll(req.URL.Path, "{token}", ci.Token)

	req.Header.Set("Content-Type", "application/json")

	resp, err := ci.Next.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
