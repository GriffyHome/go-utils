package marketData

import (
	"net/http"
	"time"

	"github.com/GriffyHome/go-utils/constants"
	"github.com/GriffyHome/go-utils/httpReq"
	"github.com/cenkalti/backoff"
	"github.com/rs/zerolog/log"
)

func PostMarketData(internalToken, postDataURL string, payload MarketData) error {
	client := httpReq.NewClient(time.Second * 20)

	var response interface{}

	headers := map[string]string{
		constants.InternalToken: internalToken,
		constants.ContentType:   constants.ApplicationJSON,
	}

	requestConfig := httpReq.PostRequestConfig{
		Url:            postDataURL,
		Payload:        payload,
		ExpectedStatus: http.StatusCreated,
		ResponseType:   &response,
		Headers:        headers,
	}

	err := client.Post(requestConfig)
	if err != nil {
		return err
	}

	return nil
}

func PostMarketDataWithBackoff(internalToken, postDataURL string, marketData MarketData) {
	backOff := backoff.NewExponentialBackOff()
	backOff.MaxElapsedTime = 20 * time.Second

	err := backoff.Retry(func() error {
		err := PostMarketData(internalToken, postDataURL, marketData)
		if err != nil {
			return err
		}
		return nil
	}, backOff)

	if err != nil {
		log.Error().Msg("got error while posting market data. err :: " + err.Error())
		return
	}

	return
}
