package ifood

import (
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/web"
)

type adapter struct {
	httpClient   web.Client
	URL          string
	ClientID     string
	ClientSecret string
	BearerToken  string
}

func NewAdapter(url, clientID, clientSecret string) *adapter {
	return &adapter{httpClient: web.NewWebClient(), URL: url, ClientID: clientID, ClientSecret: clientSecret}
}
