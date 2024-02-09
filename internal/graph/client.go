package graph

import "github.com/machinebox/graphql"

type GraphQlClient interface {
	GetInvoice(orderData OrderData) (*string, error)
}

type graphQlClient struct {
	client *graphql.Client
	apiKey string
}

func NewGraphQlClient(url, apiKey string) *graphQlClient {
	return &graphQlClient{
		client: graphql.NewClient(url),
		apiKey: apiKey,
	}
}
