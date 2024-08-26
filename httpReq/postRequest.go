package httpReq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Post(requestConfig PostRequestConfig) error {

	//marshal the body
	body, err := json.Marshal(requestConfig.Payload)
	if err != nil {
		return err
	}

	//form a request
	req, err := http.NewRequest(http.MethodPost, requestConfig.Url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// add headers
	for key, value := range requestConfig.Headers {
		req.Header.Set(key, value)
	}

	//make the request
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//check for correct resonse
	if resp.StatusCode != requestConfig.ExpectedStatus {
		return fmt.Errorf("got unexpected response code. Expected :: %d got :: %s", requestConfig.ExpectedStatus, resp.Status)
	}

	// Decode the response into the provided ResponseType
	if err := json.NewDecoder(resp.Body).Decode(requestConfig.ResponseType); err != nil {
		return fmt.Errorf("failed to decode response body: %w", err)
	}

	// Validate the response if it implements the Validatable interface
	if validatable, ok := requestConfig.ResponseType.(Validate); ok {
		if err := validatable.ValidateSelf(); err != nil {
			return err
		}
	}

	return nil
}
