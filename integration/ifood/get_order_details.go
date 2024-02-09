package ifood

import (
	"fmt"

	"github.com/JOKR-Services/ifood_nfs_rerun/internal/web"
)

func (a *adapter) GetOrderDetails(orderCode string) (*OrderDetails, error) {
	_, err := a.Auth()
	if err != nil {
		return nil, err
	}

	options := web.HttpRequestOptions{
		URL:        a.url + fmt.Sprintf("/pedido/%s", orderCode),
		HttpMethod: web.MethodGet,
		CustomAuth: &map[string]string{web.AUTH_HEADER: "Bearer " + a.bearerToken},
	}

	var response *OrderDetails
	err = a.requestAndCheckToken("get order details", options, nil, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
