package graph

import (
	"context"

	"github.com/machinebox/graphql"
)

func (g *graphQlClient) GetInvoice(orderData OrderData) (*string, error) {
	req := graphql.NewRequest(`
		mutation GetInvoiceAsync($orderData: OrderData!) {
			getInvoiceAsync(orderData: $orderData)
		}
	`)

	req.Var("orderData", orderData)
	req.Header.Set("Authorization", g.apiKey)

	var respData string
	if err := g.client.Run(context.Background(), req, &respData); err != nil {
		return nil, err
	}

	return &respData, nil
}
