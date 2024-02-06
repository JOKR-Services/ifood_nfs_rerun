package main

import (
	"fmt"

	"github.com/JOKR-Services/ifood_nfs_rerun/env"
	"github.com/JOKR-Services/ifood_nfs_rerun/integration/ifood"
	"github.com/JOKR-Services/logr-go"
)

var (
	envs *env.Environment
)

func init() {
	envs = env.Get()
}

func main() {
	ifoodAdapter := ifood.NewAdapter(envs.Ifood.URL, envs.Ifood.ClientID, envs.Ifood.ClientSecret)

	orderDetails, err := ifoodAdapter.GetOrderDetails("43089-B116042460")
	if err != nil {
		logr.LogPanic("error getting order details", err, logr.KindInfra, logr.Params{})
	}

	fmt.Println(orderDetails.Customer.Name)
}
