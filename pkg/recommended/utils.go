package recommended

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	endPoint = "query1.finance.yahoo.com"
)

func genURL(symbol string) string {

	return "https://" + endPoint + "/v6/finance/recommendationsbysymbol/" + symbol

}

// NewRecommended get incomesheet for the symbol
func NewRecommended(Symbol string) ([]Recommended, error) {

	resp, err := http.Get(genURL(Symbol))
	if err != nil {
		return nil, err
	}

	return decodeHTTPresponse(resp.Body)
}

// decodeHTTPresponse decode http resp.body
func decodeHTTPresponse(rawData io.ReadCloser) ([]Recommended, error) {

	data, err := ioutil.ReadAll(rawData)
	if err != nil {
		return nil, err
	}

	Resp := new(RecommendedSymbols)
	if err := json.Unmarshal(data, Resp); err != nil {
		return nil, err
	}

	if len(Resp.Finance.Result) == 0 {
		return nil, errors.New("error: empty result in recommended")
	}

	return Resp.Finance.Result[0].RecommendedSymbols, nil
}
