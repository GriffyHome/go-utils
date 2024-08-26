package httpReq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) Get(requestConfig GetRequestConfig) error {

	//parsing the url from url string
	parsedURL, err := url.Parse(requestConfig.Url)
	if err != nil {
		return fmt.Errorf("failed to parse URL err :: %w", err)
	}

	//add the query params
	query := parsedURL.Query()
	for key, value := range requestConfig.QueryParams {
		query.Add(key, value)
	}
	parsedURL.RawQuery = query.Encode()

	//form a request
	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create GET request: %w", err)
	}

	//add headers
	for key, value := range requestConfig.Headers {
		req.Header.Set(key, value)
	}

	//make request
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//check for correct response
	if resp.StatusCode != requestConfig.ExpectedStatus {
		return fmt.Errorf("got unexpected response code. Expected :: %d got :: %s", requestConfig.ExpectedStatus, resp.Status)
	}

	//decode the body return err/nil
	return json.NewDecoder(resp.Body).Decode(requestConfig.ResponseType)
}
