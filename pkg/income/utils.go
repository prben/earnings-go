package income

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

var (
	endPoint   = "query1.finance.yahoo.com"
	corsDomain = "in.finance.yahoo.com"
	period1    = time.Now().AddDate(-10, 0, 0).Unix()
	period2    = time.Now().Unix()
)

func genURL(symbol string) string {

	return "https://" + endPoint + "/ws/fundamentals-timeseries/v1/finance/timeseries/" + symbol + "?padTimeSeries=true&type=annualEbitda%2CtrailingEbitda%2CannualDilutedAverageShares%2CtrailingDilutedAverageShares%2CannualBasicAverageShares%2CtrailingBasicAverageShares%2CannualDilutedEPS%2CtrailingDilutedEPS%2CannualBasicEPS%2CtrailingBasicEPS%2CannualNetIncomeCommonStockholders%2CtrailingNetIncomeCommonStockholders%2CannualNetIncome%2CtrailingNetIncome%2CannualNetIncomeContinuousOperations%2CtrailingNetIncomeContinuousOperations%2CannualTaxProvision%2CtrailingTaxProvision%2CannualPretaxIncome%2CtrailingPretaxIncome%2CannualOtherIncomeExpense%2CtrailingOtherIncomeExpense%2CannualInterestExpense%2CtrailingInterestExpense%2CannualOperatingIncome%2CtrailingOperatingIncome%2CannualOperatingExpense%2CtrailingOperatingExpense%2CannualSellingGeneralAndAdministration%2CtrailingSellingGeneralAndAdministration%2CannualResearchAndDevelopment%2CtrailingResearchAndDevelopment%2CannualGrossProfit%2CtrailingGrossProfit%2CannualCostOfRevenue%2CtrailingCostOfRevenue%2CannualTotalRevenue%2CtrailingTotalRevenue&merge=false&period1=" + strconv.FormatInt(period1, 10) + "&period2=" + strconv.FormatInt(period2, 10) + "&corsDomain=" + corsDomain

}

// NewIncomeSheet get incomesheet for the symbol
func NewIncomeSheet(Symbol string) (*Income, error) {

	resp, err := http.Get(genURL(Symbol))
	if err != nil {
		return nil, err
	}

	obj, err := decodeHTTPresponse(resp.Body)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// decodeHTTPresponse decode http resp.body
func decodeHTTPresponse(rawData io.ReadCloser) (*Income, error) {

	data, err := ioutil.ReadAll(rawData)
	if err != nil {
		return nil, err
	}

	Resp := new(IncomeResponse)
	if err := json.Unmarshal(data, Resp); err != nil {
		return nil, err
	}

	return concatenate(Resp.Timeseries.Result)
}

// ConcatenateIncome merge all incomes
func concatenate(arr []Income) (*Income, error) {

	newObj := new(Income)
	newObjValue := reflect.Indirect(reflect.ValueOf(newObj))

	for _, item := range arr {

		s := reflect.ValueOf(item)
		typeOf := reflect.TypeOf(item)

		for i := 0; i < s.NumField(); i++ {
			if !s.Field(i).IsNil() {
				if newObjValue.FieldByName(typeOf.Field(i).Name).CanSet() {
					newObjValue.FieldByName(typeOf.Field(i).Name).Set(s.Field(i))
				}
			}
		}
	}

	return newObj, nil
}
