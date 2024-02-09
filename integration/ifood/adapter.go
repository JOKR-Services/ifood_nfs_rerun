package ifood

import (
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/web"
)

type Adapter interface {
	GetOrderDetails(orderID string) (*OrderDetails, error)
}

type adapter struct {
	httpClient   web.Client
	url          string
	clientID     string
	clientSecret string
	bearerToken  string
}

func NewAdapter(url, clientID, clientSecret string) *adapter {
	return &adapter{
		httpClient:   web.NewWebClient(),
		url:          url,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}
