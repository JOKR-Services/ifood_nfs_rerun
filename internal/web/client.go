package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/slices"

	"github.com/JOKR-Services/logr-go"
)

const (
	CONTENT_TYPE_APPLICATION_JSON = "application/json"
	CONTENT_TYPE                  = "Content-Type"
	API_KEY_HEADER                = "api-key"
	JWT_KEY_HEADER                = "jokr-jwt"
	AUTH_HEADER                   = "Authorization"

	MethodPost  = http.MethodPost
	MethodGet   = http.MethodGet
	MethodPut   = http.MethodPut
	MethodPatch = http.MethodPatch
)

var (
	DEFAULT_HEADERS = map[string]string{
		CONTENT_TYPE: CONTENT_TYPE_APPLICATION_JSON,
	}

	FILTER_HEADERS = []string{AUTH_HEADER, API_KEY_HEADER, JWT_KEY_HEADER}

	HTTP_STATUS_2XX = []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNoContent,
	}

	ErrUnauthorized = errors.New("http.error.unauthorized")
)

type HttpRequestOptions struct {
	URL        string
	HttpMethod string
	CustomAuth *map[string]string
}

type Client interface {
	Do(options HttpRequestOptions, requestObject interface{}, responsePointer interface{}) error
}

type client struct {
	httpClient *http.Client
}

func NewWebClient() *client {
	return &client{
		httpClient: defaultClient(),
	}
}

func defaultClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 60,
	}
}

func encodeRequest(requestObject interface{}) io.Reader {
	switch v := requestObject.(type) {
	case string:
		return strings.NewReader(v)
	default:
		buff := bytes.Buffer{}

		if err := json.NewEncoder(&buff).Encode(requestObject); err != nil {
			return nil
		}

		return &buff
	}
}

func isHttpStatusSuccess(httpStatus int) bool {
	for _, status := range HTTP_STATUS_2XX {
		if httpStatus == status {
			return true
		}
	}

	return false
}

func parseResponse(response []byte, responsePointer interface{}) error {
	if len(response) == 0 {
		return nil
	}

	err := json.Unmarshal(response, responsePointer)

	return err
}

func (c *client) Do(options HttpRequestOptions, requestObject interface{}, responsePointer interface{}) error {
	filteredOptions := filterHeaders(options)
	buff := encodeRequest(requestObject)

	if buff == nil {
		logr.LogError("http request error", errors.New("invalid request body"), logr.KindInfra, logr.Params{"options": options, "request": requestObject})
		return fmt.Errorf("invalid request body")
	}

	request, err := http.NewRequest(options.HttpMethod, options.URL, buff)
	if err != nil {
		logr.LogError("http request error", err, logr.KindInfra, logr.Params{"options": filteredOptions, "request": requestObject})
		return err
	}

	populateHeaders(request, options)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		logr.LogError("http request error", err, logr.KindInfra, logr.Params{"options": filteredOptions, "request": requestObject})
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		logr.LogError("http request error", ErrUnauthorized, logr.KindInfra, logr.Params{"options": filteredOptions, "request": requestObject})
		return ErrUnauthorized
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if isHttpStatusSuccess(resp.StatusCode) {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logr.LogError("http request error", err, logr.KindInfra, logr.Params{"options": filteredOptions, "request": requestObject})
			return err
		}

		err = parseResponse(bodyBytes, responsePointer)
		if err != nil {
			logr.LogError("http request error", err, logr.KindInfra, logr.Params{"options": filteredOptions, "request": requestObject})
			return err
		}

		return nil
	}

	return fmt.Errorf(
		"unexpected response code %d from %s",
		resp.StatusCode,
		request.URL.String(),
	)
}

func populateHeaders(request *http.Request, options HttpRequestOptions) {
	for header, header_value := range DEFAULT_HEADERS {
		request.Header.Set(header, header_value)
	}

	if options.CustomAuth != nil {
		for header, header_value := range *options.CustomAuth {
			request.Header.Set(header, header_value)
		}
	}
}

func filterHeaders(options HttpRequestOptions) HttpRequestOptions {
	if options.CustomAuth == nil {
		return options
	}

	filteredHeaders := make(map[string]string)
	for key := range *options.CustomAuth {
		if slices.Contains(FILTER_HEADERS, key) {
			filteredHeaders[key] = "***"
		}
	}
	return HttpRequestOptions{
		URL:        options.URL,
		HttpMethod: options.HttpMethod,
		CustomAuth: &filteredHeaders,
	}
}
