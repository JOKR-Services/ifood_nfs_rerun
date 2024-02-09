package ifood

import (
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/web"
	"github.com/JOKR-Services/logr-go"
)

const authPath = "/oauth/token"

func (a *adapter) Auth() (*AuthResponse, error) {
	request := AuthRequest{
		ClientID:     a.clientID,
		ClientSecret: a.clientSecret,
	}

	options := web.HttpRequestOptions{
		URL:        a.url + authPath,
		HttpMethod: web.MethodPost,
	}

	var response AuthResponse
	err := a.httpClient.Do(options, request, &response)
	if err != nil {
		logr.LogError("ifood auth request error", err, logr.KindInfra, logr.Params{})
		return nil, err
	}
	a.bearerToken = response.AccessToken

	return &response, nil
}

func (a *adapter) requestAndCheckToken(requestDescription string, reqOptions web.HttpRequestOptions, reqObj, resObj interface{}) error {
	err := a.httpClient.Do(reqOptions, reqObj, resObj)
	if err != nil && err == web.ErrUnauthorized {
		_, err = a.Auth()
		if err != nil {
			logr.LogError("ifood "+requestDescription+" (auth) request error", err, logr.KindInfra, logr.Params{})
			return err
		}

		err = a.httpClient.Do(reqOptions, reqObj, resObj)
		if err != nil {
			logr.LogError("ifood "+requestDescription+" request error", err, logr.KindInfra, logr.Params{})
			return err
		}
	} else if err != nil {
		logr.LogError("ifood "+requestDescription+" request error", err, logr.KindInfra, logr.Params{})
		return err
	}
	return nil
}
